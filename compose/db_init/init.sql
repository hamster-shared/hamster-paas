DROP TABLE IF EXISTS t_cl_account;
CREATE TABLE IF NOT EXISTS t_cl_account (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    address varchar(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    app_id_index int NOT NULL
);
DROP TABLE IF EXISTS t_cl_app;
CREATE TABLE IF NOT EXISTS t_cl_app (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    app_id int NOT NULL,
    account varchar(50) NOT NULL,
    name varchar(50) NOT NULL,
    description varchar(255) NOT NULL,
    chain varchar(50) NOT NULL,
    network varchar(50) NOT NULL,
    api_key varchar(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    http_link varchar(100) NOT NULL,
    websocket_link varchar(100) NOT NULL,
    UNIQUE KEY (app_id, account)
);
DROP TABLE IF EXISTS t_cl_chain;
CREATE TABLE IF NOT EXISTS t_cl_chain (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(50) NOT NULL,
    network varchar(50) NOT NULL,
    http_address varchar(255) NOT NULL,
    websocket_address varchar(255) NOT NULL
);
INSERT INTO t_cl_chain (name, network, http_address, websocket_address)
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
        'starkware',
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
DROP TABLE IF EXISTS t_cl_code_examples;
CREATE TABLE IF NOT EXISTS t_cl_code_examples (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    chain varchar(50) NOT NULL,
    cli TEXT NOT NULL,
    javascript TEXT NOT NULL,
    python TEXT NOT NULL,
    go TEXT NOT NULL
);