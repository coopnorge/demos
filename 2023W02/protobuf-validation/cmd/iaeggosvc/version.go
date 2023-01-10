package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/aucampia/eg/service-golang/internal/version"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Args:  cobra.ExactArgs(0),
		RunE:  versionCmdRunE,
		Short: "Print version information",
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func versionCmdRunE(cmd *cobra.Command, args []string) (err error) {
	info, err := version.Get()
	if err != nil {
		return err
	}
	infos, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	_, err = fmt.Printf("%s\n", infos)
	if err != nil {
		return err
	}
	return nil
}
