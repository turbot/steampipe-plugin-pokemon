# Table: pokemon_pokemon

Pokémon are the creatures that inhabit the world of the Pokémon games. They can
be caught using Pokéballs and trained by battling with other Pokémon. Each
Pokémon belongs to a specific species but may take on a variant which makes it
differ from other Pokémon of the same species, such as base stats, available
abilities and typings.

## Examples

### Basic info

```sql
select
  name,
  id,
  abilities,
  stats,
  types
from
  pokemon_pokemon
```

### List all Pokémon heavier than 200 hg (20 kg)

```sql
select
  name,
  id,
  weight
from
  pokemon_pokemon
where
  weight >= 200
```
