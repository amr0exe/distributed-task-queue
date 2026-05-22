# Helper Commands

``` bash
docker compose up -d
```

``` bash
docker compose run --rm goose goose up
docker compose run --rm goose goose down
```

``` bash
# fish shell
set -x DBSTRING "postgres://tskQ:password@localhost:5432/tsk_db?sslmode=disable"

#bash shell
export DBSTRING="postgres://tskQ:password@localhost:5432/tsk_db?sslmode=disable"
```

## Requests

## Create Task

``` bash
curl -X POST http://localhost:8080/task \
-H "Content-Type: application/json" \
-d '{"title":"learn distributed systems"}'
```

## Get all Task

``` bash
curl http://localhost:8080/all
```

## Delete Task

``` bash
curl -X DELETE http://localhost:8080/task/id_here
```

## Update Task

``` bash
curl -X PUT http://localhost:8080/task/id_here \
-H "Content-Type: application/json" \
-d '{"title":"halo systems", "is_completed": true}'
```
