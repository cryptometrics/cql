# Graph

## Updating the Schema

To add new endpoints, add a json file to schema/model and run the generate script:

```sh
python3 generate.py
```

To re-generate a change to the schema:

```sh
$ go get github.com/99designs/gqlgen; go run github.com/99designs/gqlgen generate
```
