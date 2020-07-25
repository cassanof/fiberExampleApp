#!/bin/bash

curl -X POST -H "Content-Type: application/json" --data "{\"title\": \"Hacking\", \"author\": \"Jon Erickson\", \"rating\":5}" http://localhost:3000/api/v1/book
