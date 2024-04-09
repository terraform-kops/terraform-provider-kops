package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/terraform-kops/terraform-provider-kops/pkg/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.NewProvider})
}
