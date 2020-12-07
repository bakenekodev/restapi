-- name: selectUsers
SELECT id, name, surname, telephone, car_id 
FROM users

-- name: selectUserByID
SELECT id, name, surname, telephone, car_id 
FROM users 
WHERE id = $1

-- name: insertUser
INSERT INTO users(id, name, surname, telephone, car_id) 
VALUES($1, $2, $3, $4, $5)
RETURNING id

-- name: updateUserByID
UPDATE users 
SET name = $1, surname = $2, telephone = $3, car_id = $4 
WHERE id = $5

-- name: deleteUserByID
DELETE FROM users 
WHERE id = $1 

-- name: selectUserID
SELECT id
FROM users
WHERE id = $1

-- name: selectPassword
SELECT password, id 
FROM credentials 
WHERE login = $1

--name: checkLogin
SELECT login
FROM credentials
WHERE login = $1

--name: insertLogin
INSERT INTO credentials(login, password)
VALUES ($1, $2)
RETURNING id

-- name: selectCars
SELECT id, mark, model, year, seats
FROM cars

-- name: selectCarByID
SELECT id, mark, model, year, seats 
FROM cars 
WHERE id = $1

-- name: insertCar
INSERT INTO cars(mark, model, year, seats) 
VALUES ($1, $2, $3, $4)
RETURNING id

-- name: updateCarByID
UPDATE cars 
SET mark = $1, model = $2, year = $3, seats = $4 
WHERE id = $5

-- name: deleteCarByID
DELETE FROM cars 
WHERE id = $1

-- name: selectCarIdByUser
SELECT car_id 
FROM users 
WHERE id = $1

-- name: selectCarID
SELECT id
FROM cars
WHERE id = $1

-- name: insertDriverRoute
INSERT INTO dirvers_trips(driver_id, start_lat, start_lng, end_lat, end_lng, start_time, end_time) 
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id