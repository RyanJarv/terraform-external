// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	"github.com/hashicorp/terraform/external/grpcwrap"
	plugin "github.com/hashicorp/terraform/external/plugin6"
	simple "github.com/hashicorp/terraform/external/provider-simple-v6"
	"github.com/hashicorp/terraform/external/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
