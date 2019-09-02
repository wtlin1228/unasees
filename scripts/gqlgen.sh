#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f internal/gql/generated/generated.go \
    internal/gql/models/generated/generated.go \
    internal/gql/resolvers/generated/generated.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"