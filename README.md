# Simple CRUD App w/ Gorilla/Mux, MariaDB
[![Go Report Card](https://goreportcard.com/badge/github.com/mrtkp9993/SimpleCRUDApp)](https://goreportcard.com/report/github.com/mrtkp9993/SimpleCRUDApp)
[![CodeFactor](https://www.codefactor.io/repository/github/mrtkp9993/simplecrudapp/badge)](https://www.codefactor.io/repository/github/mrtkp9993/simplecrudapp)

## Features 

Basic CRUD operations (Create-Read-Update-Delete).

## Database Scheme

Data table:

```sql
create table products
(
    id           int(11) unsigned auto_increment primary key,
    name         tinytext null,
    manufacturer tinytext null
);
```

You can generate data from http://filldb.info/.

Users table:

```sql
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` text NOT NULL,
  `saltedpassword` text NOT NULL,
  `salt` text NOT NULL,
  PRIMARY KEY (`id`)
);
```

````Password+Salt```` is encrypted with ``bcrypt``with 10 rounds and stored in ``saltedpassword``column.

## Example Requests

To get all entries from table:
```
curl --user user1:pass1 127.0.0.1:8000/api/products/list
```

To get an entry with `id` (where id equals 10):
```
curl --user user1:pass1 127.0.0.1:8000/api/products/10
```

To create an entry:
```
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"name": "ABC", "manufacturer": "ACME"}' \
	 --user user1:pass1 127.0.0.1:8000/api/products/new
```

To update an entry:
```
curl --request PUT \ 
     --data '{"name": "ABC", "manufacturer": "ACME"}' \ 
     --user user1:pass1 127.0.0.1:8000/api/products/11
```

To delete an entry:
```
curl --request DELETE --user user1:pass1 127.0.0.1:8000/api/products/10
```
