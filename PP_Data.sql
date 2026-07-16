USE clothing_ecommerce_db;

-- Akun Admin & Customer (Password dummy: admin123 & user123)
INSERT INTO users (email, password, role) VALUES 
('admin@store.com', 'admin123', 'admin'),
('imam@store.com', 'user123', 'customer');

INSERT INTO profiles (user_id, full_name, phone, address) VALUES 
(1, 'Admin Gudang Fesyen', '081111111111', 'Jakarta Selatan, DKI Jakarta'),
(2, 'Imam Pembeli', '082222222222', 'Bandung, Jawa Barat');

-- Kategori Pakaian
INSERT INTO categories (name) VALUES 
('Atasan Pria'), ('Outerwear'), ('Celana');

-- Katalog Produk
INSERT INTO products (category_id, sku, name, description, size, color, price, stock) VALUES 
(1, 'TSH-BLK-L', 'Kaos Polos Oversize', 'Kaos katun bambu super adem 24s', 'L', 'Hitam', 120000.00, 20),
(1, 'TSH-WHT-M', 'Kaos Polos Heavyweight', 'Kaos bahan tebal tidak terawang 16s', 'M', 'Putih', 135000.00, 15),
(2, 'JKT-DNM-XL', 'Jaket Denim Vintage', 'Jaket jeans bertekstur klasik klasik', 'XL', 'Biru Navy', 350000.00, 5),
(3, 'CHN-KRM-L', 'Celana Chino Slimfit', 'Celana formal kasual bahan stretch', 'L', 'Krem', 210000.00, 10);