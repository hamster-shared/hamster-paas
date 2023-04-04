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
    name varchar(50) NOT NULL,
    network varchar(50) NOT NULL,
    http_address varchar(255) NOT NULL,
    websocket_address varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);
INSERT INTO t_cl_rpc_chain (name, network, http_address, websocket_address)
VALUES (
        'ethereum',
        'mainnet',
        'https://ethereum.api.hamsternet.io',
        'wss://ethereum-ws.api.hamsternet.io'
    ),
    (
        'ethereum',
        'testnet-goerli',
        'https://ethereum-goerli.api.hamsternet.io',
        'wss://ethereum-goerli-ws.api.hamsternet.io'
    ),
    (
        'sui',
        'mainnet',
        'https://sui.api.hamsternet.io',
        'wss://sui-ws.api.hamsternet.io'
    ),
    (
        'avalanche',
        'mainnet',
        'https://avalanche.api.hamsternet.io',
        ''
    ),
    (
        'optimism',
        'mainnet',
        'https://optimism.api.hamsternet.io',
        'wss://optimism-ws.api.hamsternet.io'
    ),
    (
        'near',
        'mainnet',
        'https://near.api.hamsternet.io',
        ''
    ),
    (
        'starknet',
        'mainnet',
        'https://starkware.api.hamsternet.io',
        ''
    ),
    (
        'bsc',
        'mainnet',
        'https://bsc.api.hamsternet.io',
        'wss://bsc-ws.api.hamsternet.io'
    ),
    (
        'aptos',
        'mainnet',
        'https://aptos.api.hamsternet.io',
        ''
    ),
    (
        'polygon',
        'mainnet',
        'https://polygon.api.hamsternet.io',
        ''
    );
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