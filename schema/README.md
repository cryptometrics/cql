# schema

The `schema` directory holds code to generate repetative files accross the codebase:

- protomodel
- model
- model accessors
- graph/schema
- endpoint accessors

# Building and Running

Build the dockerfile:

```
docker-compose up -d --build
```

Run the generate method:

```
docker-compose run generate
```

# Testing

To test run

```
docker-compose run test-generate
```