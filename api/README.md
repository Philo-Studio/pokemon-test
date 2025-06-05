# Pokemon list API

## Install & Setup

### Start DB

```sh
$ docker run -p "5789:5432" --name "poke-list-db" -h "localhost" -e POSTGRES_PASSWORD="postgres" -e POSTGRES_USER="postgres" -e POSTGRES_DB="postgres" -d postgres:15.4
```

### Load db init state

```sh
$ psql -h localhost -U postgres -d postgres -p 5789 -f ./init_db.sql -W

# Then type password postgres when prompted
```

```sh
$ make
```

### Set .env

```sh
$ cp ./.env.example ./.env
```

## Dev

```sh
$ make dev
```
