# Go Demo App

This is a demo app with two REST APIs:
1. Get /data: To get all records in the database table `data`.
2. Post /data: To add a new record to the database table `data`.

### Build App Docker Image

Use the following command to build the app docker image:

```bash
docker build -t myapp .
```

### Start MySQL and App

You can edit `docker-compose.yml`, 
replace volumes `~/data/mysql` with your local folder and set environment `MYSQL_ROOT_PASSWORD` with your password.
Use the following command to start MySQL and this app.

```bash
docker-compose up -d
```

### Create MySQL Schema

Change to `sql/` directory, run the following command to create database schema (only need do once):

```bash
docker exec -i mysql mysql -uroot -proot < schema.sql
```

### Test APIs

You can use cURL to test REST APIs.

1. Get all data:

```bash
curl -i -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/data
```

2. Add new data:

```bash
curl -i -X POST -H "Content-Type: application/json" -d '{"id":"Taipei", "location":{"lat":25.032969, "lng":121.565414}}' http://127.0.0.1:8080/data
```

### Stop & Remove MySQL and App

```bash
docker-compose down
```
