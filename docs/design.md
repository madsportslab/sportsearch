# design

sportsndx is a golang semantic search indexer and parser, create your own keywords and syntax, and infer the meaning of each search query.

## semantic parsing

* who
  * player
  * team
  * can be none
* what
  * statistic
  * math operation
* when
  * default is the current season
  * last n
  * time range
  * time period
  * career
  * season
  * most recent
  * playoffs
  * regular season

## indexes

these will be resident in memory as maps, value will be
id

### players

* first
* last
* full
* short

### teams

* full
* city
* abv
* name
* conf
* div
