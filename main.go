package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tuapuikia/terraform-provider-zabbix/provider"
)

func main() {
	p := plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	}

	plugin.Serve(&p)
}
