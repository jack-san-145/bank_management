CREATE TABLE Transaction (
    sender_id        INT NOT NULL,
    receiver_id      INT NOT NULL,
    sender_name      VARCHAR(50) NOT NULL,
    receiver_name    VARCHAR(50) NOT NULL,
    amount           INT NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);