
-- Insert Products
INSERT INTO products (id, name, price, stock) VALUES
(1, 'Espresso', 25000, 100),
(2, 'Latte', 30000, 120),
(3, 'Cappuccino', 35000, 80),
(4, 'Mocha', 40000, 60),
(5, 'Americano', 20000, 150);

-- Insert Raw Materials
INSERT INTO raw_materials (id, name, unit, description) VALUES
(1, 'Coffee Beans', 'kg', 'High-quality beans'),
(2, 'Milk', 'L', 'Fresh milk'),
(3, 'Sugar', 'kg', 'Granulated sugar');

-- Insert Transactions
INSERT INTO transactions (id, user_id, total_amount, created_at) VALUES
(1, 1, 75000, '2025-01-01'),
(2, 2, 105000, '2025-01-02');

-- Insert Transaction Items
INSERT INTO transaction_items (transaction_id, product_id, quantity) VALUES
(1, 1, 2),
(1, 2, 1),
(2, 2, 2),
(2, 3, 1);
