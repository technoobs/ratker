# The helper script that is used to manage database.
# -- Backup schema only: mysqldump --no-data -u someuser -p mydatabase
# -- Backup data only: mysqldump --no-create-info --no-create-db

#!bin/bash
docker exec ratkerDevMaria /usr/bin/mysqldump -u root --password=password RatkerDev > sample.sql