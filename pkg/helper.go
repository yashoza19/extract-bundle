package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	bundle "github.com/yashoza19/extract-bundles/pkg/bundle"

	log "github.com/sirupsen/logrus"
)

const Docker = "docker"
const Podman = "podman"

func GenerateTemporaryDirs() {
	command := exec.Command("rm", "-rf", "tmp")
	_, _ = RunCommand(command)

	command = exec.Command("rm", "-rf", "./output/")
	_, _ = RunCommand(command)

	command = exec.Command("mkdir", "./output/")
	_, err := RunCommand(command)
	if err != nil {
		log.Fatal(err)
	}

	command = exec.Command("mkdir", "tmp")
	_, err = RunCommand(command)
	if err != nil {
		log.Fatal(err)
	}
}

func RunCommand(cmd *exec.Cmd) ([]byte, error) {
	command := strings.Join(cmd.Args, " ")
	log.Infof("running: %s\n", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("%s failed with error: (%v) %s", command, err, string(output))
	}
	if len(output) > 0 {
		log.Debugf("command output :%s", output)
	}
	return output, nil
}

const catalogIndex = "audit-catalog-index"

func ExtractIndexDB(image string, containerEngine string) error {
	log.Info("Extracting database...")
	// Remove image if exists already
	command := exec.Command(containerEngine, "rm", catalogIndex)
	_, _ = RunCommand(command)

	// Download the image
	command = exec.Command(containerEngine, "create", "--name", catalogIndex, image, "\"yes\"")
	_, err := RunCommand(command)
	if err != nil {
		return fmt.Errorf("unable to create container image %s : %s", image, err)
	}

	// Extract
	command = exec.Command(containerEngine, "cp", fmt.Sprintf("%s:/database/index.db", catalogIndex), "./output/")
	_, err = RunCommand(command)
	if err != nil {
		return fmt.Errorf("unable to extract the image for index.db %s : %s", image, err)
	}
	return nil
}

func GetDataFromIndexDB() error {
	// Connect to the database
	db, err := sql.Open("sqlite3", "./output/index.db")
	if err != nil {
		return fmt.Errorf("unable to connect in to the database : %s", err)
	}

	sql, err := bundle.BuildBundlesQuery()
	if err != nil {
		return err
	}

	row, err := db.Query(sql)
	if err != nil {
		return fmt.Errorf("unable to query the index db : %s", err)
	}

	auditBundle := make(map[string]string)

	defer row.Close()
	for row.Next() {
		var bundleName string
		var bundlePath string

		err = row.Scan(&bundleName, &bundlePath)
		if err != nil {
			log.Errorf("unable to scan data from index %s\n", err.Error())
		}
		log.Infof("Generating data from the bundle (%s)", bundleName)

		//map to get bundleName,BundlePath
		//write to json file

		// the csv is pruned from the database to save space.
		// See that is store only what is needed to populate the package manifest on cluster, all the extra
		// manifests are pruned to save storage space

		defer row.Close()

		auditBundle[bundleName] = bundlePath
		//auditBundle = add(auditBundle, map[string]string{bundleName: bundlePath})
		//build map here
	}
	//marshal auditBundle to json file
	jsonStr, err := json.Marshal(auditBundle)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	err = ioutil.WriteFile("test.json", jsonStr, 0644)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return nil
}
