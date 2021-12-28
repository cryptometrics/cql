# schema

The `schema` directory holds code to generate repetative files accross the codebase:

- protomodel
- model
- graph/schema

# Building and Running

Build the dockerfile:

```
docker-compose up -d --build
```

Run teh generate method:

```
docker-compose run generate
```

# Testing

To test run

```
docker-compose run test-generate
```
