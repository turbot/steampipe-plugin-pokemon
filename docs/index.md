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
```

```text
+-------+----------+---------------------------------------------------------------------+
| score | comments | title                                                               |
+-------+----------+---------------------------------------------------------------------+
| 242   | 300      | Query Pokemon API with SQL                                      |
| 121   | 127      | Why Uber Engineering Switched from Postgres to MySQL (2016)         |
| 70    | 12       | Show HN: QueryCal – calculate metrics from your calendars using SQL |
| 17    | 10       | Global Associative Arrays in PostgreSQL                             |
+-------+----------+---------------------------------------------------------------------+
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
