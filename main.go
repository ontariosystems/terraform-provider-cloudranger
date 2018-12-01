package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jaxi/terraform-provider-cloudranger/cloudranger"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudranger.Provider,
	})
}
