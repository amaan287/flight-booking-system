### command to run postgres locally using docker

- docker run --name postgres-container -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=postgres -p 5432:5432 -d postgres
