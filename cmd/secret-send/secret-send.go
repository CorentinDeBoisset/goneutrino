package main

import (
	"fmt"

	"github.com/corentindeboisset/secert-send/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	ExecutableName string
	Version        string

	rootCmd *cobra.Command

	confPath string
)

func init() {
	rootCmd = &cobra.Command{
		Use:   ExecutableName,
		Short: "Secret-send is a web server to securely send secrets to third-parties",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			verbosity, err := cmd.Flags().GetCount("verbose")
			if err != nil {
				return fmt.Errorf("failed to parse the verbosity options: %w", err)
			}
			logger.SetLogLevel(verbosity)

			return nil
		},
	}
	rootCmd.PersistentFlags().CountP("verbose", "v", "Add some verbosity to the output")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print version and exit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s: %s\n", ExecutableName, Version)
		},
	}
	rootCmd.AddCommand(versionCmd)

	runServerCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start a web server",
		Run: func(cmd *cobra.Command, args []string) {
			logger.InfoLog("Starting the server")

			// if err := watchRegistry(confPath); err != nil {
			// 	if verbosity >= 2 {
			// 		fmt.Printf("\n\nAn error occured when executing the command:\n%+v\n", err)
			// 	} else {
			// 		fmt.Printf("\n\nAn error occured when executing the command:\n%v\n", err)
			// 	}
			// 	os.Exit(1)
			// }
		},
	}
	runServerCmd.Flags().StringVarP(&confPath, "conf", "c", "", "Configuration file")
	if err := runServerCmd.MarkFlagRequired("conf"); err != nil {
		panic(fmt.Sprintf("An error occured when configuring the CLI: %v", err))
	}

	rootCmd.AddCommand(runServerCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.ErrorLog("An error occured when executing the command: %v", err)
	}
}
