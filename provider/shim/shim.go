package shim

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	p "github.com/timescale/terraform-provider-timescale/internal/provider"
)

func NewProvider(version string) func() provider.Provider {
	return p.New(version)
}