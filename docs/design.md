# design

sportsearch is a golang semantic search indexer and parser, create your own keywords and syntax, and infer the meaning of each search query.

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

### parse each word, give it a meaning

* find all teams, find all names

### find all names

names most likely will have a pair, like "lebron james", if the query only
had "james", this would cause ambiguity, but should considered and given most likely matches.  though a unique first name like "lebron" should give
a pretty accurate match


## sample queries

* `bos vs cle` # returns recent games' boxscore
* `lebron` # returns season stats since time is not mentioned
* `lebron last 5`
* `boston`

## player names

first and last names will be duplicated.  use a simple data tree
structure which is only 1 level deep, this will allow us to search efficiently.


## brainstorm

* first pass filter should identify each word in the query string without correlation, second pass should associate meaning, combining multiple words like first/last name
* the checks are for exact matches and will not provide a likely match