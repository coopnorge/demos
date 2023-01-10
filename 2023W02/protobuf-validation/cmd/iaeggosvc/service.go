package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/aucampia/eg/service-golang/internal/config"
	"gitlab.com/aucampia/eg/service-golang/internal/server"
)

var (
	serviceCmd = &cobra.Command{
		Use: "service",
	}
	dumpConfigCmd = &cobra.Command{
		Use:  "dump-config",
		Args: cobra.ExactArgs(0),
		RunE: dumpConfigCmdRunE,
	}
	healthcheckCmd = &cobra.Command{
		Use:  "check-health",
		Args: cobra.ExactArgs(0),
		RunE: healthcheckCmdRunE,
	}
	runCmd = &cobra.Command{
		Use:  "run",
		Args: cobra.ExactArgs(0),
		RunE: runCmdRunE,
	}
)

func init() {
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(dumpConfigCmd)
	serviceCmd.AddCommand(healthcheckCmd)
	serviceCmd.AddCommand(runCmd)
}

func dumpConfigCmdRunE(cmd *cobra.Command, args []string) (err error) {
	logrus.Debugf("entry: %s", commandPath(cmd))

	cfg := &config.ServerConfig{}
	err = config.Load(cfg, nil)
	if err != nil {
		return err
	}
	cfgStr, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Printf("%s\n", cfgStr)
	if err != nil {
		return err
	}
	return err
}

func healthcheckCmdRunE(cmd *cobra.Command, args []string) (err error) {
	cfg := &config.ServerConfig{}
	err = config.Load(cfg, nil)
	if err != nil {
		return err
	}

	conn, err := net.DialTimeout("tcp", cfg.Address(), time.Duration(3)*time.Second)

	if err != nil {
		return fmt.Errorf("error trying to connect to %s: %w", cfg.Address(), err)
	}
	defer func() {
		err := conn.Close()

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Fatalf("failed closing the healthcheck connection")
			panic(err)
		}
	}()

	_, err = fmt.Printf("OK\n")
	if err != nil {
		return err
	}
	return nil
}

func runCmdRunE(cmd *cobra.Command, args []string) (err error) {
	logrus.WithField("commandPath", commandPath(cmd)).Debug("entry")
	cfg := &config.ServerConfig{}
	err = config.Load(cfg, nil)
	if err != nil {
		return err
	}
	logrus.WithField("config", cfg).Info("config")

	logrus.Info("creating a new server")
	srv, err := server.NewServer(cmd.Context(), cfg)
	if err != nil {
		return fmt.Errorf("failed to create the server: %w", err)
	}
	logrus.Info("starting the server")
	err = srv.Start()
	if err != nil {
		return fmt.Errorf("failed to start the server: %w", err)
	}
	return nil
}
