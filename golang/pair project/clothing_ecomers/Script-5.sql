create database pproject;
use pproject


/* Daftar table 
 	- master table
	 	1. products
	 	2. categories
	- support table
		1. users
		2. profiles
		3. orders
		4. order details
		5. payments
 */


/* ------------ */
/* master table */
/* ------------ */
 
  create table if not exists categories (
 	categoryid INT auto_increment primary key,
 	name VARCHAR(255) null
 );
  
 create table if not exists products (
 	productid INT auto_increment primary key,
 	categoryid INT null,
 	sku VARCHAR(30) null,
 	name VARCHAR(255) null,
 /* description text null,  */ -- kondisional bisa dipakai atau tidak
 	size VARCHAR(10) null,
 	color VARCHAR(30) null,
 	price DECIMAL(10,2) null,
 	stock INT null,
 	created_at TIMESTAMP null,
 	update_at TIMESTAMP null
 );
 

 /* ------------- */
 /* table support */
 /* ------------- */
 
 create table if not exists users (
 	userid INT auto_increment primary key,
 	email VARCHAR(50) UNIQUE,
 	password VARCHAR(50) null
 /* role VARCHAR(50) null, */ -- tidak perlu
 /* creat_at timestamp null */ -- kondisional
);

create table if not exists profiles (
 	profileid INT auto_increment primary key,
 	userid INT null,
 	fullname VARCHAR(255) null,
 	phone VARCHAR(255) null,
 	address TEXT null
 );
 create table if not exists orders (
 	orderid INT auto_increment primary key,
 	userid INT null,
 	order_date TIMESTAMP null,
 	status VARCHAR(20) null,
 	total_amount DECIMAL(10,2) null,
 	created_at TIMESTAMP null
 );
 
 create table if not exists order_details (
 	order_detail_id INT auto_increment primary key,
 	orderid INT null,
 	productid INT null,
 	quantity INT null,
 	price_at_buy DECIMAL(10,2) null,
 	subtotal DECIMAL(10,2) null
 );
 
 create table if not exists payment (
 	paymentid INT auto_increment primary key,
 	orderid INT null,
 	payment_method VARCHAR(30) null,
 	payment_status VARCHAR(30) null,
 	amount DECIMAL(10,2) null,
 	paid_at TIMESTAMP
 );
 
 /* --------------- */
 /* menotasikan FK  */
 /* --------------- */
 
ALTER TABLE profiles 
ADD CONSTRAINT fk_profile_userid
FOREIGN KEY (userid)
REFERENCES users(userid);

ALTER TABLE products 
ADD CONSTRAINT fk_categoryid
FOREIGN KEY (categoryid)
REFERENCES categories(categoryid);

ALTER TABLE orders 
ADD CONSTRAINT fk_userid
FOREIGN KEY (userid)
REFERENCES users(userid);

ALTER TABLE order_details 
ADD CONSTRAINT fk_order_details
FOREIGN KEY (orderid)
REFERENCES products(productid);

ALTER TABLE order_details 
ADD CONSTRAINT fk_productid
FOREIGN KEY (productid)
REFERENCES products(productid);

ALTER TABLE payment
ADD CONSTRAINT fk_orderid
FOREIGN KEY (orderid)
REFERENCES orders(orderid);


/* -------------------- */
/* INSERT DATA MASTER   */
/* -------------------- */

INSERT INTO categories (categoryid, name) VALUES
(1, 'T-Shirt'),
(2, 'Jacket'),
(3, 'Pants'),
(4, 'Shoes'),
(5, 'Accessories');


INSERT INTO products (productid, categoryid, sku, name, /*description,*/ size, color, price, stock, created_at, update_at) VALUES
(1, 1, 'TSH-001', 'T-Shirt Affends', /*'Kaos katun premium dengan bahan lembut.',*/ 'M', 'Black', 129000, 50, NOW(), NOW()),
(2, 1, 'TSH-002', 'T-Shirt Cole', /*'Kaos oversized untuk gaya kasual.',*/ 'L', 'White', 149000, 35, NOW(), NOW()),
(3, 2, 'JKT-001', 'Jaket Crooz', /*'Hoodie fleece hangat dan nyaman.',*/ 'XL', 'Grey', 299000, 20, NOW(), NOW()),
(4, 2, 'JKT-002', 'Denim Jacket', /*'Jaket denim klasik untuk pria dan wanita.',*/ 'L', 'Blue', 399000, 18, NOW(), NOW()),
(5, 3, 'PNT-001', 'Celana Jeans Levis', /*'Celana jeans slim fit stretch.',*/ '32', 'Dark Blue',349000, 25, NOW(), NOW()),
(6, 3, 'PNT-002', 'Cargo Pants', /*'Celana cargo dengan banyak kantong.',*/ '34', 'Army Green', 289000, 22, NOW(), NOW()),
(7, 4, 'SHO-001', 'Sepatu Nike', /*'Sepatu sneakers ringan untuk aktivitas sehari-hari.',*/ '42', 'White', 599000, 15, NOW(), NOW()),
(8, 4, 'SHO-002', 'Kaos Kaki Levis', /*'Sepatu kanvas casual.',*/ '41', 'Black', 359000, 30, NOW(), NOW()),
(9, 5, 'ACC-001', 'Sabuk Nike', /*'Ikat pinggang kulit asli.',*/ 'All Size', 'Brown', 179000, 40, NOW(), NOW()),
(10, 5, 'ACC-002', 'Topi Nike', /*'Topi baseball dengan desain minimalis.',*/ 'All Size', 'Navy', 99000, 45, NOW(), NOW());


-- menampilkan sekali order (belum sampai payment)
SELECT
    pr.fullname,
    pd.name,
    od.quantity,
    od.subtotal
FROM orders o
JOIN profiles pr
ON o.userid = pr.userid
JOIN order_details od
ON o.orderid = od.orderid
JOIN products pd
ON od.productid = pd.productid;


-- menampilkan sekali order sampai payment
SELECT
    u.userid,
    p.fullname,
    u.email,
    o.orderid,
    o.order_date,
    pr.name AS product_name,
    pr.sku,
    od.quantity,
    od.price_at_buy,
    od.subtotal,
    o.total_amount,
    pay.payment_method,
    pay.payment_status,
    pay.amount,
    pay.paid_at
FROM users u
JOIN profiles p
    ON u.userid = p.userid
JOIN orders o
    ON u.userid = o.userid
JOIN order_details od
    ON o.orderid = od.orderid
JOIN products pr
    ON od.productid = pr.productid
JOIN payment pay
    ON o.orderid = pay.orderid
ORDER BY o.orderid, pr.name;

-- menampilkan daftar riwayat
SELECT
    pr.fullname,
    pd.name,
    od.quantity,
    pay.payment_method,
    pay.payment_status,
    pay.amount
FROM orders o
JOIN profiles pr
ON o.userid = pr.userid
JOIN order_details od
ON o.orderid = od.orderid
JOIN products pd
ON od.productid = pd.productid
JOIN payment pay
ON o.orderid = pay.orderid;

