// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package config

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initSetCommand() *cobra.Command {
	setCommand := &cobra.Command{
		Use:     "set",
		Short:   "",
		Long:    "",
		Example: "",
		Args:    cobra.ExactArgs(2),
		Run:     runSetCommand,
	}
	return setCommand
}

func runSetCommand(cmd *cobra.Command, args []string) {
	// We're assuming the config file already exists, we should probably
	// create one in case it doesn't
	settingsKey := args[0]

	if settingsKey == "board_manager.additional_urls" {
		urls := strings.Split(args[1], ",")
		viper.Set(args[0], urls)
	} else {
		viper.Set(args[0], args[1])
	}
	viper.WriteConfig()
}
