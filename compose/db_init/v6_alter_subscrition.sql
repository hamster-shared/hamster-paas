alter table t_cl_subscription
    add column error_message varchar(255);

alter table t_cl_deposit
    add column error_message varchar(255);

alter table t_cl_consumer
    add column error_message varchar(255);