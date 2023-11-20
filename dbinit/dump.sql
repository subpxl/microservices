-- Create a database
CREATE DATABASE IF NOT EXISTS todoservice;

-- Use the created database
USE todoservice;

create table if not exists Todo(
		id SERIAL PRIMARY KEY,
		Todo varchar(255)
	);

