DROP TABLE IF EXISTS t_cl_function_consumer_event;
CREATE TABLE IF NOT EXISTS t_cl_function_consumer_event (
     id BIGINT PRIMARY KEY AUTO_INCREMENT,
     request_id varchar(200),
     result     varchar(200),
     error_info varchar(200),
     created DATETIME NOT NULL
);