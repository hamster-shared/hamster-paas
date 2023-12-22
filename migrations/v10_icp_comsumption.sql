create table t_icp_consumption
(
    id                int auto_increment,
    canister_id       varchar(50) default null,  
    cycles            decimal(10,2) default null,
    update_time       timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
) comment 'icp 每日消费';