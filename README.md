# GolangPostgresCourse
 Geek Brains Golang Postgres 

 To initialize Database using migrate/migrate package in docker follow these steps:

 1. Find your postgres database ip with this command:
 docker inspect postgres | grep "IPAddress"
 
 The result should be something like this:
            "SecondaryIPAddresses": null,
            "IPAddress": "172.17.0.2",
                    "IPAddress": "172.17.0.2"

2. Run docker with the following command using the IP that you got from step 1:
docker run -v "$(pwd)/migrations":/migrations migrate/migrate -path=/migrations/ -database "postgres://bituser:bit@172.17.0.2:5432/bitmedb?sslmode=disable" -verbose up


