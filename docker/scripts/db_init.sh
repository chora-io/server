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

# seed database
/home/db/scripts/db_seed.sh

# stop background process
pg_ctl -D ${PGDATA} stop
