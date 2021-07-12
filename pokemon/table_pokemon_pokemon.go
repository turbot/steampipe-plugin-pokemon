package pokemon

import (
	"context"

	"github.com/mtslzr/pokeapi-go"

	"net/url"
	"strconv"

	//"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePokemonPokemon(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pokemon_pokemon",
		Description: "Pokémon are the creatures that inhabit the world of the Pokémon games.",
		List: &plugin.ListConfig{
			Hydrate: pokemonList,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate:    getPokemon,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The identifier for this resource.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromGo(),
			},
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
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemon,
			},
			{
				Name:        "forms",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "game_indices",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "height",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemon,
			},
			{
				Name:        "held_items",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "is_default",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getPokemon,
			},
			{
				Name:        "location_area_encounters",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getPokemon,
			},
			{
				Name:        "moves",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "order",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemon,
			},
			{
				Name:        "species",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "sprites",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "stats",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "types",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemon,
			},
			{
				Name:        "weight",
				Description: "A list of abilities this Pokémon could potentially have.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemon,
			},

			// Standard columns
			{
				Name:        "title",
				Description: "Pokemon name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func pokemonList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	offset := 0
	newOffset := 0

	logger := plugin.Logger(ctx)

	for true {
		l, err := pokeapi.Resource("pokemon", offset)

		if err != nil {
			plugin.Logger(ctx).Error("pokemon_pokemon.pokemonList", "query_error", err)
			return nil, err
		}

		for _, pokemon := range l.Results {
			d.StreamListItem(ctx, pokemon)
		}

		if len(l.Next) == 0 {
			break
		}

		u, err := url.Parse(l.Next)
		logger.Warn("URL", u)
		if err != nil {
			plugin.Logger(ctx).Error("pokemon_pokemon.pokemonList", "url_parse_error", err)
			return nil, err
		}

		m, _ := url.ParseQuery(u.RawQuery)
		logger.Warn("Raw query", m)
		newOffset, err = strconv.Atoi(m["offset"][0])
		logger.Warn("New offset", newOffset)
		if err != nil {
			plugin.Logger(ctx).Error("pokemon_pokemon.pokemonList", "str_conversion_error", err)
			return nil, err
		}

		offset = newOffset
		logger.Warn("Final offset", offset)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getPokemon(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	var name string
	//var id string

	/*
		if h.Item != nil {
			data := h.Item.(Pokemon)
			name = types.SafeString(data.Name)
		} else {
			name = d.KeyColumnQuals["name"].GetStringValue()
			id = d.KeyColumnQuals["id"].GetStringValue()
		}
	*/

	name = d.KeyColumnQuals["name"].GetStringValue()
	//id = d.KeyColumnQuals["id"].GetStringValue()

	logger.Warn("Name", name)

	l, err := pokeapi.Pokemon(name)
	if err != nil {
		plugin.Logger(ctx).Error("pokemon_pokemon.pokemonGet", "query_error", err)
		return nil, err
	}
	return l, nil
}
