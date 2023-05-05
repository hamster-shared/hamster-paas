<div align="center">
    <img src="https://socialify.git.ci/hamster-shared/hamster-paas/image?description=1&descriptionEditable=One-stop%20Toolkit%20and%20Middleware%20Platform%20for%20Web3.0%20Developers&font=KoHo&logo=https%3A%2F%2Fhamsternet.io%2F_nuxt%2Flogo.668de5a2.png&owner=1&pattern=Floating%20Cogs&theme=Auto" width="640" height="320" alt="logo" />

# <a href="https://develop.alpha.hamsternet.io/chainlink/dashboard">Hamster PaaS</a>

[![Discord](https://badgen.net/badge/icon/discord?icon=discord&label)](https://discord.gg/qMWUvs7jkV)
[![Telegram](https://badgen.net/badge/icon/telegram?icon=telegram&label)](https://t.me/hamsternetio)
[![Twitter](https://badgen.net/badge/icon/twitter?icon=twitter&label)](https://twitter.com/Hamsternetio)

_One-stop Toolkit and Middleware Platform for Web3.0 Developers_

</div>

## Getting Started

To get started with Hamster PaaS, follow the steps below:

1. **Prepare your environment file:** You will need to create a `.env` file in the root directory of the project. You can find an example file named `.env.example` in the repository. Copy and rename this file to `.env`, and update the necessary fields according to your configuration.

Here is an example of the contents of the `.env` file:

```
# GIN_MODE=release
PORT=9898
MEILI_SEARCH="http://127.0.0.1:7700"
NGINX_LOG_PATH=./compose/nginx/log/access.log

MYSQL_DATABASE=aline
MYSQL_USER=aline
MYSQL_PASSWORD=changeme
MYSQL_ROOT_PASSWORD=changeme
MYSQL_HOST=127.0.0.1
MYSQL_PORT=3306

ALINE_DB_USER=root
ALINE_DB_PASSWORD=changeme
ALINE_DB_HOST=127.0.0.1
ALINE_DB_PORT=30303
ALINE_DB_NAME=aline

FROM_EMAIL=
EMAIL_PASSWORD=
```

2. **Start Docker containers:** Run the following command in the terminal to start the required Docker containers:

```bash
docker-compose up -d
```

This command will start all the necessary services, such as database and search engine, in the background.

3. **Run the application:** Finally, execute the following command to start the Hamster PaaS application:

```bash
go run .
```

The application should now be up and running. You can access the Hamster PaaS dashboard by navigating to `http://localhost:9898` in your browser.

## About Hamster

Hamster is aiming to build the one-stop infrastructure developer toolkits for Web3.0. It defines itself as a development, operation and maintenance DevOps service platform, providing a set of development tools as well as O&M tools, empowering projects in Web3.0 to improve their coding and delivery speed, quality and efficiency, as well as product reliability & safety.

With Hamster, developers or project teams realize the development, verification and O&M stages of their blockchain projects in an automatic, standardized and tooled approach: from contract template of multiple chains, contract/frontend code build, security check, contract deployment to the contract operation and maintenance.

Together with its developer toolkits, Hamster offers the RPC service and decentralized computing power network service when users choose to deploy their contracts via Hamster.

At the same time, the contract security check part within the developer toolkits is offered separately to its to-C customers, who could check their contracts to avoid potential security risks.

## Contributors

This project exists thanks to all the people who contribute.

 <a href="https://github.com/hamster-shared/hamster-paas/contributors">
  <img src="https://contrib.rocks/image?repo=hamster-shared/hamster-paas" />
 </a>

## License

[MIT](LICENSE)
