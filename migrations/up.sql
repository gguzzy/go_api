create table products(
	id int primary key not null AUTO_INCREMENT,
    name varchar(60) not null,
    price float not null,
    description varchar(200),
    quantity_available int not null
)
