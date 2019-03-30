# Simple CRUD App w/ Gorilla/Mux, MariaDB
[![Go Report Card](https://goreportcard.com/badge/github.com/mrtkp9993/SimpleCRUDApp)](https://goreportcard.com/report/github.com/mrtkp9993/SimpleCRUDApp)

## Features 

Basic CRUD operations (Create-Read-Update-Delete).

## Database Scheme

```sql
create table products
(
    id           int(11) unsigned auto_increment primary key,
    name         tinytext null,
    manufacturer tinytext null
);
```

You can generate data from http://filldb.info/.

## Example Requests

To get all entries from table:
```
curl 127.0.0.1:8000/api/products/list
```

To get an entry with `id` (where id equals 10):
```
curl 127.0.0.1:8000/api/products/10
```

To create an entry:
```
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"name": "ABC", "manufacturer": "ACME"}' \
	 127.0.0.1:8000/api/products/new
```

To update an entry:
```
curl --request PUT \ 
     --data '{"name": "ABC", "manufacturer": "ACME"}' \ 
     127.0.0.1:8000/api/products/11
```

To delete an entry:
```
curl --request DELETE 127.0.0.1:8000/api/products/10
```