#!/bin/bash

PROJECT_NAME="${1:-my-go-app}"

echo "Creating instructor template: $PROJECT_NAME"

# Create directory tree
mkdir -p "$PROJECT_NAME"/cmd/api
mkdir -p "$PROJECT_NAME"/internal/api/{handlers,middlewares}
mkdir -p "$PROJECT_NAME"/internal/models
mkdir -p "$PROJECT_NAME"/internal/repositories/{mongodb,sqlconnect}
mkdir -p "$PROJECT_NAME"/pkg/utils
mkdir -p "$PROJECT_NAME"/proto

# Create starter files
touch "$PROJECT_NAME"/cmd/api/{server.go,.env}
touch "$PROJECT_NAME"/internal/api/routers.go
touch "$PROJECT_NAME"/internal/repositories/mongodb/mongoconnect.go
touch "$PROJECT_NAME"/internal/repositories/sqlconnect/sqlconfig.go
touch "$PROJECT_NAME"/pkg/utils/{error_handling,jwt_processing}.go
touch "$PROJECT_NAME"/{go.mod,go.sum}

echo "Done! Run 'cd $PROJECT_NAME' to navigate into your project."
