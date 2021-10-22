package pokemon

import (
	"context"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePokemonBerry(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pokemon_berry",
		Description: "Berries are small fruits that can provide HP and status condition restoration, stat enhancement, and even damage negation when eaten by Pokémon.",
		List: &plugin.ListConfig{
			Hydrate: listBerry,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name"}),
			// TODO: Add support for 'id' key column
			// KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate: getBerry,
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
				Name:        "firmness",
				Description: "The firmness of this berry, used in making Pokéblocks or Poffins.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBerry,
			},
			{
				Name:        "flavors",
				Description: "A list of references to each flavor a berry can have and the potency of each of those flavors in regard to this berry.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBerry,
			},
			{
				Name:        "growth_time",
				Description: "Time it takes the tree to grow one stage, in hours. Berry trees go through four of these growth stages before they can be picked.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getBerry,
			},
			{
				Name:        "id",
				Description: "The identifier for this resource.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getBerry,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "item",
				Description: "Berries are actually items. This is a reference to the item specific data for this berry.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBerry,
			},
			{
				Name:        "max_harvest",
				Description: "The maximum number of these berries that can grow on one tree in Generation IV.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getBerry,
			},
			{
				Name:        "natural_gift_power",
				Description: "The power of the move 'Natural Gift' when used with this Berry.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getBerry,
			},
			{
				Name:        "natural_gift_type",
				Description: "The type inherited by 'Natural Gift' when used with this Berry.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBerry,
			},
			{
				Name:        "size",
				Description: "The size of this Berry, in millimeters.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getBerry,
			},
			{
				Name:        "smoothness",
				Description: "The smoothness of this Berry, used in making Pokéblocks or Poffins.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getBerry,
			},
			{
				Name:        "soil_dryness",
				Description: "The speed at which this Berry dries out the soil as it grows. A higher rate means the soil dries more quickly.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getBerry,
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

func listBerry(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("listBerry")

	offset := 0

	for true {
		resources, err := pokeapi.Resource("berry", offset)

		if err != nil {
			plugin.Logger(ctx).Error("pokemon_berry.listBerry", "query_error", err)
			return nil, err
		}

		for _, berry := range resources.Results {
			d.StreamListItem(ctx, berry)
		}

		// No next URL returned
		if len(resources.Next) == 0 {
			break
		}

		urlOffset, err := extractUrlOffset(resources.Next)
		if err != nil {
			plugin.Logger(ctx).Error("pokemon_berry.listBerry", "extract_url_offset_error", err)
			return nil, err
		}

		// Set next offset
		offset = urlOffset
	}

	return nil, nil
}

func getBerry(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("getBerry")

	var name string

	if h.Item != nil {
		result := h.Item.(structs.Result)
		name = result.Name
	} else {
		name = d.KeyColumnQuals["name"].GetStringValue()
	}

	logger.Debug("Name", name)

	berry, err := pokeapi.Berry(name)

	if err != nil {
		plugin.Logger(ctx).Error("pokemon_berry.pokemonGet", "query_error", err)
		return nil, err
	}

	return berry, nil
}
