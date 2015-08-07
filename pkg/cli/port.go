// Copyright (C) 2015 Scaleway. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package cli

import (
	"github.com/scaleway/scaleway-cli/vendor/github.com/Sirupsen/logrus"

	"github.com/scaleway/scaleway-cli/pkg/commands"
)

var cmdPort = &Command{
	Exec:        runPort,
	UsageLine:   "port [OPTIONS] SERVER [PRIVATE_PORT[/PROTO]]",
	Description: "Lookup the public-facing port that is NAT-ed to PRIVATE_PORT",
	Help:        "List port mappings for the SERVER, or lookup the public-facing port that is NAT-ed to the PRIVATE_PORT",
}

func init() {
	cmdPort.Flag.BoolVar(&portHelp, []string{"h", "-help"}, false, "Print usage")
	cmdPort.Flag.StringVar(&portGateway, []string{"g", "-gateway"}, "", "Use a SSH gateway")
}

// FLags
var portHelp bool      // -h, --help flag
var portGateway string // -g, --gateway flag

func runPort(cmd *Command, rawArgs []string) {
	if portHelp {
		cmd.PrintUsage()
	}
	if len(rawArgs) < 1 {
		cmd.PrintShortUsage()
	}

	args := commands.PortArgs{
		Gateway: portGateway,
		Server:  rawArgs[0],
	}
	ctx := cmd.GetContext(rawArgs)
	err := commands.RunPort(ctx, args)
	if err != nil {
		logrus.Fatalf("Cannot execute 'port': %v", err)
	}
}
