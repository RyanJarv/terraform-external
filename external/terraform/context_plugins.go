// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package terraform

import (
	"github.com/hashicorp/terraform/external/addrs"
	"github.com/hashicorp/terraform/external/configs"
	"github.com/hashicorp/terraform/external/providers"
	"github.com/hashicorp/terraform/external/provisioners"
	"github.com/hashicorp/terraform/external/schemarepo"
	"github.com/hashicorp/terraform/external/schemarepo/loadschemas"
	"github.com/hashicorp/terraform/external/states"
)

// contextPlugins is a deprecated old name for loadschemas.Plugins
type contextPlugins = loadschemas.Plugins

func newContextPlugins(
	providerFactories map[addrs.Provider]providers.Factory,
	provisionerFactories map[string]provisioners.Factory,
	preloadedProviderSchemas map[addrs.Provider]providers.ProviderSchema,
) *loadschemas.Plugins {
	return loadschemas.NewPlugins(providerFactories, provisionerFactories, preloadedProviderSchemas)
}

// Schemas is a deprecated old name for schemarepo.Schemas
type Schemas = schemarepo.Schemas

func loadSchemas(config *configs.Config, state *states.State, plugins *loadschemas.Plugins) (*schemarepo.Schemas, error) {
	return loadschemas.LoadSchemas(config, state, plugins)
}
