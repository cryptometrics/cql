# cql

cql takes json models in schema/model and generates corresponding go models, graphql types, and endpoint references. The idea is to programtically generate boilerplate work in building api pipelines, then fill in the blanks in individual exportable packages, such as cql/coinbase, to be used as SDKs for market analysis.

## Getting Started
  * [Install Bazel](https://github.com/cryptometrics/cql#installing-bazel)
  * [Build Bazel](https://github.com/cryptometrics/cql#build-bazel)
  * [servers](https://github.com/cryptometrics/cql/blob/main/cmd/README.md)
    * [Graph](https://github.com/cryptometrics/cql/blob/main/graph/README.md)
      * [Updating the Schema](https://github.com/cryptometrics/cql/tree/main/graph#updating-the-schema)
  * Software Development Kits
    * [Coinbase](https://github.com/cryptometrics/cql/blob/main/coinbase/README.md)
      * [Connecting](https://github.com/cryptometrics/cql/blob/main/coinbase/README.md#connecting)
      * [WebSocket](https://github.com/cryptometrics/cql/blob/main/coinbase/README.md#websocket)
    * [Iex](https://github.com/cryptometrics/cql/blob/main/iex/README.md)
    * [Kraken](https://github.com/cryptometrics/cql/blob/main/kraken/README.md)
  * [Resources](https://github.com/cryptometrics/cql#resources)

## Install Bazel

Bazel is an open-source build and test tool similar to Make, Maven, and Gradle. It uses a human-readable, high-level build language. Bazel supports projects in multiple languages and builds outputs for multiple platforms. Bazel supports large codebases across multiple repositories, and large numbers of users.

To install, following the intructions [here](https://docs.bazel.build/versions/4.2.2/bazel-overview.html#how-do-i-use-bazel)

If you're on macOS, [you can install Bazel via Homebrew](https://docs.bazel.build/versions/4.2.2/install-os-x.html#step-2-install-bazel-via-homebrew):

```sh
$ brew install bazel
```

## Build Bazel

```
$ ./build.sh
```

## Resouces

- [Coinbase Pro Asyncronous Websocket Client Documentation](https://readthedocs.org/projects/copra/downloads/pdf/latest/)
- [Coinbase Pro API Reference](https://docs.pro.coinbase.com/)
- [IEX Cloud API Reference](https://iexcloud.io/docs/api/)
- [Kraken API Reference](https://docs.kraken.com/rest/)
- [Monorepo](https://en.wikipedia.org/wiki/Monorepo)
