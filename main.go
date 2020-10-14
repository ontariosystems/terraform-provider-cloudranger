package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/ontariosystems/terraform-provider-cloudranger/cloudranger"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudranger.Provider,
	})
}
