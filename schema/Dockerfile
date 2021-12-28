#Deriving the latest base image
FROM ruby:latest

WORKDIR /usr/src/schema

COPY Gemfile ./
COPY Gemfile.lock ./
COPY test ./

RUN bundle install
