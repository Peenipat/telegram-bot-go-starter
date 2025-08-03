-- 1. เมนูอาหาร
CREATE TABLE menu_items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    category TEXT,
    price NUMERIC(10, 2) NOT NULL,
    is_available BOOLEAN DEFAULT TRUE
);

-- 2. คำสั่งซื้อ/ออเดอร์
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    total_amount NUMERIC(10, 2) NOT NULL,
    table_number INT,
    payment_method TEXT,
    status TEXT DEFAULT 'completed'  -- หรือใช้ enum ภายหลังก็ได้
);

-- 3. รายการอาหารในแต่ละออเดอร์
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    menu_item_id INT REFERENCES menu_items(id),
    quantity INT NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

-- 4. การจองโต๊ะ
CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    customer_name TEXT NOT NULL,
    phone TEXT,
    table_number INT,
    reserved_time TIMESTAMP NOT NULL,
    status TEXT DEFAULT 'pending'  -- pending, confirmed, cancelled
);

-- 5. ความคิดเห็นจากลูกค้า
CREATE TABLE feedbacks (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE SET NULL,
    customer_name TEXT,
    comment TEXT,
    rating INT CHECK (rating BETWEEN 1 AND 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 6. วัตถุดิบในระบบ
CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    unit TEXT NOT NULL,  -- เช่น 'kg', 'pcs'
    stock_level NUMERIC(10, 2) DEFAULT 0
);

-- 7. การเคลื่อนไหวของสต็อกวัตถุดิบ
CREATE TABLE stock_movements (
    id SERIAL PRIMARY KEY,
    ingredient_id INT REFERENCES ingredients(id) ON DELETE CASCADE,
    amount NUMERIC(10, 2) NOT NULL,
    type TEXT CHECK (type IN ('in', 'out')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
