#!/bin/bash

# init data directory
initdb -D ${PGDATA}

# start in background
pg_ctl -D ${PGDATA} -l ${LOGDIR}/logfile start

# wait for startup to finish
wait_postgresql() {
  while ! pg_isready -q; do
    echo "Starting postgres..."
    sleep 1
  done
}
wait_postgresql

# create database
psql postgres://postgres:password@localhost:5432 -c 'CREATE DATABASE server;'

# run initial migrations required for seeding
psql postgres://postgres:password@localhost:5432/server -f migrations/001_init.sql

# loop through data seed objects and insert data
jq -c '.[]' data/db_data.json | while read -r i; do
  psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
    '$(jq -r '.iri' <<< "$i")',
    '$(jq -r '.jsonld' <<< "$i")'
  );"
done

# stop background process
pg_ctl -D ${PGDATA} stop
