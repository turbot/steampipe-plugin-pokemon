package pokemon

import (
	"context"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePokemonPokemon(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pokemon_pokemon",
		Description: "Pokémon are the creatures that inhabit the world of the Pokémon games.",
		List: &plugin.ListConfig{
			Hydrate: listPokemon,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name"}),
			// TODO: Add support for 'id' key column
			//KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate: getPokemon,
			// Bad error message is a result of https://github.com/mtslzr/pokeapi-go/issues/29
			ShouldIgnoreError: isNotFoundError([]string{"invalid character 'N' looking for beginning of value"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name for this resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "abilities",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "base_experience",
				Description: "The base experience gained for defeating this Pokémon.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemon,
			},
			{
				Name:        "forms",
				Description: "A list of forms this Pokémon can take on.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "game_indices",
				Description: "A list of game indices relevant to Pokémon item by generation",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "height",
				Description: "The height of this Pokémon in decimeters.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemon,
			},
			{
				Name:        "held_items",
				Description: "A list of items this Pokémon may be holding when encountered.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "id",
				Description: "The identifier for this resource.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemon,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "is_default",
				Description: "Set for exactly one Pokémon used as the default for each species.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getPokemon,
			},
			{
				Name:        "location_area_encounters",
				Description: "A link to a list of location areas, as well as encounter details pertaining to specific versions.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getPokemon,
			},
			{
				Name:        "moves",
				Description: "A list of moves along with learn methods and level details pertaining to specific version groups.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "order",
				Description: "Order for sorting. Almost national order, except families are grouped together.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemon,
			},
			{
				Name:        "species",
				Description: "The species this Pokémon belongs to.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "sprites",
				Description: "A set of sprites used to depict this Pokémon in the game.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "stats",
				Description: "A list of base stat values for this Pokémon.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "types",
				Description: "A list of details showing types this Pokémon has.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "weight",
				Description: "The weight of this Pokémon in hectograms.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemon,
			},

			// Standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listPokemon(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("listPokemon")

	offset := 0

	for true {
		resources, err := pokeapi.Resource("pokemon", offset)

		if err != nil {
			plugin.Logger(ctx).Error("pokemon_pokemon.listPokemon", "query_error", err)
			return nil, err
		}

		for _, pokemon := range resources.Results {
			d.StreamListItem(ctx, pokemon)
		}

		// No next URL returned
		if len(resources.Next) == 0 {
			break
		}

		urlOffset, err := extractUrlOffset(resources.Next)
		if err != nil {
			plugin.Logger(ctx).Error("pokemon_pokemon.listPokemon", "extract_url_offset_error", err)
			return nil, err
		}

		// Set next offset
		offset = urlOffset
	}

	return nil, nil
}

func getPokemon(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("getPokemon")

	var name string

	if h.Item != nil {
		result := h.Item.(structs.Result)
		name = result.Name
	} else {
		name = d.KeyColumnQuals["name"].GetStringValue()
	}

	logger.Debug("Name", name)

	pokemon, err := pokeapi.Pokemon(name)

	if err != nil {
		plugin.Logger(ctx).Error("pokemon_pokemon.pokemonGet", "query_error", err)
		return nil, err
	}

	return pokemon, nil
}
