# Graph

To start the graphql server go to the `cmd/graphql` directory and generate the `graphql` binary with `go build`.  Then run

```
./graphql start --port=8080
```

and visit http://localhost:8080/

Here is an example query using the coinbase SDK:

```graphql
query {
  coinbaseWallets {
    id,
    destinationTagName,
    swiftDepositInformation {
      bankName
    },
    sepaDepositInformation {
      bankName
    },
    destinationTagRegex
  }
}
```


## Updating the Schema

To add new endpoints, add a json file to schema/model and run the generate script:

```sh
python3 generate.py
```

To re-generate a change to the schema:

```sh
$ gqlgen generate
```
