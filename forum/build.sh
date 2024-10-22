#!/usr/bin/env bash

printHeader() {
    text=$1
    echo -e "\n$text"
    echo "--------------------------------------------------------------------------------"
}

DB_FILE="data.db"

# Create the database and initialize schema if the database does not exist
if [ ! -f "$DB_FILE" ]; then
    echo "$DB_FILE not found. Creating a new database..."

    # Use sqlite3 to create tables and initialize the database schema
    sqlite3 $DB_FILE <<EOF
  CREATE TABLE IF NOT EXISTS example_table (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );
EOF

else
    echo "$DB_FILE already exists. Skipping database creation..."
fi

printHeader "Building Space Chat Forum docker image"
docker image build -f Dockerfile -t forum . && printHeader "Building Space Chat Forum image done" || exit 1

printHeader "Running the docker container"
docker container run -p 8080:8080 -d --name forum forum && printHeader "Space Chat Forum container now running on port 8080" || exit 1

printHeader "Pruning unused docker objects"
docker system prune -f && printHeader "Basic prune complete" || exit 1

printHeader "Remove additonal unused images"
docker rmi alpine:3.20 && docker rmi golang:1.22.4-alpine && printHeader "Pruning unused images completed" || exit 1

printHeader "Space Chat Forum now available at http://localhost:8080/"
