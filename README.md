# graphql-go
Example API on graphql go using gqlgen
https://hasura.io/blog/building-a-graphql-api-with-golang-postgres-and-hasura/


## Setups
```
go install github.com/99designs/gqlgen@latest
# false import
go run github.com/99designs/gqlgen init
```

Insatll postgres
```
psql postgres
postgres# CREATE database graphql
```

```sql
CREATE TABLE movies (
	id serial PRIMARY KEY,
	title VARCHAR (50) NOT NULL,
	url VARCHAR (50) NOT NULL,
    release_date timestamp NOT NULL DEFAULT NOW()

);
```

## Run
```
go run server.go
```

### Request
```graphql
query getMovies {
  movies {
    title
    url
  }
}
```

### Response
```json
{
  "data": {
    "movies": [
      {
        "title": "Rise of GraphQL Warrior Pt1",
        "url": "https://riseofgraphqlwarriorpt1.com/"
      },
      {
        "title": "Iron Man",
        "url": "https://ironman.com"
      }
    ]
  }
}
```

Client starts at: localhost:38080

## Development
Changes to schema is made in `./graph/schema.graphqls` file and then we run `go run github.com/99designs/gqlgen generate`, to get generate the resolvers in `./graph/schema.resolvers.go` file.
