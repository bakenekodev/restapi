# GO Lang REST API

> RESTful API to handle requests from the HitchHike application. Connected with a MySQL server from Google Cloud.


## Endpoints

### User

#### Get All Users
``` bash
GET api/users
```
#### Get Single User
``` bash
GET api/users/{id}
```

#### Delete User
``` bash
DELETE api/users/{id}
```

#### Create User
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

#### Update User
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

## Car

#### Get All Cars
``` bash
GET api/cars
```
#### Get Single Car
``` bash
GET api/cars/{id}
```

#### Delete Car
``` bash
DELETE api/cars/{id}
```

#### Create Car
``` bash
POST api/cars

# Request sample
{
   "mark":"Nissan",
   "model":"GTR",
   "year":"2020",
   "seats":"4"
}
```

#### Update Car
``` bash
PUT api/cars/{id}

# Request sample
{
   "mark":"Nissan",
   "model":"GTR",
   "year":"2020",
   "seats":"4"
}
```