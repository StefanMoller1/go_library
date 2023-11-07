# Golang Project

## About the project

Library management project in golang

## Getting started

Starting the project
```
$ docker compose --profile app up -d --build
```

Starting the database only and running the project locally
```
$ docker compose --profile psql up -d
$ go run cmd/main.go
```

## Endpoints
Select All Books
```
$ curl 0.0.0.0:8000/api/v1/library
```

Select All Books
```
$ curl 0.0.0.0:8000/api/v1/<id>
```

Create Book
```
curl  -H 'Content-Type: application/json' -X POST 0.0.0.0:8000/api/v1/library -d '{"name": "Test Book", "description":"Some book to Test With", "author":"Test Author"}'
```

Update Book
```
curl  -H 'Content-Type: application/json' -X PUT 0.0.0.0:8000/api/v1/library/<id> -d '{"name": "Test Book", "description":"Some book to Test With", "author":"New Author"}'
```

Delete Book
```
curl -X DELETE 0.0.0.0:8000/api/v1/library/af033fcb-53c6-4fe5-a24f-101c5a085757
```