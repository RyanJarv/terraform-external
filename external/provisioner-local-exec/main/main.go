// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	localexec "github.com/hashicorp/terraform/external/builtin/provisioners/local-exec"
	"github.com/hashicorp/terraform/external/grpcwrap"
	"github.com/hashicorp/terraform/external/plugin"
	"github.com/hashicorp/terraform/external/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}
