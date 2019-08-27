#!/bin/sh
printf "\nRegenerating gqlgen file\n"
rm -f internal/gql/generated.go \
      internal/gql/models/generated.go \
      internal/gql/resolvers/generated.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"