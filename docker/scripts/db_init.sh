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

# run initial migrations required for seeding the database
psql postgres://postgres:password@localhost:5432/server -f migrations/001_init.sql

# seed database with data from chora genesis state (geonode metadata)
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVgcMUtYdPMWrUioXrSLoXfFdPJJqph9D8HMhU5LrkjAPfiN6Pbm.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/geonode.jsonld\", \"name\": \"geonode1\"}")'
);"

# seed database with data from chora genesis state (group metadata)
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVgqPqUq7pLobCdUsdakkEavthyoYmZtp1M3LH1x34rBNaheSyTK.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/group.jsonld\", \"name\": \"group1\"}")'
);"
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVgv2KxZTpf6HSFwHFRotLwSLjb3mvvn83wmcve9iBfpaoq1NDLZ.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/group_policy.jsonld\", \"name\": \"policy1\"}")'
);"
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVh21baKVbWVLEZLQuxwsh3NpnDieEnRXgnxLCpW9fUvGQE1vd4X.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/group_policy.jsonld\", \"name\": \"policy2\"}")'
);"
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVg41sKJB6fSDEAuUAxWsS39P5TG21Fsf2yonFzeBM9rh2xWWzLt.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/group_policy.jsonld\", \"name\": \"policy3\"}")'
);"
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVhmVtYLmWVVq7hWN4TbHDoWwa83VDHfK943JmNTcWWofb2AFQzP.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/group_policy.jsonld\", \"name\": \"policy4\"}")'
);"
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVh48LAR8M1jwY4RYiMiHnHFvw7vWWFerQy4RJVTqPjLPqm7sCs7.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/group_policy.jsonld\", \"name\": \"policy5\"}")'
);"
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVgAuyZF5uLrDx4FSkJAhKqJHcHCcQbAXKMwPcYnFBwRyWZKwu9x.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/group_member.jsonld\", \"name\": \"member1\"}")'
);"
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVgreSNsuF4yknA3tutdDVmy4berwCsECEWLVNuu7LnhTvQucAVf.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/group_member.jsonld\", \"name\": \"member2\"}")'
);"
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVhMZW81LXu1b75jdpPe5kHauFsgnwrYk541qPtAC8dempwGQtN2.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/group_member.jsonld\", \"name\": \"member3\"}")'
);"
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVgHZQcvE1TS6KzCQhpuhJ6dcyKFKnYkQM6J79uPLq7prArb2gPS.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/group_member.jsonld\", \"name\": \"member4\"}")'
);"

# seed database with data from chora genesis state (voucher metadata)
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVhHkkea9GmoZMeZnayM39RWJj1cjfhFUXg8FCDv2Qw3vYNAbwYS.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/voucher.jsonld\", \"name\": \"voucher1\"}")'
);"
psql postgres://postgres:password@localhost:5432/server -c "INSERT INTO data (iri, jsonld) VALUES (
  'chora:13toVhP9o1vzvSbLCTLeycFBxekJMW7oG6uL9xhcKbrMekfnAEgja57.rdf',
  '$(echo "{\"@context\": \"https://schema.chora.io/contexts/voucher.jsonld\", \"name\": \"voucher2\"}")'
);"

# stop background process
pg_ctl -D ${PGDATA} stop
