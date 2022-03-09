/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
//TODO
*/
package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yashoza19/extract-bundles/cmd"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "audit-tool",
		Short: "An analytic tool to audit operator bundles and index catalogs",
		Long: "The audit is an analytic tool which uses the Operator Framework solutions. " +
			"Its purpose is to obtain and report and aggregate data provided by checks and analyses done in " +
			"the operator bundles, packages and channels from an index catalog image.\n\n" +
			"Note that the latest version of the reports generated for all images can be checked in its repository, " +
			"inside of `testdata/report`. \n **NOTE** The file names has been created by using the kind/type " +
			"of the report, image name and date. " +
			"(E.g. `testdata/report/bundles_quay.io_operatorhubio_catalog_latest_2021-04-22.xlsx`)" +
			"For further information use the --help and check : https://github.com/operator-framework/audit",
	}

	rootCmd.AddCommand(cmd.NewCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
