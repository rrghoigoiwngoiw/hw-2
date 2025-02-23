create database dbonlinestore
\c dbonlinestore;

create user db_user with password 'password';

grant connect on database dbonlinestore to db_user;

grant select, insert, update, delete on all tables in schema public to db_user; --?????

create table Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

create table Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
);


create table Orders (
    id SERIAL PRIMARY KEY,
    user_id int not NULL,
    order_date date not NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE OrderProducts (
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES Orders(id),
    FOREIGN KEY (product_id) REFERENCES Products(id)
);


INSERT INTO Users (username, password, email) 
VALUES ('john_doe', 'password123', 'john.doe@example.com');

insert into Products (name, price)
values ('слон', '100000000');


update Users
set name = 'John doe', email = 'john.doe@example.com'
where id = 1;

update Products
set price = 120000000
where id = 1;

delete from Users
where id = 1;

delete from Products
where id = 1;


insert into Orders (user_id, order_date, total_amount)
values (1, '2022-01-01', 100000000);

insert into OrderProducts (order_id, product_id, quantity)
values (1, 1, 1);

delete from OrderProducts
where order_id = 1;

delete from Orders
where id = 1;

select * from Products;

select * from Orders;

select * from OrderProducts;

select * from Users;

select o.id, o.order_date, o.total_amount
from Orders o
join Users u on o.user_id = u.id
where u.id = 1;

SELECT 
    u.id AS user_id,
    u.name AS user_name,
    SUM(o.total_amount) AS total_spent,
    AVG(p.price) AS average_product_price
FROM Users u
JOIN Orders o ON u.id = o.user_id
JOIN OrderProducts op ON o.id = op.order_id
JOIN Products p ON op.product_id = p.id
WHERE u.id = 1
GROUP BY u.id, u.name;

create index idx_users_email on Users(email)

create index idx_orders_id on Orders(order_id)
create index idx_order_date on Orders(order_date)

create index idx_product_name on Products(name)
create index idx_product_price on Products(price)