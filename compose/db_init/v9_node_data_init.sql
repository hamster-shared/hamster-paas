INSERT INTO t_cl_rpc_node_resource_standard (id,chain_protocol,cpu,memory,disk,region,cost_per_month)
VALUES (1,'Ethereum','4','32','3000','US East',530.00),
       (2,'Sui','8','32','300','US East',657.20),
       (3,'Near','4','8','1300','US East',318.00),
       (4,'Starknet','2','4','500','US East',95.40),
       (5,'Aptos','8','32','2000','US East',551.20),
       (6,'Optimism','4','16','200','US East',164.30),
       (7,'Avalanche','8','16','1000','US East',413.40);

INSERT INTO t_cl_black_height (id, black_height, event_type) VALUES (1, 36902587, 'Transfer');