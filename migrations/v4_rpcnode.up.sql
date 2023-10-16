create table t_rpc_node (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(50) not null comment 'rpc 节点名',
    user_id int,
    chain_protocol  tinyint default 0 comment '链协议（evm,其他）',
    status tinyint not null default 0 comment '链状态',
    public_ip varchar(40)  comment '公网ip',
    region  varchar(20) comment '机器所在区域',
    launch_time timestamp   comment '启动时间',
    resource varchar(100) comment '资源规格',
    chain_version   varchar(20) comment '部署链的版本',
    next_payment_date timestamp null comment '下一次支付时间',
    payment_per_month decimal(10,2) comment '每月支付金额',
    remaining_sync_time varchar(20) comment '剩余同步时间',
    current_height  int(11) comment '当前区块高度',
    block_time  varchar(20) comment '平均出块时间',
    http_endpoint   varchar(75) comment 'http 请求地址',
    websocket_endpoint varchar(75) comment  'websocket 请求地址',
    created datetime NOT NULL
) comment 'rpc 节点表';

create table t_rpc_node_resource_standard(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    chain_protocol  tinyint default 0 comment '链协议（evm,其他）',
    cpu varchar(20),
    memory varchar(20),
    disk varchar(20),
    cost_per_month decimal(10,2) comment '每月花费'
) comment 'rpc 节点资源月租收费表';

create table t_user_charge_account(
    user_id int primary key ,
    seed varchar(64),
    address varchar(64)  comment '收费账户地址'
) comment '用户收费账户表';

create table t_order(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    order_id varchar(50) not null,
    user_id int not null,
    order_time timestamp not null,
    order_type tinyint not null  default 0 comment '订单类型',
    resource_type   varchar(50) comment '资源类型',
    status tinyint not null  default 0 comment '资源状态',
    chain varchar(25) comment '链类型',
    amount decimal(10,2) comment '总价',
    pay_address varchar(50) comment '支付地址',
    receive_address varchar(50) comment '收款地址',
    pay_tx varchar(64) comment '交易事务号',
    index (order_id),
    index (user_id)
);
