# cql

## Installation 

To start the server run 

```sh
$ go run server.go
```

## Usage

To re-generate a change to the schema:

```sh
$ go run github.com/99designs/gqlgen generate
```

## cql/model V. cql/graph/model

There may be cases where the json being returned by a service is in snake case, rather than camel case. In this instances, we can't rely on gqlgen to generate the models. Instead, we have to make custom models. This is really straightforward, all you you need to do is add the custom model to cql/model and update the gqlgen.yml file:

```
  <gql schema name>:
    model: cql/model.<go struct name>
```
