package pokemon

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type pokemonConfig struct {
	MaxItems *int `cty:"max_items"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"max_items": {
		Type: schema.TypeInt,
	},
}

func ConfigInstance() interface{} {
	return &pokemonConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) pokemonConfig {
	if connection == nil || connection.Config == nil {
		return pokemonConfig{}
	}
	config, _ := connection.Config.(pokemonConfig)
	return config
}
