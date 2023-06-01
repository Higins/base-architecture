# Blocks

#### BE start backend

```bash
./local.sh
```
Graphql play groud
http://localhost:8411/play

#### BE graphql resolvers

```bash
go get github.com/99designs/gqlgen@v0.17.12;go run github.com/99designs/gqlgen generate --verbose
```
#### BE db migration UP

```bash
cd migrations
go get github.com/pressly/goose/v3/cmd/goose;GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=blocks dbname=blocks sslmode=disable port=5443 host=localhost password=blocks" go run github.com/pressly/goose/v3/cmd/goose up
```
#### BE db migration DOWN

```bash
cd migrations
go get github.com/pressly/goose/v3/cmd/goose;GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=blocks dbname=blocks sslmode=disable port=5443 host=localhost password=blocks" go run github.com/pressly/goose/v3/cmd/goose down
```
#### BE Test data for login

```bash
Test data folder: migrations/test_data/user.sql
email: admin@admin.hu
password: admin
```
#### BE Test db

```bash
docker-compose up
```


#### Architecture

https://github.com/bxcodec/go-clean-arch
