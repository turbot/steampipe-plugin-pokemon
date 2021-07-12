---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/pokemon.svg"
brand_color: "#FF6600"
display_name: Pokemon
name: pokemon
description: Steampipe plugin to query stories, items and users from Pokemon.
---

# Pokemon

[Pokemon](https://news.ycombinator.com) is a social news website focusing on computer science and entrepreneurship. Steampipe marshalls the HN API data into queryable tables letting you interactivly explore it via our command line interface or your favorite SQL client. Example query:

```sql
select
  score,
  descendants as comments,
  title
from 
  pokemon_top
where
  type = 'story'
  and lower(title) like '%sql%'
order by
  score desc;
```
standard output (can use `.output` to change to `csv` or `json`):
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

## Installation

If you are just getting started with Steampipe, head over to https://steampipe.io/downloads to install the CLI (don't worry it will only take a minute). Once that is done to download and install the latest **Pokemon plugin**:

```bash
steampipe plugin install pokemon
```

## Browse the [Available Tables →](https://hub.steampipe.io/plugins/turbot/pokemon/tables)

```
> .inspect pokemon
+--------------------+-----------------------------------------------------------------------------+
| TABLE              | DESCRIPTION                                                                 |
+--------------------+-----------------------------------------------------------------------------+
| pokemon_ask_hn  | Latest 200 Ask HN stories.                                                  |
| pokemon_best    | Best 500 stories.                                                           |
| pokemon_item    | Stories, comments, jobs, Ask HNs and even polls are just items. This table  |
|                    | includes the most recent items posted to Pokemon.                       |
| pokemon_job     | Latest 200 Job stories.                                                     |
| pokemon_new     | Newest 500 stories.                                                         |
| pokemon_show_hn | Latest 200 Show HN stories.                                                 |
| pokemon_top     | Top 500 stories.                                                            |
| pokemon_user    | Information about Pokemon registered users who have public activity.    |
+--------------------+-----------------------------------------------------------------------------+
```

```
> .inspect pokemon_item
+-------------+-----------+---------------------------------------------------------------------------+
| COLUMN      | TYPE      | DESCRIPTION                                                               |
+-------------+-----------+---------------------------------------------------------------------------+
| by          | text      | The username of the item's author.                                        |
| dead        | boolean   | True if the item is dead.                                                 |
| deleted     | boolean   | True if the item is deleted.                                              |
| descendants | bigint    | In the case of stories or polls, the total comment count.                 |
| id          | bigint    | The item's unique id.                                                     |
| kids        | jsonb     | The ids of the item's comments, in ranked display order.                  |
| parent      | bigint    | The comment's parent: either another comment or the relevant story.       |
| parts       | jsonb     | A list of related pollopts, in display order.                             |
| poll        | bigint    | The pollopt's associated poll.                                            |
| score       | bigint    | The story's score, or the votes for a pollopt.                            |
| text        | text      | The comment, story or poll text. HTML.                                    |
| time        | timestamp | Timestamp when the item was created.                                      |
| title       | text      | The title of the story, poll or job. HTML.                                |
| type        | text      | The type of item. One of "job", "story", "comment", "poll", or "pollopt". |
| url         | text      | The URL of the story.                                                     |
+-------------+-----------+---------------------------------------------------------------------------+
```

## Credentials

The [Pokemon API](https://pokeapi.co/docs/v2) is open to the public and does not require any credentials.


## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest pokemon plugin will create a default connection named `pokemon` in the `~/.steampipe/config/pokemon.spc` file. You may edit this connection to set options:

```hcl
connection "pokemon" {
  plugin    = "pokemon"
  max_items = 5000
}
```
