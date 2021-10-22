# Table: pokemon_berry

Berries (Japanese: きのみ Tree Fruit) are small, juicy, fleshy fruit. As in the real world, a large variety exists in the Pokémon world, with a large range of flavors, names, and effects. First found in the Generation II games, many Berries have since become critical held items in battle, where their various effects include HP and status condition restoration, stat enhancement, and even damage negation.

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
  pokemon_berry
```

### List all Berry where size more than 100 mm

```sql
select
  name,
  id,
  size
from
  pokemon_berry
where
  size > 100
```

### List all Berry where soil_dryness between 10 and 30

```sql
select
  name,
  id,
  soil_dryness
from
  pokemon_berry
where 
  soil_dryness between 10 and 30
```
