/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yashoza19/extract-bundles/pkg"
	"github.com/yashoza19/extract-bundles/pkg/bundle"
)

var flags = bundle.Flags{}

// extractBundleCmd represents the extractBundle command
func NewCmd() *cobra.Command {
	extractBundleCmd := &cobra.Command{
		Use:   "extractBundle",
		Short: "A brief description of your command",
		Long: "A longer description that spans multiple lines and likely contains examples" +
			"and usage of using your command. For example:" +
			"Cobra is a CLI library for Go that empowers applications." +
			"This application is a tool to generate the needed files" +
			"to quickly create a Cobra application.",
		RunE: run,
	}

	extractBundleCmd.Flags().StringVar(&flags.IndexImage, "index-image", "",
		"index image and tag which will be audit")

	if err := extractBundleCmd.MarkFlagRequired("index-image"); err != nil {
		log.Fatalf("Failed to mark `index-image` flag for `index` sub-command as required")
	}
	extractBundleCmd.Flags().StringVar(&flags.ContainerEngine, "container-engine", pkg.Docker,
		fmt.Sprintf("specifies the container tool to use. If not set, the default value is docker. "+
			"Note that you can use the environment variable CONTAINER_ENGINE to inform this option. "+
			"[Options: %s and %s]", pkg.Docker, pkg.Podman))

	return extractBundleCmd

}

func run(cmd *cobra.Command, args []string) error {
	log.Info("Running capabilities run function")

	pkg.GenerateTemporaryDirs()

	if err := pkg.ExtractIndexDB(flags.IndexImage, flags.ContainerEngine); err != nil {
		log.Fatalf("Unable to ExtractIndexDB: %v", err)
	}
	if err := pkg.GetDataFromIndexDB(); err != nil {
		log.Fatalf("Unable to GetDataFromIndexDB: %v", err)
	}

	return nil
}
