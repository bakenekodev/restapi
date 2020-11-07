# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete books. No database implementation yet

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
```

``` bash
go build | ./restapi
```

## Endpoints

### Get All Users
``` bash
GET api/users
```
### Get Single Users
``` bash
GET api/users/{id}
```

### Delete Users
``` bash
DELETE api/users/{id}
```

### Create User
``` bash
POST api/users

# Request sample
{
    "id": "2",
    "login": "numele",
    "password": "123",
    "name": "Marius",
    "surname": "Bitca",
    "phone": "069000000",
    "car": null
}
```

### Update User
``` bash
PUT api/users/{id}

# Request sample
{
    "login": "numele",
    "password": "123",
    "name": "Marius",
    "surname": "Bitca",
    "phone": "069000000",
    "car": null
}
```