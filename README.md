# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete users. Connected with a MySQL server.

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

# Request sample if user has a car
{
   "login": "name",
   "password": "pass",
   "name": "New",
   "surname": "User",
   "phone": "069000000",
   "car": {
      "mark":"Nissan",
      "model":"GTR",
      "year":"2020",
      "seats":"4"
   }
}
# If car is null
{
   "login":"name",
   "password":"pass",
   "name":"New",
   "surname":"User",
   "phone":"069000000",
   "car": null
}
```

### Update User
``` bash
PUT api/users/{id}

# Request sample
{
   "login":"name",
   "password":"pass",
   "name":"New",
   "surname":"User",
   "phone":"069000000",
   "car":{
      "mark":"Nissan",
      "model":"GTR",
      "year":"2020",
      "seats":"4"
   }
}

# If car is null
{
   "login":"name",
   "password":"pass",
   "name":"New",
   "surname":"User",
   "phone":"069000000",
   "car": null
}
```