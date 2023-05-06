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
    image varchar(255) NOT NULL DEFAULT '',
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
        image,
        native_token,
        explorer_url,
        network_url,
        name,
        network,
        http_address,
        websocket_address
    )
VALUES (
        -- ethereum mainnet
        '1',
        'https://chain-logo.api.hamsternet.io/ethereum.svg',
        'ETH',
        'https://mainnet.etherscan.io',
        'https://mainnet.infura.io/v3/',
        'ethereum',
        'mainnet',
        'https://ethereum.api.hamsternet.io',
        'wss://ethereum-ws.api.hamsternet.io'
    ),
    (
        -- ethereum goerli
        '5',
        'https://chain-logo.api.hamsternet.io/ethereum.svg',
        'ETH',
        'https://goerli.etherscan.io',
        'https://rpc.ankr.com/eth_goerli',
        'ethereum',
        'testnet-goerli',
        'https://ethereum-goerli.api.hamsternet.io',
        'wss://ethereum-goerli-ws.api.hamsternet.io'
    ),
    (
        -- ethereum sepolia
        'aa36a7',
        'https://chain-logo.api.hamsternet.io/ethereum.svg',
        'ETH',
        'https://sepolia.etherscan.io',
        'https://rpc.ankr.com/eth_sepolia',
        'ethereum',
        'testnet-sepolia',
        'https://ethereum-sepolia.api.hamsternet.io',
        'wss://ethereum-sepolia-ws.api.hamsternet.io'
    ),
    (
        -- sui testnet 不属于 ethereum
        '0',
        'https://chain-logo.api.hamsternet.io/sui.svg',
        'SUI',
        'https://explorer.sui.io/',
        '',
        'sui',
        'testnet',
        'https://sui-testnet.api.hamsternet.io',
        ''
    ),
    (
        'a86a',
        'https://chain-logo.api.hamsternet.io/avalanche.svg',
        'AVAX',
        'https://snowtrace.io/',
        'https://api.avax.network/ext/bc/C/rpc',
        'avalanche',
        'mainnet',
        'https://avalanche.api.hamsternet.io',
        'https://avalanche-ws.api.hamsternet.io'
    ),
    -- (
    --     'a869',
    --     'AVAX',
    --     'https://explorer.avax-test.network/',
    --     'https://api.avax-test.network/ext/bc/C/rpc',
    --     'avalanche',
    --     'testnet-fuji',
    --     'https://avalanche-fuji.api.hamsternet.io',
    --     'https://avalanche-fuji-ws.api.hamsternet.io'
    -- ),
    (
        'a',
        'https://chain-logo.api.hamsternet.io/optimism.svg',
        'ETH',
        'https://optimistic.etherscan.io/',
        'https://mainnet.optimism.io',
        'optimism',
        'mainnet',
        'https://optimism.api.hamsternet.io',
        'https://optimism-ws.api.hamsternet.io'
    ),
    (
        '1a4',
        'https://chain-logo.api.hamsternet.io/optimism.svg',
        'ETH',
        'https://goerli-explorer.optimism.io',
        'https://goerli.optimism.io',
        'optimism',
        'testnet-goerli',
        'https://optimism-goerli.api.hamsternet.io',
        'https://optimism-goerli-ws.api.hamsternet.io'
    ),
    (
        '4E454152',
        'https://chain-logo.api.hamsternet.io/near.svg',
        'NEAR',
        'https://explorer.mainnet.aurora.dev/',
        'https://mainnet.aurora.dev',
        'near',
        'mainnet',
        'https://near.api.hamsternet.io',
        ''
    ),
    -- (
    --     '0',
    --     'NEAR',
    --     'https://explorer.mainnet.near.org/',
    --     'https://rpc.mainnet.near.org',
    --     'near',
    --     'testnet-testnet',
    --     'https://near-testnet.api.hamsternet.io',
    --     ''
    -- ),
    (
        -- aptos 不属于 eth 类型
        '0',
        'https://chain-logo.api.hamsternet.io/aptos.svg',
        'APT',
        'https://explorer.aptoslabs.com/?network=mainnet',
        'https://fullnode.mainnet.aptoslabs.com/v1',
        'aptos',
        'mainnet',
        'https://aptos.api.hamsternet.io',
        ''
    ),
    -- (
    --     '0',
    --     'unknown',
    --     'https://explorer.mainnet.near.org/',
    --     'https://rpc.mainnet.near.org',
    --     'aptos',
    --     'testnt-testnet',
    --     'https://aptos-testnet.api.hamsternet.io',
    --     ''
    -- ),
    -- (
    --     '0',
    --     'unknown',
    --     'https://explorer.mainnet.near.org/',
    --     'https://rpc.mainnet.near.org',
    --     'aptos',
    --     'testnt-devnet',
    --     'https://aptos-devnet.api.hamsternet.io',
    --     ''
    -- ),
    (
        -- starknet 不属于 eth 类型
        '0',
        'https://chain-logo.api.hamsternet.io/starknet.svg',
        'ETH',
        'https://starkscan.co/',
        '',
        'starknet',
        'mainnet',
        'https://starknet.api.hamsternet.io',
        ''
    ),
    (
        '0',
        'https://chain-logo.api.hamsternet.io/starknet.svg',
        'ETH',
        'https://testnet.starkscan.co/',
        '',
        'starknet',
        'testnet-goerli',
        'https://starknet-goerli.api.hamsternet.io',
        ''
    ),
    (
        '89',
        'https://chain-logo.api.hamsternet.io/polygon.svg',
        'MATIC',
        'https://polygonscan.com',
        'https://polygon-rpc.com/',
        'polygon',
        'mainnet',
        'https://polygon.api.hamsternet.io',
        'https://polygon-ws.api.hamsternet.io'
    ),
    (
        '13881',
        'https://chain-logo.api.hamsternet.io/polygon.svg',
        'MATIC',
        'https://mumbai.polygonscan.com/',
        'https://rpc-mumbai.maticvigil.com',
        'polygon',
        'testnet-mumbai',
        'https://polygon-mumbai.api.hamsternet.io',
        'https://polygon-mumbai-ws.api.hamsternet.io'
    ),
    (
        '501',
        'https://chain-logo.api.hamsternet.io/hamster.svg',
        'HA',
        'https://hamsternet.io',
        'https://rpc-moonbeam.hamster.newtouch.com',
        'hamster',
        'testnet-moonbeam',
        '',
        ''
    ),
    (
        -- ton 不属于 ethereum
        '0',
        'https://chain-logo.api.hamsternet.io/ton.svg',
        'TON',
        'https://tonscan.org/',
        '',
        'ton',
        'mainnet',
        'https://ton.api.hamsternet.io',
        ''
    ),
    (
        -- arbitrum
        '42161',
        'https://chain-logo.api.hamsternet.io/arbitrum.svg',
        'ETH',
        'https://arbiscan.io/',
        'https://arb1.arbitrum.io/rpc',
        'arbitrum',
        'mainnet',
        'https://arbitrum.api.hamsternet.io',
        'wss://arbitrum-ws.api.hamsternet.io'
    ),
    (
        -- arbitrum goerli
        '421613',
        'https://chain-logo.api.hamsternet.io/arbitrum.svg',
        'ETH',
        'https://goerli.arbiscan.io/',
        'https://goerli-rollup.arbitrum.io/rpc',
        'arbitrum',
        'testnet-goerli',
        'https://arbitrum-goerli.api.hamsternet.io',
        'wss://arbitrum-goerli-ws.api.hamsternet.io'
    ),
    (
        -- irisnet mainnet
        '1a20',
        'https://chain-logo.api.hamsternet.io/irisnet.svg',
        'ERIS',
        'https://irishub.iobscan.io/',
        'https://evmrpc.irishub-1.irisnet.org',
        'irisnet',
        'mainnet',
        'https://irisnet.api.hamsternet.io',
        ''
    ),
    (
        -- irisnet nyancat
        '4130',
        'https://chain-logo.api.hamsternet.io/irisnet.svg',
        'ERIS',
        'https://nyancat.iobscan.io/',
        'https://rpc.nyancat.irisnet.org/',
        'irisnet',
        'testnet-nyancat',
        'https://irisnet-nyancat.api.hamsternet.io',
        ''
    )
    ;
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