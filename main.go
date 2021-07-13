package main

import (
	"github.com/turbot/steampipe-plugin-pokemon/pokemon"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: pokemon.Plugin})
}
