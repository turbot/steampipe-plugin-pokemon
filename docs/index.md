---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/pokemon.svg"
brand_color: "#FF6600"
display_name: Pokémon
name: pokemon
description: Steampipe plugin to query Pokémon, items, moves, and more from PokéAPI.
---

# Pokémon

Pokémon is a media franchise that includes games, films, television shows, manga, and more.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  title,
  id,
  height,
  weight
from
  pokemon_pokemon
order by
  id
```

```text
+------------+----+--------+--------+
| title      | id | height | weight |
+------------+----+--------+--------+
| bulbasaur  | 1  | 7      | 69     |
| ivysaur    | 2  | 10     | 130    |
| venusaur   | 3  | 20     | 1000   |
| charmander | 4  | 6      | 85     |
| charmeleon | 5  | 11     | 190    |
| charizard  | 6  | 17     | 905    |
| squirtle   | 7  | 5      | 90     |
| wartortle  | 8  | 10     | 225    |
| blastoise  | 9  | 16     | 855    |
| caterpie   | 10 | 3      | 29     |
| metapod    | 11 | 7      | 99     |
| butterfree | 12 | 11     | 320    |
| weedle     | 13 | 3      | 32     |
| kakuna     | 14 | 6      | 100    |
| beedrill   | 15 | 10     | 295    |
| pidgey     | 16 | 3      | 18     |
| pidgeotto  | 17 | 11     | 300    |
| pidgeot    | 18 | 15     | 395    |
| rattata    | 19 | 3      | 35     |
| raticate   | 20 | 7      | 185    |
+------------+----+--------+--------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/pokemon/tables)**

## Get started

### Install

Download and install the latest Pokémon plugin:

```bash
steampipe plugin install pokemon
```

### Credentials

The [PokéAPI](https://pokeapi.co/docs/v2) is open to the public and does not require any credentials.

### Configuration

No configuration is needed. Installing the latest Pokémon plugin will create a config file (`~/.steampipe/config/pokemon.spc`) with a single connection named `pokemon`:
```hcl
connection "pokemon" {
  plugin    = "pokemon"
}
```

## Get Involved

* Open source: https://github.com/turbot/steampipe-plugin-pokemon
* Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
