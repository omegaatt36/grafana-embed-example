--liquibase formatted sql

-- changeset create-tables:1
CREATE TABLE organizations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(256) NOT NULL
);

-- changeset create-tables:2
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(256) NOT NULL
);

-- changeset create-tables:3
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(256) NOT NULL
);

-- changeset create-tables:4
CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    organization_id INT NOT NULL,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    created_at INT NOT NULL,
    status ENUM('init', 'shipping', 'done') NOT NULL,
    FOREIGN KEY (organization_id) REFERENCES organizations(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- changeset insert-data:1
INSERT INTO organizations (name) VALUES ('Organization 1'), ('Organization 2'), ('Organization 3');

-- changeset insert-data:2
INSERT INTO users (name) VALUES
('User 1-1'), ('User 1-2'),
('User 2-1'), ('User 2-2'),
('User 3-1'), ('User 3-2');

-- changeset insert-data:3
INSERT INTO products (name) VALUES
('Product 1'), ('Product 2'), ('Product 3'), ('Product 4'), ('Product 5');

-- changeset insert-data:4
INSERT INTO orders (organization_id, user_id, product_id, created_at, status) VALUES
-- Organization 1, User 1-1
(1, 1, 1, 1704230400, 'init'),
(1, 1, 2, 1704316800, 'shipping'),
(1, 1, 3, 1704403200, 'done'),
(1, 1, 1, 1704489600, 'init'),
(1, 1, 2, 1704576000, 'shipping'),
-- Organization 1, User 1-2
(1, 2, 3, 1704662400, 'done'),
(1, 2, 4, 1704748800, 'init'),
(1, 2, 5, 1704835200, 'shipping'),
(1, 2, 3, 1704921600, 'done'),
(1, 2, 4, 1705008000, 'init'),
-- Organization 2, User 2-1
(2, 3, 1, 1705094400, 'shipping'),
(2, 3, 2, 1705180800, 'done'),
-- Organization 2, User 2-2
(2, 4, 3, 1705267200, 'init'),
(2, 4, 4, 1705353600, 'shipping'),
-- Organization 3, User 3-1
(3, 5, 5, 1705440000, 'done'),
(3, 5, 1, 1705526400, 'init'),
(3, 5, 2, 1705612800, 'shipping'),
-- Organization 3, User 3-2
(3, 6, 3, 1705699200, 'done'),
(3, 6, 4, 1705785600, 'init'),
(3, 6, 5, 1705872000, 'shipping');