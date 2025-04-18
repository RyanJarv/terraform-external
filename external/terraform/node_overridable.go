// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package terraform

import (
	"github.com/hashicorp/terraform/external/addrs"
	"github.com/hashicorp/terraform/external/configs"
)

// GraphNodeOverridable represents a node in the graph that can be overridden
// by the testing framework.
type GraphNodeOverridable interface {
	GraphNodeResourceInstance

	ConfigProvider() addrs.AbsProviderConfig
	SetOverride(override *configs.Override)
}
