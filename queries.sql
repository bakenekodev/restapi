-- name: selectUsers
SELECT id, login, password, name, surname, telephone, car_id 
FROM users

-- name: selectUserByID
SELECT id, login, password, name, surname, telephone, car_id 
FROM users 
WHERE id = $1

-- name: insertUser
INSERT INTO users(login, password, name, surname, telephone, photo, car_id) 
VALUES($1, $2, $3, $4, $5, 1, $6)

-- name: updateUserByID
UPDATE users 
SET login = $1, password = $2, name = $3, surname = $4, telephone = $5, photo = 1, car_id = $6 
WHERE id = $7

-- name: deleteUserByID
DELETE FROM users 
WHERE id = $1 

-- name: selectUserID
SELECT id
FROM users
WHERE id = $1

-- name: selectCars
SELECT id, mark, model, year, seats
FROM cars

-- name: selectCarByID
SELECT id, mark, model, year, seats 
FROM cars 
WHERE id = $1

-- name: insertCar
INSERT INTO cars(mark, model, year, seats, driving_lic) 
VALUES ($1, $2, $3, $4, 1)

-- name: updateCarByID
UPDATE cars 
SET mark = $1, model = $2, year = $3, seats = $4, driving_lic = 1 
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