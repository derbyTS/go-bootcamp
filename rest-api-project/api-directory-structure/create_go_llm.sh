#!/bin/bash

# Get project name from command line argument (defaults to 'my-go-app')
PROJECT_NAME="${1:-my-go-app}"

echo "Creating Go project template: $PROJECT_NAME"

# Create directory structure (-p creates nested directories automatically)
mkdir -p "$PROJECT_NAME"/{cmd/api,internal/{user,auth,platform},proto}

# Create starter files
touch "$PROJECT_NAME"/.env
touch "$PROJECT_NAME"/cmd/api/main.go
touch "$PROJECT_NAME"/internal/user/{handler,model,repository}.go
touch "$PROJECT_NAME"/internal/auth/{jwt,middleware}.go
touch "$PROJECT_NAME"/internal/platform/{database,mongodb,response}.go

echo "Done! Run 'cd $PROJECT_NAME' to start."
