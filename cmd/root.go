// Copyright © 2018 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"os"

	"github.com/Scalify/puppet-master-client-go"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type puppetMasterConfig struct {
	Verbose  bool   `ignore:"true"`
	Endpoint string `split_words:"true" required:"true" envconfig:"PUPPET_MASTER_ENDPOINT"`
	APIToken string `split_words:"true" required:"true" envconfig:"PUPPET_MASTER_API_TOKEN"`
}

var (
	client *puppetmaster.Client
	logger *logrus.Logger
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "puppet-master-cli",
	Short: "A simple and smart CLI for the puppet-master api.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config := &puppetMasterConfig{}
		envconfig.MustProcess("", config)
		logger = logrus.New()
		var err error

		if config.Verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
			logger.Fatalf("failed to get verbose flag: %v", )
		}

		if config.Verbose {
			logger.SetLevel(logrus.DebugLevel)
		} else {
			logger.SetLevel(logrus.InfoLevel)
		}

		client, err = puppetmaster.NewClient(config.Endpoint, config.APIToken)
		if err != nil {
			logger.Fatalf("failed to create puppet-master client: %v", err)
		}
	},
}

func init() {
	RootCmd.PersistentFlags().Bool("verbose", false, "Verbose mode enables detailed logs")
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
