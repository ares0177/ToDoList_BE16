create database ToDoList;

use ToDoList;

CREATE TABLE user (
  user_id int NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  email varchar(45) NOT NULL,
  phone_number varchar(14) NOT NULL,
  password varchar(45) NOT NULL,
  status int NOT NULL DEFAULT '1',
  created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id),
  UNIQUE KEY hp_unique (phone_number),
  UNIQUE KEY email_unique (email)
  );
 
 CREATE TABLE tasks (
  idTasks int NOT NULL AUTO_INCREMENT,
  user_id int NOT NULL,
  tasks_name varchar(45) NOT NULL,
  status int NOT NULL DEFAULT '1',
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (idTasks)
  );
 
 CREATE TABLE planning (
  idPlanning int NOT NULL AUTO_INCREMENT,
  user_id int NOT NULL,
  name varchar(45) NOT NULL,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (idPlanning)
  );
 
 CREATE TABLE planning_tasks (
  idTasks int NOT NULL,
  idPlanning int NOT NULL,
  PRIMARY KEY (idPlanning,idTasks)
  );
 
show tables;

select * from user;

select * from tasks;

select * from planning;

select * from planning_tasks;