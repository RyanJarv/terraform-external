// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	"github.com/hashicorp/terraform/external/builtin/providers/terraform"
	"github.com/hashicorp/terraform/external/grpcwrap"
	"github.com/hashicorp/terraform/external/plugin"
	"github.com/hashicorp/terraform/external/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(terraform.NewProvider())
		},
	})
}
