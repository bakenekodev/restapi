-- name: selectUsers
SELECT id, login, password, name, surname, telephone, car_id 
FROM users

-- name: selectUserByID
SELECT id, login, password, name, surname, telephone, car_id 
FROM users 
WHERE id = ?

-- name: insertUser
INSERT INTO users(login, password, name, surname, telephone, photo, car_id) 
VALUES(?, ?, ?, ?, ?, 1, ?)

-- name: updateUserByID
UPDATE users 
SET login = ?, password = ?, name = ?, surname = ?, telephone = ?, photo = 1, car_id = ? 
WHERE id = ?

-- name: deleteUserByID
DELETE FROM users 
WHERE id = ?  

-- name: selectCarByID
SELECT id, mark, model, year, seats 
FROM cars 
WHERE id = ?

-- name: insertCar
INSERT INTO cars(mark, model, year, seats, driving_lic) 
VALUES (?, ?, ?, ?, 1)

-- name: updateCarByID
UPDATE cars 
SET mark = ?, model = ?, year = ?, seats = ?, driving_lic = 1 
WHERE id = ?

-- name: deleteCarByID
DELETE FROM cars 
WHERE id = ?

-- name: selectCarIdByID
SELECT car_id 
FROM users 
WHERE id = ?

-- name: selectID
SELECT id
FROM users
WHERE id = ?
LIMIT 1

-- name: insertDriverRoute
INSERT INTO dirvers_trips(driver_id, start_lat, start_lng, end_lat, end_lng, start_time, end_time) 
VALUES (?, ?, ?, ?, ?, ?, ?)