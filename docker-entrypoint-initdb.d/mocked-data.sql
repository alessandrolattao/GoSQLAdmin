-- Create tables
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS reviews (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    comment TEXT,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Populate the users table
INSERT INTO users (name, email) VALUES
('User1', 'user1@example.com'),
('User2', 'user2@example.com'),
('User3', 'user3@example.com'),
('User4', 'user4@example.com'),
('User5', 'user5@example.com'),
('User6', 'user6@example.com'),
('User7', 'user7@example.com'),
('User8', 'user8@example.com'),
('User9', 'user9@example.com'),
('User10', 'user10@example.com'),
('User11', 'user11@example.com'),
('User12', 'user12@example.com'),
('User13', 'user13@example.com'),
('User14', 'user14@example.com'),
('User15', 'user15@example.com'),
('User16', 'user16@example.com'),
('User17', 'user17@example.com'),
('User18', 'user18@example.com'),
('User19', 'user19@example.com'),
('User20', 'user20@example.com'),
('User21', 'user21@example.com'),
('User22', 'user22@example.com'),
('User23', 'user23@example.com'),
('User24', 'user24@example.com'),
('User25', 'user25@example.com'),
('User26', 'user26@example.com'),
('User27', 'user27@example.com'),
('User28', 'user28@example.com'),
('User29', 'user29@example.com'),
('User30', 'user30@example.com'),
('User31', 'user31@example.com'),
('User32', 'user32@example.com'),
('User33', 'user33@example.com'),
('User34', 'user34@example.com'),
('User35', 'user35@example.com'),
('User36', 'user36@example.com'),
('User37', 'user37@example.com'),
('User38', 'user38@example.com'),
('User39', 'user39@example.com'),
('User40', 'user40@example.com'),
('User41', 'user41@example.com'),
('User42', 'user42@example.com'),
('User43', 'user43@example.com'),
('User44', 'user44@example.com'),
('User45', 'user45@example.com'),
('User46', 'user46@example.com'),
('User47', 'user47@example.com'),
('User48', 'user48@example.com'),
('User49', 'user49@example.com'),
('User50', 'user50@example.com'),
('User51', 'user51@example.com'),
('User52', 'user52@example.com'),
('User53', 'user53@example.com'),
('User54', 'user54@example.com'),
('User55', 'user55@example.com'),
('User56', 'user56@example.com'),
('User57', 'user57@example.com'),
('User58', 'user58@example.com'),
('User59', 'user59@example.com'),
('User60', 'user60@example.com');

-- Populate the products table
INSERT INTO products (name, price, stock) VALUES
('Product1', 10.00, 100),
('Product2', 15.50, 50),
('Product3', 20.00, 30),
('Product4', 25.00, 10),
('Product5', 5.00, 500),
('Product6', 8.99, 120),
('Product7', 12.49, 60),
('Product8', 14.99, 40),
('Product9', 22.00, 25),
('Product10', 30.00, 15),
('Product11', 35.50, 5),
('Product12', 40.00, 8),
('Product13', 45.00, 12),
('Product14', 50.00, 20),
('Product15', 55.00, 18),
('Product16', 60.00, 9),
('Product17', 65.00, 7),
('Product18', 70.00, 6),
('Product19', 75.00, 5),
('Product20', 80.00, 4),
('Product21', 85.00, 3),
('Product22', 90.00, 2),
('Product23', 95.00, 1),
('Product24', 100.00, 50),
('Product25', 105.00, 45),
('Product26', 110.00, 40),
('Product27', 115.00, 35),
('Product28', 120.00, 30),
('Product29', 125.00, 25),
('Product30', 130.00, 20),
('Product31', 135.00, 15),
('Product32', 140.00, 10),
('Product33', 145.00, 8),
('Product34', 150.00, 6),
('Product35', 155.00, 4),
('Product36', 160.00, 2),
('Product37', 165.00, 12),
('Product38', 170.00, 10),
('Product39', 175.00, 8),
('Product40', 180.00, 6),
('Product41', 185.00, 4),
('Product42', 190.00, 2),
('Product43', 195.00, 30),
('Product44', 200.00, 28),
('Product45', 205.00, 26),
('Product46', 210.00, 24),
('Product47', 215.00, 22),
('Product48', 220.00, 20),
('Product49', 225.00, 18),
('Product50', 230.00, 16),
('Product51', 235.00, 14),
('Product52', 240.00, 12),
('Product53', 245.00, 10),
('Product54', 250.00, 8),
('Product55', 255.00, 6),
('Product56', 260.00, 4),
('Product57', 265.00, 2),
('Product58', 270.00, 12),
('Product59', 275.00, 10),
('Product60', 280.00, 8);


-- Populate the orders table
INSERT INTO orders (user_id, product_id, quantity, total_price) VALUES
(1, 1, 2, 20.00),
(2, 2, 1, 15.50),
(3, 3, 3, 60.00),
(4, 4, 1, 25.00),
(5, 5, 10, 50.00),
(6, 6, 5, 44.95),
(7, 7, 4, 49.96),
(8, 8, 2, 29.98),
(9, 9, 1, 22.00),
(10, 10, 3, 90.00),
(11, 11, 1, 35.50),
(12, 12, 2, 80.00),
(13, 13, 1, 45.00),
(14, 14, 3, 150.00),
(15, 15, 2, 110.00),
(16, 16, 4, 240.00),
(17, 17, 1, 65.00),
(18, 18, 2, 140.00),
(19, 19, 1, 75.00),
(20, 20, 5, 400.00),
(21, 21, 3, 255.00),
(22, 22, 2, 180.00),
(23, 23, 1, 95.00),
(24, 24, 3, 300.00),
(25, 25, 4, 420.00),
(26, 26, 5, 550.00),
(27, 27, 2, 230.00),
(28, 28, 3, 360.00),
(29, 29, 1, 125.00),
(30, 30, 6, 780.00),
(31, 31, 3, 405.00),
(32, 32, 2, 280.00),
(33, 33, 4, 580.00),
(34, 34, 5, 750.00),
(35, 35, 1, 155.00),
(36, 36, 2, 320.00),
(37, 37, 3, 495.00),
(38, 38, 1, 170.00),
(39, 39, 2, 350.00),
(40, 40, 4, 720.00),
(41, 41, 5, 925.00),
(42, 42, 1, 190.00),
(43, 43, 3, 585.00),
(44, 44, 2, 400.00),
(45, 45, 1, 205.00),
(46, 46, 4, 840.00),
(47, 47, 3, 645.00),
(48, 48, 2, 440.00),
(49, 49, 5, 1125.00),
(50, 50, 6, 1380.00),
(51, 51, 2, 470.00),
(52, 52, 1, 240.00),
(53, 53, 3, 735.00),
(54, 54, 4, 1000.00),
(55, 55, 5, 1275.00),
(56, 56, 2, 520.00),
(57, 57, 1, 265.00),
(58, 58, 4, 1080.00),
(59, 59, 3, 885.00),
(60, 60, 2, 560.00);

-- Populate the reviews table
INSERT INTO reviews (product_id, user_id, rating, comment) VALUES
(1, 1, 5, 'Excellent quality!'),
(2, 2, 4, 'Very useful product'),
(3, 3, 3, 'Average, but works'),
(4, 4, 2, 'Could be better'),
(5, 5, 5, 'Outstanding value'),
(6, 6, 4, 'Very practical'),
(7, 7, 3, 'It’s okay'),
(8, 8, 2, 'Not great'),
(9, 9, 5, 'Highly recommend'),
(10, 10, 4, 'Good for the price'),
(11, 11, 5, 'Perfect!'),
(12, 12, 4, 'Well made'),
(13, 13, 3, 'Just fine'),
(14, 14, 2, 'Below expectations'),
(15, 15, 5, 'Top-notch'),
(16, 16, 4, 'Pretty good'),
(17, 17, 3, 'Acceptable'),
(18, 18, 2, 'Not worth it'),
(19, 19, 5, 'Fantastic'),
(20, 20, 4, 'Good overall'),
(21, 21, 3, 'It’s fine'),
(22, 22, 2, 'Needs improvement'),
(23, 23, 5, 'Excellent design'),
(24, 24, 4, 'Well packaged'),
(25, 25, 3, 'Average quality'),
(26, 26, 2, 'Not great'),
(27, 27, 5, 'Superb value'),
(28, 28, 4, 'Very efficient'),
(29, 29, 3, 'Okay for the price'),
(30, 30, 2, 'Disappointing'),
(31, 31, 5, 'Highly recommend'),
(32, 32, 4, 'Pretty solid'),
(33, 33, 3, 'Meets expectations'),
(34, 34, 2, 'Could improve'),
(35, 35, 5, 'Amazing!'),
(36, 36, 4, 'Really good'),
(37, 37, 3, 'Not bad'),
(38, 38, 2, 'Not worth the money'),
(39, 39, 5, 'Absolutely love it'),
(40, 40, 4, 'Decent purchase'),
(41, 41, 3, 'Fine product'),
(42, 42, 2, 'Wouldn’t buy again'),
(43, 43, 5, 'Highly recommend!'),
(44, 44, 4, 'Good build quality'),
(45, 45, 3, 'It’s okay'),
(46, 46, 2, 'Could do better'),
(47, 47, 5, 'Excellent!'),
(48, 48, 4, 'Nice product'),
(49, 49, 3, 'Satisfied'),
(50, 50, 2, 'Not impressed'),
(51, 51, 5, 'Fantastic product!'),
(52, 52, 4, 'Good overall value'),
(53, 53, 3, 'Just okay'),
(54, 54, 2, 'Disappointing build'),
(55, 55, 5, 'Highly recommend!'),
(56, 56, 4, 'Great buy'),
(57, 57, 3, 'Meh'),
(58, 58, 2, 'Not what I hoped for'),
(59, 59, 5, 'Outstanding quality!'),
(60, 60, 4, 'Happy with the purchase');

