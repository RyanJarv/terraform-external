// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package promising

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func init() {
	tracer = otel.Tracer("github.com/hashicorp/terraform/external/promising")
}
