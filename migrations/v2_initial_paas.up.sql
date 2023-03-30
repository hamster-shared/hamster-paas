DROP TABLE IF EXISTS t_user_middleware;
CREATE TABLE IF NOT EXISTS t_user_middleware (
    user_id BIGINT NOT NULL PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    chain VARCHAR(50) NOT NULL,
    network VARCHAR(50) NOT NULL,
    created DATETIME NOT NULL
    );

DROP TABLE IF EXISTS t_cl_subscription;
CREATE TABLE IF NOT EXISTS t_cl_subscription (
    id BIGINT PRIMARY KEY,
    subscription_id BIGINT,
    name VARCHAR(50) NOT NULL,
    created DATETIME NOT NULL,
    chain varchar(50) NOT NULL,
    network varchar(50) NOT NULL,
    consumers TINYINT(3) NOT NULL,
    balance decimal(18, 6) NOT NULL,
    user_id bigint NOT NULL,
    admin char(42) NOT NULL,
    transaction_tx char(42) NOT NULL,
    status char(20) NOT NULL,
    INDEX(user_id)
    );

DROP TABLE IF EXISTS t_cl_consumer;
CREATE TABLE IF NOT EXISTS t_cl_consumer (
    id BIGINT NOT NULL PRIMARY KEY,
    subscription_id BIGINT NOT NULL,
    created datetime NOT NULL,
    consumer_address char(42) NOT NULL,
    transaction_tx char(42) NOT NULL,
    status char(20) NOT NULL,
    user_id BIGINT NOT NULL,
    INDEX(user_id)
    );

DROP TABLE IF EXISTS t_cl_request_template;
CREATE TABLE IF NOT EXISTS t_cl_request_template (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    created datetime NOT NULL,
    script text NOT NULL,
    author VARCHAR(20) NOT NULL,
    description VARCHAR(255) NOT NULL
    );

DROP TABLE IF EXISTS t_cl_request;
CREATE TABLE IF NOT EXISTS t_cl_request (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    created datetime NOT NULL,
    script text NOT NULL,
    user_id BIGINT NOT NULL,
    INDEX(user_id)
    );

DROP TABLE IF EXISTS t_cl_request_execute;
CREATE TABLE IF NOT EXISTS t_cl_request_execute (
    id BIGINT NOT NULL PRIMARY KEY ,
    subscription_id BIGINT NOT NULL,
    consumer_address char(42) NOT NULL,
    secretsloction tinyint NOT NULL,
    args varchar(255) NOT NULL,
    transaction_tx char(42) NULL,
    status char(20),
    user_id BIGINT NOT NULL,
    created DATETIME NOT NULL,
    INDEX(user_id)
    );

DROP TABLE IF EXISTS t_cl_deposit;
CREATE TABLE IF NOT EXISTS t_cl_deposit (
    id BIGINT NOT NULL PRIMARY KEY ,
    created datetime NOT NULL,
    request_name char(50) NOT NULL,
    consumer_address char(42) NOT NULL,
    amount decimal(18, 6) NOT NULL,
    transaction_tx char(42) NOT NULL,
    status char(20),
    user_id BIGINT NOT NULL,
    INDEX(user_id)
    );