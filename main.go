package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/ncecere/terraform-provider-ufad/ad"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ad.Provider})
}
