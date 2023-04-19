alter table t_cl_deposit
    add column  address char(50) comment 'deposit wallet address';

alter table t_cl_subscription
    add column  balance decimal(18, 6) DEFAULT 0.0 comment 'subscription id balance';