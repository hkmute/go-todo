#!/bin/bash

# This script is used to run the migration scripts in the database container.

source .env
~/go/bin/tern migrate --config ./packages/db/migrations/tern.conf --migrations ./packages/db/migrations
