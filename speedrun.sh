#!/usr/bin/env bash

BASE=http://localhost:8080

echo "== ping =="
curl -s $BASE/ping | jq

echo "== create book =="
RESP=$(curl -s -X POST $BASE/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Speed Run","author":"Arya"}')

echo $RESP | jq

BOOK_ID=$(echo $RESP | jq -r .id)
echo "BOOK_ID=$BOOK_ID"

echo "== protected =="
curl -s $BASE/protected/books \
  -H "Authorization: Bearer secret-token" | jq

echo "== delete =="
curl -s -X DELETE $BASE/books/$BOOK_ID | jq