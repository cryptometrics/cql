# meta

The `meta` library holds code to generate repetative files accross the codebase:

- protomodel
- model
- model accessors
- graph/schema
- endpoint accessors

# Building and Running

Build the dockerfile:

```
docker-compose -f "meta.docker-compose.yaml" up -d --build
```

Run the generate method:

```
docker-compose -f "meta.docker-compose.yaml" run generate
```

# Testing

To test run

```
docker-compose -f "meta.docker-compose.yaml" run test_generate
```
