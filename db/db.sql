CREATE DATABASE retailer_db;

USE retailer_db;

CREATE TABLE products (
    id VARCHAR(10) PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    quantity INT NOT NULL
);

CREATE TABLE orders (
    id VARCHAR(10) PRIMARY KEY,
    customer_id VARCHAR(10) NOT NULL,
    product_id VARCHAR(10) NOT NULL,
    quantity INT NOT NULL,
    status ENUM('order placed', 'processed', 'failed') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
show databases
SHOW TABLes;
SELECT * FROM products;
DESC products;
ALTER TABLE products MODIFY id INT PRIMARY KEY AUTO_INCREMENT;
ALTER TABLE products DROP PRIMARY KEY;
ALTER TABLE products MODIFY id INT NOT NULL AUTO_INCREMENT PRIMARY KEY;

CREATE TABLE products_new (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_name LONGTEXT,
    price DOUBLE,
    quantity BIGINT
);
INSERT INTO products_new (product_name, price, quantity)
SELECT product_name, price, quantity FROM products;



DROP TABLE products;
ALTER TABLE products_new RENAME TO products;

DESC products;
ALTER TABLE products MODIFY id INT NOT NULL AUTO_INCREMENT PRIMARY KEY;

ALTER TABLE products MODIFY id INT NOT NULL AUTO_INCREMENT PRIMARY KEY;
 SHOW CREATE TABLE products;
 
 ALTER TABLE orders MODIFY COLUMN id VARCHAR(20);






