# cql

cql takes json models in schema/model and generates a corresponding go models, graphql types, endpoint references. The idea is to programtically generate the more mundane boilerplate work in building api pipelines. cql package can be used as a general-purpose microserverice for algorithm work, data analysis, etc.

## Installing Bazel

Bazel is an open-source build and test tool similar to Make, Maven, and Gradle. It uses a human-readable, high-level build language. Bazel supports projects in multiple languages and builds outputs for multiple platforms. Bazel supports large codebases across multiple repositories, and large numbers of users.

To install, following the intructions [here](https://docs.bazel.build/versions/4.2.2/bazel-overview.html#how-do-i-use-bazel)

If you're on macOS, [you can install Bazel via Homebrew](https://docs.bazel.build/versions/4.2.2/install-os-x.html#step-2-install-bazel-via-homebrew):

```sh
$ brew install bazel
```

## Building Bazel

To build bazel run

```sh
$ bazel run //:gazelle
```

Then build
```sh
$ bazel build //...
```

This will generate a BUILD.bazel file with go_library, go_binary, and go_test targets for each package in your project. You can run the same command in the future to update existing build files with new source files, dependencies, and options.

## Updating Repos Bazel

Whenever we add new go dependencies we need to update the repos through bazel:

```sh
$ bazel run //:gazell-update-repos
```


## .auth.env

To use api endpoints that require authentication, create a cryptometrics/cql/.auth.env file and add your keys. The following are the supported APIs

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

## Graphql usage

To start the graphql server cd into cmd/graphql and build the binary

```sh
$ go build
```

then from the command line start the server:

```sh
$ ./graphql start --port=8080
```

## Updating the Schema

To add new endpoints, add a json file to schema/model and run the generate script:

```sh
python3 generate.py
```

To re-generate a change to the schema:

```sh
$ go get github.com/99designs/gqlgen; go run github.com/99designs/gqlgen generate
```

## Docs

- [Coinbase Pro Asyncronous Websocket Client Documentation](https://readthedocs.org/projects/copra/downloads/pdf/latest/)
- [Coinbase Pro API Reference](https://docs.pro.coinbase.com/)
- [IEX Cloud API Reference](https://iexcloud.io/docs/api/)
- [Kraken API Reference](https://docs.kraken.com/rest/)
