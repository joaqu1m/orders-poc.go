drop database if exists mydb;
create database mydb;
use mydb;

create table `user`(
	id INT PRIMARY KEY AUTO_INCREMENT,
    login VARCHAR(45),
    password VARCHAR(16)
);

create table `order`(
	id INT PRIMARY KEY AUTO_INCREMENT,
    customer_name VARCHAR(45),
    customer_email VARCHAR(45),
    status VARCHAR(45),
	source VARCHAR(45)
);

create table `item`(
	id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(45),
    price DECIMAL(7,2)
);

create table `order_item`(
	id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT,
    foreign key (order_id) references order(id),
    product_id INT,
    foreign key (product_id) references item(id),
    quantity INT
);

INSERT INTO `item`(name, price) values
('Pera', 2.90),
('Maçã', 3.90),
('Banana', 4.90),
('Abacaxi', 12.90);

INSERT INTO `order`(customer_name, customer_email, status, source) values
('joca', 'jocaa', 'ok', 'sla');

INSERT INTO user (login, password) values
('adm', '123');

select * from `item`;
select * from `order`;
select * from `user`;