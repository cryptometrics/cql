# cql

cql takes json models in schema/model and generates a corresponding go models, graphql types, endpoint references. The idea is to programtically generate the more mundane boilerplate work in building api pipelines. cql package can be used as a general-purpose microserverice for algorithm work, data analysis, etc.

## Installation

Clone the repository and run `go build`

## .env

To use api endpoints that require authentication, create a cryptometrics/cql/.env file and add your keys. The following are the supported APIs

### Coinbase Pro

```
CB_PRO_URL=
CB_PRO_ACCESS_PASSPHRASE=
CB_PRO_ACCESS_KEY=
CB_PRO_SECRET=
```

### Kraken

```
KRAKEN_URL=
KRAKEN_KEY=
KRAKEN_SECRET=
```

### IEX

```
IEX_URL=
IEX_KEY=
IEX_SECRET=
```

## Usage

To start the graphql server run

```sh
$ go run server.go
```

Then go to http://localhost:8080/

## Updating the Schema

To add new endpoints, add a json file to schema/model and run the generate script:

```sh
python3 generate.py
```

To re-generate a change to the schema:

```sh
$ go run github.com/99designs/gqlgen generate
```

## Docs

- [Coinbase Pro Asyncronous Websocket Client Documentation](https://readthedocs.org/projects/copra/downloads/pdf/latest/)
- [Coinbase Pro API Reference](https://docs.pro.coinbase.com/)
- [IEX Cloud API Reference](https://iexcloud.io/docs/api/)
- [Kraken API Reference](https://docs.kraken.com/rest/)
