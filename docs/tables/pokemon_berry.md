# Table: pokemon_berry

Berries are small fruits that can provide HP and status condition restoration, stat enhancement, and even damage negation when eaten by Pokémon.

## Examples

### Basic info

```sql
select
  name,
  id, 
  growth_time,
  max_harvest,
  natural_gift_power,
  size,
  smoothness,
  soil_dryness
from 
  pokemon_berry;
```

### List all berries larger than 100 mm

```sql
select
  name,
  id,
  size
from
  pokemon_berry
where
  size > 100;
```

### List all berries with electric natural gift type

```sql
select
  name,
  id,
  jsonb_pretty(natural_gift_type) as natural_gift_type
from
  pokemon_berry
where 
  natural_gift_type ->> 'name' = 'electric';
```
