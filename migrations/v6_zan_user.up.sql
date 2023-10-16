create table t_zan_user
(
    id           int auto_increment,
    user_id      bigint      not null comment '用户id',
    access_token varchar(64) null comment 'zan 平台token',
    created      timestamp   null comment '创建时间',
    constraint t_zan_user_pk
        primary key (id)
)
    comment 'zan授权用户表';

create index t_zan_user_user_id_index
    on t_zan_user (user_id);

