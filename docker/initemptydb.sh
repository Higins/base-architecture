#!/usr/bin/env bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER blocks WITH
      LOGIN
      SUPERUSER
      CREATEDB
      CREATEROLE
      INHERIT
      REPLICATION
      CONNECTION LIMIT -1
      PASSWORD 'blocks';

    CREATE DATABASE blocks;

    GRANT ALL PRIVILEGES ON DATABASE blocks TO blocks;
EOSQL

echo 'user and database created done!'
