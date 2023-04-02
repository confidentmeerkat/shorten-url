# urlshort #

urlshort is a URL Shortener with a web page for regular users, API for developers, and can redirect to the original URL.

## How to use ##

You can use urlshort in 2 ways:

1. Using docker

2. Using go

### Using docker ###

For using urlshort with docker, you should have installed `docker` and `docker-compose`.

After that clone urlshort repository:

```
https://github.com/barahouei/urlshort.git
```
 Then in the urlshort directory, use the following command:

```
docker-compose up
```

### Using go ###

For using urlshort with go, you should have installed `go` and `postgresql`.

After that, follow the below instructions:

Clone urlshort repository:

```
https://github.com/barahouei/urlshort.git
```

Go to the urlshort directory.

Create a database with a custom name like `shortener` then import `init.sql` file from `configs` directory to the database.

Set the following environment variables in your machine:

Environment Variable | Value
---------------------|------------
POSTGRES_HOST        | localhost
POSTGRES_PORT        | 5432
POSTGRES_USER        | your postgres user like `postgres`
POSTGRES_PASSWORD    | your postgres password
POSTGRES_DB          | your created database name like `shortener`
POSTGRES_SSL_MODE    | disable
SHORTENER_DOMAIN     | http://localhost:8080

Install dependencies:

```
go mod download && go mod verify
```

Now you can either run or build and run:

Run:

```
go run main.go
```

Build and run:

```
go build -o urlshort && ./urlshort
```