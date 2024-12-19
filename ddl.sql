CREATE TABLE Users (
	user_id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	email VARCHAR(50) UNIQUE NOT NULL,
	password TEXT NOT NULL,
	deposit_amount DECIMAL(15, 2) NOT NULL DEFAULT 0
);

CREATE TABLE Cars (
	car_id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	stock_availability INTEGER NOT NULL,
	rental_costs DECIMAL(15, 2) NOT NULL,
	category VARCHAR(50) NOT NULL
);

CREATE TABLE Rentals (
	rental_id SERIAL PRIMARY KEY,
	user_id INTEGER REFERENCES Users(user_id),
	car_id INTEGER REFERENCES Cars(car_id),
	duration INTEGER NOT NULL,
	start_date DATE DEFAULT CURRENT_DATE,
	end_date DATE,
	total_costs DECIMAL(15, 2) NOT NULL,
	status VARCHAR(50) NOT NULL
);