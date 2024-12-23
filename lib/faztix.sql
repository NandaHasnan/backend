create table users (
id serial primary key,
first_name varchar(80),
last_name varchar(80),
phone_number varchar(18),
image varchar(255)
email varchar(80),
password varchar(300),
point int,
created_at timestamp default now(),
updated_at timestamp
)  

create table movie (
id serial primary key,
title varchar(100),
description varchar(255),
image varchar(255),
banner varchar(255),
genre varchar(100),
release_date date,
duration time,
director varchar(255),
cast varchar(255),
synopsis text,
location varchar(255),
date_time time,
created_at timestamp default now(),
updated_at timestamp
)

-- create table ticket (
-- id serial primary key,

-- created_at timestamp default now(),
-- updated_at timestamp
-- )

create table cinema (
id serial primary key,
cinema_name varchar(255),
cinema_logo varchar(255),
price INT,
location varchar(255),
date date,
time TIME,
created_at timestamp default now(),
updated_at timestamp
)

create table seats (
id serial primary key,
cinema_id int REFERENCES cinema(id),
num_seat int,
created_at timestamp default now(),
updated_at timestamp
) 

create table order (
id serial primary key,
user_id INT REFERENCES users(id)
movie_id INT REFERENCES movie(id),
cinema_id INT REFERENCES cinema(id), -- yang diambil price
seat VARCHAR(255),
total_payment INT,
created_at timestamp default now(),
updated_at timestamp
)

create table payment (
id serial primary key,
order_id INT REFERENCES order(id),
payment_methot VARCHAR(80),
payment_status VARCHAR(80),
created_at timestamp default now(),
updated_at timestamp
)
