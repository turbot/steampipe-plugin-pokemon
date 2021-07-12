package pokemon

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-pokemon",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		//DefaultTransform: transform.FromJSONTag().NullIfZero(),
		DefaultTransform: transform.FromCamel(),
		DefaultConcurrency: &plugin.DefaultConcurrencyConfig{
			TotalMaxConcurrency:   500,
			DefaultMaxConcurrency: 200,
		},
		TableMap: map[string]*plugin.Table{
			"pokemon_pokemon": tablePokemonPokemon(ctx),
		},
	}
	return p
}
