DROP TABLE IF EXISTS t_cl_rpc_account;
CREATE TABLE IF NOT EXISTS t_cl_rpc_account (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    address varchar(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);
DROP TABLE IF EXISTS t_cl_rpc_app;
CREATE TABLE IF NOT EXISTS t_cl_rpc_app (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    app_id int NOT NULL,
    account varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    chain varchar(50) NOT NULL,
    network varchar(50) NOT NULL,
    api_key varchar(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    http_link varchar(100) NOT NULL,
    websocket_link varchar(100) NOT NULL,
    UNIQUE KEY (account, name)
);
DROP TABLE IF EXISTS t_cl_rpc_chain;
CREATE TABLE IF NOT EXISTS t_cl_rpc_chain (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    chain_id varchar(255) NOT NULL,
    native_token varchar(255) NOT NULL,
    explorer_url varchar(255) NOT NULL,
    network_url varchar(255) NOT NULL,
    name varchar(50) NOT NULL,
    network varchar(50) NOT NULL,
    http_address varchar(255) NOT NULL,
    websocket_address varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);
INSERT INTO t_cl_rpc_chain (
        chain_id,
        native_token,
        explorer_url,
        network_url,
        name,
        network,
        http_address,
        websocket_address
    )
VALUES (
        '1',
        'ETH',
        'https://mainnet.etherscan.io',
        'https://mainnet.infura.io/v3/',
        'ethereum',
        'mainnet',
        -- 'https://ethereum.api.hamsternet.io',
        'http://54.69.42.237:9912',
        'wss://ethereum-ws.api.hamsternet.io'
    ),
    (
        '13881',
        'MATIC',
        'https://explorer-mumbai.maticvigil.com/',
        'https://rpc-mumbai.maticvigil.com',
        'polygon',
        'testnet-mumbai',
        'http://54.69.42.237:9912',
        ''
    ),
    (
        '501',
        'Hamster',
        'https://hamsternet.io',
        'https://rpc-moonbeam.hamster.newtouch.com',
        'hamster',
        'testnet-moonbeam',
        'http://54.69.42.237:9912',
        ''
    ),
    (
        '5',
        'GETH',
        'https://goerli.etherscan.io',
        'https://goerli.infura.io/v3/',
        'ethereum',
        'testnet-goerli',
        -- 'https://ethereum-goerli.api.hamsternet.io',
        'http://54.69.42.237:9912',
        'wss://ethereum-goerli-ws.api.hamsternet.io'
    ),
    (
        'aa36a7',
        'Ether',
        'https://sepolia.etherscan.io',
        'https://sepolia.infura.io/v3/',
        'ethereum',
        'testnet-sepolia',
        'http://54.69.42.237:9912',
        ''
    ),
    (
        'a86a',
        'AVAX',
        'https://snowtrace.io/',
        'https://api.avax.network/ext/bc/C/rpc',
        'avalanche',
        'mainnet',
        -- 'https://avalanche.api.hamsternet.io',
        'http://54.69.42.237:9912',
        ''
    ),
    (
        'a869',
        'AVAX',
        'https://explorer.avax-test.network/',
        'https://api.avax-test.network/ext/bc/C/rpc',
        'avalanche',
        'testnet-goerli',
        -- 'https://ethereum-goerli.api.hamsternet.io',
        'http://54.69.42.237:9912',
        'wss://ethereum-goerli-ws.api.hamsternet.io'
    ),
    (
        'a',
        'ETH',
        'https://optimistic.etherscan.io/',
        'https://mainnet.optimism.io',
        'optimism',
        'mainnet',
        -- 'https://ethereum-goerli.api.hamsternet.io',
        'http://54.69.42.237:9912',
        -- 'wss://ethereum-goerli-ws.api.hamsternet.io'
        ''
    ),
    (
        '1a4',
        'ETH',
        'https://optimistic.etherscan.io/',
        'https://mainnet.optimism.io',
        'optimism',
        'mainnet',
        -- 'https://ethereum-goerli.api.hamsternet.io',
        'http://54.69.42.237:9912',
        -- 'wss://ethereum-goerli-ws.api.hamsternet.io'
        ''
    ),
    (
        '1a4',
        'ETH',
        'https://optimistic.etherscan.io/',
        'https://goerli.optimism.io',
        'optimism',
        'testnet-goerli',
        -- 'https://ethereum-goerli.api.hamsternet.io',
        'http://54.69.42.237:9912',
        -- 'wss://ethereum-goerli-ws.api.hamsternet.io'
        ''
    ),
    (
        '1a4',
        'ETH',
        'https://optimistic.etherscan.io/',
        'https://goerli.optimism.io',
        'optimism',
        'testnet-goerli',
        -- 'https://ethereum-goerli.api.hamsternet.io',
        'http://54.69.42.237:9912',
        -- 'wss://ethereum-goerli-ws.api.hamsternet.io'
        ''
    ),
    (
        '0',
        'NEAR',
        'https://explorer.mainnet.near.org/',
        'https://rpc.mainnet.near.org',
        'near',
        'mainnet',
        -- 'https://ethereum-goerli.api.hamsternet.io',
        'http://54.69.42.237:9912',
        -- 'wss://ethereum-goerli-ws.api.hamsternet.io'
        ''
    );
-- (
--     0,
--     'NEAR',
--     'https://explorer.testnet.near.org/',
--     'https://rpc.testnet.near.org',
--     'near',
--     'testnet-testnet',
--     -- 'https://ethereum-goerli.api.hamsternet.io',
--     'http://54.69.42.237:9912',
--     -- 'wss://ethereum-goerli-ws.api.hamsternet.io'
--     ''
-- );
DROP TABLE IF EXISTS t_cl_rpc_code_examples;
CREATE TABLE IF NOT EXISTS t_cl_rpc_code_examples (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    chain varchar(50) NOT NULL,
    cli TEXT NOT NULL,
    javascript TEXT NOT NULL,
    python TEXT NOT NULL,
    go TEXT NOT NULL
);
DROP TABLE IF EXISTS t_cl_oracle_request_event;
CREATE TABLE IF NOT EXISTS t_cl_oracle_request_event (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    chain varchar(255) NOT NULL,
    network varchar(255) NOT NULL,
    transaction_hash varchar(255) NOT NULL,
    requesting_contract varchar(255) NOT NULL,
    request_initiator varchar(255) NOT NULL,
    subscription_id int NOT NULL,
    subscription_owner varchar(255) NOT NULL,
    block_number int NOT NULL,
    tx_index int NOT NULL,
    block_hash varchar(255) NOT NULL,
    _index int NOT NULL,
    removed tinyint(1) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);