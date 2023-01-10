package main

import (
	"fmt"
	"os"
	"path"

	"github.com/pkg/profile"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"modernc.org/mathutil"
)

var (
	rootCmd = &cobra.Command{
		Use:                path.Base(os.Args[0]),
		PersistentPreRunE:  rootPersistentPreRunE,
		PersistentPostRunE: rootPersistentPostRunE,
		RunE:               rootRunE,
		SilenceUsage:       true,
	}
	rootCmdFlagNames = struct {
		verbose string
		profile string
	}{"verbose", "profile"}
	rootCmdFlags struct {
		verbose int
		profile string
	}
	rootCmdState struct {
		profile interface{ Stop() }
	}
)

func init() {
	rootCmd.PersistentFlags().CountVarP(&rootCmdFlags.verbose, rootCmdFlagNames.verbose, "v", "counted verbosity")
	rootCmd.PersistentFlags().StringVar(&rootCmdFlags.profile, rootCmdFlagNames.profile, "", "profile type (cpu, clock)")
}

func rootPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	var setLevel = uint32(logrus.GetLevel()) + uint32(rootCmdFlags.verbose)
	var maxLevel = uint32(logrus.TraceLevel)
	var newLevel = mathutil.MinUint32(setLevel, maxLevel)

	logrus.SetLevel(logrus.Level(newLevel))
	log := logrus.WithFields(logrus.Fields{"verbosity": rootCmdFlags.verbose, "setLevel": setLevel, "maxLevel": maxLevel, "newLevel": newLevel, "logrus.GetLevel()": logrus.GetLevel()})
	log.Debugf("info")

	switch rootCmdFlags.profile {
	case "":
		// no profiling
	case "cpu":
		rootCmdState.profile = profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	case "clock":
		rootCmdState.profile = profile.Start(profile.ClockProfile, profile.ProfilePath("."))
	default:
		return fmt.Errorf("invalid profile type: %q", rootCmdFlags.profile)
	}

	return nil
}

func rootPersistentPostRunE(cmd *cobra.Command, args []string) (err error) {
	if rootCmdState.profile != nil {
		rootCmdState.profile.Stop()
	}
	return nil
}

func rootRunE(cmd *cobra.Command, args []string) error {
	logrus.Debugf("entry: %v", cmd.Use)
	return nil
}
