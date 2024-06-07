
## Migration

go run main -migrate

## Execution API

go run main.go


## Create User

```sh
curl --location 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Higor Diego",
    "email": "Test Diego"
}'
```

## Get User

```sh
curl --location 'http://localhost:8080/users'
``` 