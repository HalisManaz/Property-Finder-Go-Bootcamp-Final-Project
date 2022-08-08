CREATE SCHEMA db;
USE db;
DROP TABLE IF EXISTS products;
CREATE TABLE products(ID VARCHAR(255) PRIMARY KEY, Name VARCHAR(255), TaxRate INT, Price FLOAT);
INSERT INTO products(ID, Name, TaxRate, Price) VALUES("1","Chocolate",18,5.0);
INSERT INTO products(ID, Name, TaxRate, Price) VALUES("2","Ice Cream", 18, 12.0);
INSERT INTO products(ID, Name, TaxRate, Price) VALUES("3","Water",1,8.0);
INSERT INTO products(ID, Name, TaxRate, Price) VALUES("4","Soda",8,4.5);
INSERT INTO products(ID, Name, TaxRate, Price) VALUES("5","Bread",1,4.25);
INSERT INTO products(ID, Name, TaxRate, Price) VALUES("6","Lemonade",8,11.25);
INSERT INTO products(ID, Name, TaxRate, Price) VALUES("7","Banana (kg)",8,24);
INSERT INTO products(ID, Name, TaxRate, Price) VALUES("8","Apple (kg)",18,15);
INSERT INTO products(ID, Name, TaxRate, Price) VALUES("9","Soap",1,11.5);
INSERT INTO products(ID, Name, TaxRate, Price) VALUES("10","Tootbrush",1,17);

DROP TABLE IF EXISTS basket;
create table basket
(
    ID      varchar(255) null,
    Name    varchar(255) null,
    TaxRate float        null,
    Price   float        null,
    Amount  int          null,
    OrderID int          not null
);

DROP TABLE IF EXISTS receipts;
create table receipts
(
    TotalPrice float null,
    TaxAmount  float null,
    Discount   float null,
    AmountDue  float null,
    OrderID    int   not null
        primary key
);

DROP TABLE IF EXISTS userandorder;
create table userandorder
(
    ID      varchar(255) null,
    OrderID int          not null
        primary key
);

DROP TABLE IF EXISTS users;

create table users
(
    ID           varchar(255) not null
        primary key,
    UserName     varchar(255) null,
    Type         varchar(255) null,
    Streak       int          null,
    MonthlyTotal float        null
);