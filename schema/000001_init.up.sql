CREATE TABLE users 
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);


CREATE TABLE orders
(
    id serial not null primary key,
    latitude float,
    longitude float,
    location varchar(255),
    typeOfOrder varchar(255) not null,
    typeOfAuto varchar(255) not null,
    numberOfClient varchar(255) not null,
    status varchar(255) not null,
    created_at timestamp not null default CURRENT_TIMESTAMP
);

create table executors_orders_history
(
    id serial not null primary key,
    order_id int not null,
    executor_number varchar(255) not null,
    created_at timestamp not null default CURRENT_TIMESTAMP
)