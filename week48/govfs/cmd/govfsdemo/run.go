package main

import (
	"example.com/coopnorge/govfs/internal/processor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	runCmd = &cobra.Command{
		Use:  "run",
		Args: cobra.ExactArgs(0),
		RunE: runCmdRunE,
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("url", "", "", "location")
	runCmd.MarkFlagRequired("url")
	runCmd.Flags().StringP("mark", "", "", "mark")
	runCmd.MarkFlagRequired("mark")
}

func runCmdRunE(cmd *cobra.Command, args []string) error {
	logrus.Infof("starting")

	u, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}

	m, err := cmd.Flags().GetString("mark")
	if err != nil {
		return err
	}

	p := processor.Processor{URL: u, Mark: m}
	count, err := p.Run()
	if err != nil {
		return err
	}
	logrus.WithField("count", count).Infof("done")
	return nil
}
