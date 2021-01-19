-- name: selectUsers
SELECT id, name, surname, telephone, car_id 
FROM users

-- name: selectUserByID
SELECT id, name, surname, telephone, current_lat, current_lng, car_id 
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

-- name: selectPasswordById
SELECT password
FROM credentials
WHERE id = $1

-- name: updatePassword
UPDATE credentials
SET password = $2
WHERE id = $1

-- name: checkLogin
SELECT login
FROM credentials
WHERE login = $1

-- name: insertLogin
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

-- name: selectRoutes
SELECT driver_id, points 
FROM drivers_trips

-- name: getDriversFunc
SELECT get_drivers($1, $2, $3)

-- name: upsetDriverRoute
INSERT INTO drivers_trips(driver_id, points) 
VALUES ($1, $2)
ON CONFLICT (driver_id)
DO
UPDATE SET points = $2

-- name: deleteRoute
DELETE FROM drivers_trips
WHERE driver_id = $1

-- name: acceptDriver
INSERT INTO passengers_trips(user_id, driver_id)
VALUES ($1, $2)

-- name: checkPassengers
SELECT user_id 
FROM passengers_trips 
WHERE driver_id = $1

-- name: declineDriver
DELETE FROM passengers_trips
WHERE user_id = $1

-- name: updatePos
UPDATE users
SET current_lat = $2, current_lng = $3
WHERE id = $1