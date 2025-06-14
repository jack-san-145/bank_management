CREATE TABLE CustomersList (
    acc_no     INT PRIMARY KEY AUTO_INCREMENT,
    cus_name   VARCHAR(50) NOT NULL,
    password   VARCHAR(100) NOT NULL,
    address    VARCHAR(150),
    mobile_no  VARCHAR(15)
) AUTO_INCREMENT = 1000;