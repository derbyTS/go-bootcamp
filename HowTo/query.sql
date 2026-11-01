
-- CREATE DATABASE IF NOT EXISTS school;
-- USE school;
-- CREATE TABLE IF NOT EXISTS teachers (
-- 	id INT AUTO_INCREMENT PRIMARY KEY,
-- 	first_name VARCHAR(255) NOT NULL,
-- 	last_name VARCHAR(255) NOT NULL,
-- 	email VARCHAR(255) NOT NULL UNIQUE,
-- 	class VARCHAR(255) NOT NULL,
-- 	subject VARCHAR(255) NOT NULL,
-- 	INDEX(email)
-- )AUTO_INCREMENT=100;

USE school;

-- SHOW TABLES;

-- SELECT * from teachers;

-- SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1;


CREATE INDEX idx_class ON teachers(class);

CREATE TABLE IF NOT EXISTS students (
	id INT AUTO_INCREMENT PRIMARY KEY,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	class VARCHAR(255) NOT NULL,
	INDEX(email),
	FOREIGN KEY (class) REFERENCES teachers(class)
)AUTO_INCREMENT=100;

