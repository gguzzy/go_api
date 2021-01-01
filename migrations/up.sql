create table products(
	id int primary key not null AUTO_INCREMENT,
    name varchar(60) not null,
    price float not null,
    description varchar(200),
    quantity_available int not null
)


create table inventory (
	storeId int,
    productId int
)


create table shop (
	id int primary key AUTO_INCREMENT,
    name varchar(60),
    address varchar(60),
    time date
)
