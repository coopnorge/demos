package main

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"modernc.org/mathutil"
)

var (
	rootCmd = &cobra.Command{
		Use:               path.Base(os.Args[0]),
		PersistentPreRunE: rootPersistentPreRunE,
		RunE:              rootRunE,
		SilenceUsage:      true,
	}
)

func init() {
	var rootPFlags = rootCmd.PersistentFlags()
	rootPFlags.CountP("verbose", "v", "counted verbosity")
}

func rootPersistentPreRunE(cmd *cobra.Command, args []string) error {
	verbosity, err := cmd.Flags().GetCount("verbose")
	if err != nil {
		return err
	}

	var setLevel = uint32(logrus.GetLevel()) + uint32(verbosity)
	var maxLevel = uint32(logrus.TraceLevel)
	var newLevel = mathutil.MinUint32(setLevel, maxLevel)

	logrus.SetLevel(logrus.Level(newLevel))
	log := logrus.WithFields(logrus.Fields{"verbosity": verbosity, "setLevel": setLevel, "maxLevel": maxLevel, "newLevel": newLevel, "logrus.GetLevel()": logrus.GetLevel()})
	log.Debugf("info")

	return nil
}

func rootRunE(cmd *cobra.Command, args []string) error {
	logrus.Debugf("entry: %v", cmd.Use)
	return nil
}


