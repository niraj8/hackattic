#!/usr/bin/env bash

curl -s https://hackattic.com/challenges/backup_restore/problem?access_token=${HACKATTIC_ACCESS_TOKEN} | jq -r .dump | base64 -d > dump.zip
zcat < dump.zip > dump
sleep 2
psql -c "drop database hackattic_backup_restore"
psql -c "create database hackattic_backup_restore"
psql -d hackattic_backup_restore < dump
psql -d hackattic_backup_restore -c "select json_agg(ssn) from criminal_records where status='alive'" | tail -n 3 | head -n 1 | jq -n '{alive_ssns: input}' > alive_ssns.json
curl -X POST -H "Content-Type: application/json" -d @alive_ssns.json https://hackattic.com/challenges/backup_restore/solve?access_token=${HACKATTIC_ACCESS_TOKEN}

# cleanup
rm dump.zip dump alive_ssns.json
