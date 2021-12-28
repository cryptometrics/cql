#Deriving the latest base image
FROM ruby:latest

WORKDIR /usr/src/schema

# copy the /go partition from the cql_go_generate container, this gives us
# access to goimports
COPY --from=cql_goimports /go /go
COPY Gemfile ./
COPY Gemfile.lock ./
COPY test ./

ENV PATH="/usr/local/go/bin:${PATH}"

RUN bundle install
