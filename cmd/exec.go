// Copyright © 2018 Alexander Pinnecke <alexander.pinnecke@googlemail.com>
//

package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/scalify/puppet-master-cli/internal/pkg/exec"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec <directory>",
	Short: "Execute a single job",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.WithField("cmd", "exec")

		var baseDir string
		var err error
		if len(args) < 1 {
			if baseDir, err = os.Getwd(); err != nil {
				panic(err)
			}
		} else {
			baseDir = args[0]
		}

		execLogsVerbose, err := cmd.Flags().GetBool("executor-logs-verbose")
		if err != nil {
			log.Fatalf("failed to get executor-logs-verbose flag: %v", err)
		}

		del, err := cmd.Flags().GetBool("delete")
		if err != nil {
			log.Fatalf("failed to get delete flag: %v", err)
		}

		varsFile, err := cmd.Flags().GetString("vars")
		if err != nil {
			log.Fatalf("failed to get vars flag: %v", err)
		}

		codeFile, err := cmd.Flags().GetString("code")
		if err != nil {
			log.Fatalf("failed to get code flag: %v", err)
		}

		moduleFiles, err := cmd.Flags().GetStringSlice("module")
		if err != nil {
			log.Fatalf("failed to get modules flag: %v", err)
		}

		baseDir = path.Clean(baseDir)
		log.Infof("Executing jobs from directory %s", baseDir)

		if err := exec.Execute(log, client, baseDir, codeFile, varsFile, moduleFiles, del, execLogsVerbose); err != nil {
			log.Fatalf("failed to execute job: %v", err)
		}

		log.Infoln("Done.")
	},
}

func init() {
	RootCmd.AddCommand(execCmd)

	execCmd.Flags().Bool("delete", false, "Delete the job after execution")
	execCmd.Flags().Bool("executor-logs-verbose", false, "Verbose mode prints debug and notice logs from the executor")
	execCmd.Flags().String("vars", "vars.json", "JSON file containing variables")
	execCmd.Flags().String("code", "code.mjs", "MJS file containing the main execution code")
	execCmd.Flags().StringSlice("module", []string{}, "MJS files containing modules")
}
