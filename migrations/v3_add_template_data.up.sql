INSERT INTO t_cl_request_template (id, name, created, script, author, description)
VALUES (
           1,
           'Fetch Discord Upvote Data',
           '2023-03-30',
           '// This function retrieves the number of upvotes a Discord member has received in the past 24 hours using the Discord API.

   const getDiscordUpvotes = async (memberId, apiKey, guildId, channelId, timeRangeMs) => {
       const endpoint = ''https://discord.com/api/v9''
       const timeRangeSec = Math.round(timeRangeMs / 1000)
       const time24HoursAgo = Math.round((Date.now() - timeRangeMs) / 1000)
       const headers = {
           ''Authorization'': `Bot ${apiKey}`,
           ''Content-Type'': ''application/json''
       }
       const config = {
           method: ''GET'',
           headers: headers,
           url: `${endpoint}/guilds/${guildId}/audit-logs?limit=100&user_id=${memberId}&before=${time24HoursAgo}&action_type=MESSAGE_DELETE`
       }
       const response = await Functions.makeHttpRequest(config)
       if (response.error) {
           throw new Error(response.response.data.message)
       }
       const auditLogs = response.data.audit_log_entries
       let upvotes = 0
       for (let i = 0; i < auditLogs.length; i++) {
           const log = auditLogs[i]
           if (log.action_type === 72 && log.target_id === channelId && log.created_at >= time24HoursAgo - timeRangeSec) {
               upvotes++
           }
       }
       return Functions.encodeUint256(upvotes)
   }
   ',
           'Sam Demaree',
           'This function retrieves the number of upvotes a Discord member has received in the past 24 hours. *Note: ChatGPT was used to demonstrate that non-developers can also participate.'
       ),
       (
           2,
           'US election results from AP (Associated Press) API',
           '2023-03-30',
           '// Chainlink Function to get election results from AP (Associated Press) API. Date and API key are the only required parameters

   const getReportingUnit = (reportingUnits, statePostal) => {
     const level = statePostal === ''US'' ? ''national'' : ''state''
     const reportingUnit = reportingUnits.find((ru) => ru.level === level)
     if (!reportingUnit) {
       throw new Error(''Cannot find reporting unit'')
     }
     return reportingUnit
   }

   const getReportingUnitWinner = (reportingUnit) => {
     for (const candidate of reportingUnit.candidates) {
       if (candidate.winner === ''X'') {
         return candidate
       }
     }
     throw new Error(''Candidate not found'')
   }


   const date = args[0] // The date of the election formatted as YYYY-MM-DD
   const statePostal = args[1] // The state''s two-letter code e.g CO. `US` to get the results of a nationwide election
   const raceID = args[2] // AP-assigned race ID. Should be used with `statePostal`
   const raceType = args[3] || ''G'' // The race type the election is for. The race type can be `D(Dem Primary)`, `R(GOP Primary)`, `G(General)`, `E(Dem Caucus)`, `S(GOP Caucus)`, `X(Open Primary or special use cases)`
   const resultsType = args[4] || ''L'' // The type of results to return. `L` for live results, `T` for test results

   if (!secrets.apikey) {
     throw new Error(''Missing AP API key'')
   }

   const params = {
     level: statePostal === ''US'' ? ''national'' : ''state'',
     raceTypeID: raceType,
     format: ''json'',
     winner: ''X'',
     resultsType: resultsType,
     apikey: secrets.apikey,
   }

   if ((statePostal && !raceID) || (!statePostal && raceID)) {
     throw new Error(''Both statePostal and raceID are required if one is provided'')
   }

   if (statePostal) {
     params.statePostal = statePostal
   }

   if (raceID) {
     params.raceID = raceID
   }

   const config = {
     url: `https://api.ap.org/v3/elections/${date}`,
     params
   }

   const response = await Functions.makeHttpRequest(config)

   const races = response.data.races
   if (races.length === 0) {
     throw new Error(''Could not find any races'')
   }
   if (races.length > 1) {
     throw new Error(''Finding the winner from multiple races is not supported'')
   }

   const race = races[0]
   const reportingUnit = getReportingUnit(race.reportingUnits, statePostal)
   const raceWinner = getReportingUnitWinner(reportingUnit)


   return Functions.encodeString(JSON.stringify(raceWinner))
   ',
           'Karen Stepanyan',
           'This Function returns the winner of the US election for a given date. It uses the AP (Associated Press) API to get the results. The date is the only required parameter. API key is the only required secret.'
       ),
       (
           3,
           'Aggregate the ERC20 balance of an address across multiple chains',
           '2023-03-30',
           '// https://github.com/polar0/cross-chain-ERC20-balance-verification/blob/main/implementation/verify-balances.js

   // The address to check the balances of
   const userAddress = args[0]
   // The chains to check, formatted as:
   // name:tokenAddress,name:tokenAddress...
   const tokens = args[1].split(",").map((tokenAddress) => {
     const [chain, address] = tokenAddress.split(":")
     return { chain, address }
   })

   // Verify if there is indeed a secret (RPC URL) for each chain
   tokens.forEach((token) => {
     if (!secrets[token.chain]) {
       throw new Error(`No secret found for chain ${token.chain}`)
     }
   })

   // Prepare requests for each chain
   const requests = tokens.map((token, index) => {
     return Functions.makeHttpRequest({
       url: secrets[token.chain],
       method: "POST",
       data: {
         id: index,
         jsonrpc: "2.0",
         method: "eth_call",
         params: [
           {
             to: token.address,
             // The signature of ''balanceOf(address)'' + the user address without the 0x prefix
             data: "0x70a08231000000000000000000000000" + userAddress.slice(2),
           },
           "latest",
         ],
       },
     })
   })

   // Wait for all requests to finish
   const responses = await Promise.all(requests)

   // Parse responses
   const balances = responses.map((response) => {
     // Convert the result to a number
     return parseInt(response.data.result, 16) ?? 0
   })

   // Sum all balances
   const totalBalance = balances.reduce((a, b) => a + b, 0)

   // Return the total balance of the user
   return Functions.encodeUint256(totalBalance)
   ',
           'polarzero',
           'Find the balance of a user for a specific ERC20 token across the specified chains, and return the total balance. This balance, for example, could be used immediately in the callback function to approve or deny the user access to specific functions in the contract.'
       ),
       (
           4,
           'Find the Best DEX Trade Value for a Given Asset Pair',
           '2023-03-30',
           '// Decimals can be passed from the token contract decimals() function
   const srcToken = args[0] // Token source (selling)
   const srcDecimals = args[1]
   const destAsset = args[2] //Token destination (buying)
   const destDecimals = args[3]
   const amount = args[4] // Amount of source token to trade

   // Pull from the Paraswap DEX Aggregator router
   const paraswapRequest = await Functions.makeHttpRequest({
     url: `https://apiv5.paraswap.io/prices?srcToken=${srcToken}&srcDecimals=${srcDecimals}&destToken=${destAsset}&destDecimals=${destDecimals}&amount=${amount}&network=1`,
   })

   if (!paraswapRequest.error) {
     console.log("Optimal trade route found!")
     console.log(
       `Swap found to exchange ${
         10 ** -paraswapRequest.data.priceRoute.srcDecimals * parseInt(paraswapRequest.data.priceRoute.srcAmount)
       } of ${paraswapRequest.data.priceRoute.srcToken} into ${
         10 ** -paraswapRequest.data.priceRoute.destDecimals * parseInt(paraswapRequest.data.priceRoute.destAmount)
       } of ${paraswapRequest.data.priceRoute.destToken}`
     )
     //Sample Output: "Swap found to exchange 1 of 0x514910771af9ca656af840dff83e8264ecf986ca into 6.732330036871376 of 0x6b175474e89094c44da98b954eedeac495271d0f"
     console.log(`${paraswapRequest.data.priceRoute.bestRoute.length} best route(s) found:`)
     //If direct swap is found with one pool return that pool address
     if (paraswapRequest.data.priceRoute.bestRoute[0].percent == 100) {
       console.log(
         `One direct route found through ${paraswapRequest.data.priceRoute.bestRoute[0].swaps[0].swapExchanges[0].exchange}`
       )
       //Sample Output: One direct route found through UniswapV2
       console.log(paraswapRequest.data.priceRoute.bestRoute[0].swaps[0].swapExchanges[0].data)
       /*
       Sample Output:
       {
         router: ''0xF9234CB08edb93c0d4a4d4c70cC3FfD070e78e07'',
         path: [
           ''0x514910771af9ca656af840dff83e8264ecf986ca'',
           ''0x6b175474e89094c44da98b954eedeac495271d0f''
         ],
         factory: ''0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f'',
         initCode: ''0x96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f'',
         feeFactor: 10000,
         pools: [
           {
             address: ''0x6D4fd456eDecA58Cf53A8b586cd50754547DBDB2'',
             fee: 30,
             direction: true
           }
         ],
         gasUSD: ''2.735657''
       }
       */
     }
   } else {
     console.log("Paraswap Request error")
     console.log({ ...paraswapRequest })
   }
   return Functions.encodeUint256(parseInt(paraswapRequest.data.priceRoute.destAmount))
   ',
           'Max Melcher',
           'This example shows how to return the best DEX trade value for a give asset pair using Paraswap DEX Aggregator'
       ),
       (
           5,
           'Fetch result of soccer match from Sportsdata.io',
           '2023-03-30',
           '// Chainlink function to get the winner of soccer match. Possible return values are abbreviations of team names or ''Draw''

   const date = args[0] // Match date. basic date format YYYY-MM-DD. for example 2023-01-28
   let teams = args[1] //  competing teams in following format TEAM1/TEAM2. for example AST/LEI

   if (!secrets.soccerApiKey) {
     throw Error("Sportsdata.io API KEY is required")
   }

   const config = {
     url: `https://api.sportsdata.io/v3/soccer/scores/json/GamesByDate/${date}?key=${secrets.soccerApiKey}`
   }

   const response = await Functions.makeHttpRequest(config)

   const allMatches = response.data;

   const match = allMatches.find(match => {
     const playingTeams = `${match.AwayTeamKey}/${match.HomeTeamKey}`.toUpperCase()
     const playingTeamsReversed = `${match.HomeTeamKey}/${match.AwayTeamKey}`.toUpperCase()
     if (teams.toUpperCase() === playingTeams || teams.toUpperCase() === playingTeamsReversed) {
       return true
     }
   })

   if (!match) {
     throw new Error(''Match not found for given arguments'')
   }

   if (match.Winner === ''Scrambled'') {
     throw new Error(''Data is scrambled, use production API Key'')
   }

   let result;

   if (match.Winner === ''AwayTeam'') {
     result = match.AwayTeamKey
   } else if (match.Winner === ''HomeTeam'') {
     result = match.HomeTeamKey
   } else if (match.Winner === ''Draw'') {
     result = ''Draw''
   }

   if (!result) {
     throw new Error(''Could not get the winner team.'')
   }

   return Functions.encodeString(result)
   ',
           'Karen Stepanyan',
           'The function fetches the result of soccer match. Required arguments are match date and abbreviations of team names'
       ),
       (
           6,
           'Prompt AI for a response',
           '2023-03-30',
           'const prompt = args[0]

   if (
       !secrets.openaiKey
   ) {
       throw Error(
           "Need to set OPENAI_KEY environment variable"
       )
   }

   // example request:
   // curl https://api.openai.com/v1/completions -H "Content-Type: application/json" -H "Authorization: Bearer YOUR_API_KEY" -d ''{"model": "text-davinci-003", "prompt": "Say this is a test", "temperature": 0, "max_tokens": 7}

   // example response:
   // {"id":"cmpl-6jFdLbY08kJobPRfCZL4SVzQ6eidJ","object":"text_completion","created":1676242875,"model":"text-davinci-003","choices":[{"text":"\n\nThis is indeed a test","index":0,"logprobs":null,"finish_reason":"length"}],"usage":{"prompt_tokens":5,"completion_tokens":7,"total_tokens":12}}
   const openAIRequest = Functions.makeHttpRequest({
       url: "https://api.openai.com/v1/completions",
       method: "POST",
       headers: {
           ''Authorization'': `Bearer ${secrets.openaiKey}`
       },
       data: { "model": "text-davinci-003", "prompt": prompt, "temperature": 0, "max_tokens": 7 }
   })

   const [openAiResponse] = await Promise.all([
       openAIRequest
   ])
   console.log("raw response", openAiResponse)

   const result = openAiResponse.data.choices[0].text
   return Functions.encodeString(result)
   ',
           'Patrick Collins',
           'Ask OpenAI (or any AI model you want to interact with) for information on-chain.'
       ),
       (
           7,
           'Read cross-chain information',
           '2023-03-30',
           '// This example shows how to make a decentralized price feed using multiple APIs

   // Arguments can be provided when a request is initated on-chain and used in the request source code as shown below
   const contractAddress = args[0]
   const encodedAbiFunctionCall = args[1]

   if (
       !secrets.polygonKey
   ) {
       throw Error(
           "Need to set POLYGON_RPC_URL environment variable"
       )
   }

   // curl --data ''{"method":"eth_call","params":[{"to":"0x794a61358D6845594F94dc1DB02A252b5b4814aD","data":"0x35ea6a750000000000000000000000007ceb23fd6bc0add59e62ac25578270cff1b9f619"},"latest"],"id":1,"jsonrpc":"2.0"}'' -H "Content-Type: application/json" -X POST $POLYGON_RPC_URL
   // example response:
   // {"jsonrpc":"2.0","id":1,"result":"0x000000000000000000000003e80000069140000039cb03e805122904203a1f400000000000000000000000000000000000000000033e9fbcc201bc653e561a5300000000000000000000000000000000000000000002542e73dd9e8a5aecdb2a0000000000000000000000000000000000000000034895c6e6312a938da89522000000000000000000000000000000000000000000123f39e6ba5158357302ea0000000000000000000000000000000000000000004a723dc6b40b8a9a0000000000000000000000000000000000000000000000000000000000000063e965ca0000000000000000000000000000000000000000000000000000000000000004000000000000000000000000e50fa9b3c56ffb159cb0fca61f5c9d750e8128c8000000000000000000000000d8ad37849950903571df17049516a5cd4cbe55f60000000000000000000000000c84331e39d6658cd6e6b9ba04736cc4c473435100000000000000000000000003733f4e008d36f2e37f0080ff1c8df756622e6f00000000000000000000000000000000000000000000000001e758ee6c676a3f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

   // To make an HTTP request, use the Functions.makeHttpRequest function
   // Functions.makeHttpRequest function parameters:
   // - url
   // - method (optional, defaults to ''GET'')
   // - headers: headers supplied as an object (optional)
   // - params: URL query parameters supplied as an object (optional)
   // - data: request body supplied as an object (optional)
   // - timeout: maximum request duration in ms (optional, defaults to 10000ms)
   // - responseType: expected response type (optional, defaults to ''json'')

   // Ideally, you''d use multiple RPC URLs so we don''t have to trust just one
   const polygonReadRequest = Functions.makeHttpRequest({
       url: secrets.polygonKey,
       method: "POST",
       data: {
           "jsonrpc": "2.0",
           "method": "eth_call",
           "params": [
               { "to": contractAddress, data: encodedAbiFunctionCall },
               "latest"
           ],
           "id": 1
       }
   })

   // First, execute all the API requests are executed concurrently, then wait for the responses
   const [polygonResponse] = await Promise.all([
       polygonReadRequest
   ])

   console.log("raw response", polygonResponse)

   // take the "0x" off the front of the hex string
   const result = polygonResponse.data.result.slice(2)

   // loop through result and convert each 64 characters to a number
   const startingIndex = 64 * 2
   const supplyApy = "0x" + result.slice(startingIndex, startingIndex + 64)

   // convert the hex supplyApy to a number
   const supplyApyNumber = parseInt(supplyApy, 16)
   // This number is returned as a RAY, so we''d divide by 1e27, or 1e25 to get a percentage
   console.log("WETH Supply APY on AaveV3 in Polygon: ", (supplyApyNumber / 1e25), "%")

   // The source code MUST return a Buffer or the request will return an error message
   // Use one of the following functions to convert to a Buffer representing the response bytes that are returned to the client smart contract:
   // - Functions.encodeUint256
   // - Functions.encodeInt256
   // - Functions.encodeString
   // Or return a custom Buffer for a custom byte encoding
   // return Functions.encodeUint256(Math.round(medianPrice * 100))
   return Functions.encodeUint256(supplyApyNumber)
   ',
           'Patrick Collins',
           'The function reads the supply APY rate of depositing WETH into AaveV3 on Polygon'
       ),
       (
           8,
           'Fetch outcome of off-chain Snapshot.org vote',
           '2023-03-30',
           'const proposalID = args[0]

   if (!proposalID) {
     throw Error("Proposal ID is required")
   }

   const config = {
     url: "https://hub.snapshot.org/graphql?",
     method: "POST",
     headers: {
       ''content-type'': ''application/json''
     },
     params: {
       operationName: "Proposal",
       query: `query Proposal {\n  proposal(id:"${proposalID}") {\n    id\n    votes\n   scores\n  choices\n  state\n  scores_total\n quorum\n}\n}`,
       variables: null,
     },
   }

   const response = await Functions.makeHttpRequest(config)

   const state = response.data.data.proposal.state
   const totalScore = response.data.data.proposal.scores_total
   const quorum = response.data.data.proposal.quorum

   if (state !== ''closed'') {
     return Functions.encodeString(''Vote not ended'')
   }

   if (totalScore < quorum) {
     return Functions.encodeString(''Quorum not met'')
   }

   const scores = response.data.data.proposal.scores
   const choices = response.data.data.proposal.choices
   const highestIndex = scores.indexOf(Math.max(...scores));

   return Functions.encodeString(choices[highestIndex])
   ',
           'ChainLinkGod',
           'The function fetches the outcome of an off-chain Snapshot.org vote proposal using the GraphQL API. Takes into account if the vote has closed and has met quorum. Gas efficient solution for DAOs.'
       ),
       (
           9,
           'Financial metric data for dApps and blockchains sourced from Token Terminal',
           '2023-03-30',
           '
   const metric = args[0] // valid metric id that can be found on https://api.tokenterminal.com/v2/metrics
   const project = args[1] // project id
   const date = args[2] // optional date. format YYYY-MM-DD. For example 2023-02-10
   const apiKey = secrets.API_KEY;


   if (!apiKey) {
     throw Error("Tokenterminal API Key is required")
   }

   const config = {
     url: `https://api.tokenterminal.com/v2/metrics/${metric}?project_ids=${project}`,
     headers: {
       ''Authorization'': `Bearer ${apiKey}`
     }
   }

   const response = await Functions.makeHttpRequest(config)
   if (response.error) {
     throw new Error(response.response.data.message)
   }

   let data;
   if (date) {
     data = response.data.data.find(d => d.timestamp.includes(date))
   }else {
     data = response.data.data[0]
   }
   const result = Math.round(data.value * 100)

   return Functions.encodeUint256(result)
   ',
           'ChainLinkGod',
           'This Function fetches metric data from the Token Terminal API for a specific project. Supported metrics include revenue, fees, earnings, active users, TVL, volume, supply, and more. Projects includes both dApps and blockchains. Optional parameter for specific date. Requires Token Terminal Pro subscription to obtain API key.'
       ),
       (
           10,
           'Obtain outcome of off-chain vote',
           '2023-03-30',
           'const proposalId = args[0]

   // Use snapshot''s graphql API to get the final vote outcome
   const snapshotRequest = () => Functions.makeHttpRequest({
     url: `https://hub.snapshot.org/graphql`,
     method: "POST",
     data: {
       query: `{
         proposal(id: "${proposalId}") {
           choices
           scores
           scores_state
         }
       }`,
     },
   })

   const { data, error } = await snapshotRequest()

   if (error) {
     throw Error("Snapshot request failed")
   }

   const { proposal } = data.data
   const { choices, scores, scores_state } = proposal

   if (scores_state !== "final") {
     throw Error("Snapshot vote is not final")
   }

   const winningChoice = choices[scores.indexOf(Math.max(...scores))]
   return Functions.encodeString(winningChoice)
   ',
           'mykcryptodev',
           'This function fetches the final outcome of an off-chain vote on the Snapshot.org platform'
       ),
       (
           11,
           'Fetch and return available balance of Stripe account',
           '2023-03-30',
           'const apiKey = secrets.API_KEY
   const balanceCurrency = args[0] || ''usd''

   if (!apiKey) {
     throw Error("Stripe API Key is required")
   }


   const config = {
     url: `https://${apiKey}@api.stripe.com/v1/balance`,
   }

   const response = await Functions.makeHttpRequest(config)

   const balance = response.data.available.find(c => c.currency.toLowerCase() === balanceCurrency.toLowerCase())

   const balanceInCents = Math.round(balance.amount * 100)

   return Functions.encodeUint256(balanceInCents)
   ',
           'Karen Stepanyan',
           'This function will fetch Stripe account available balance of particular currency.'
       ),
       (
           12,
           'Calculate the median price of a token on Uniswap V2',
           '2023-03-30',
           '// Max sample size is 4 due to 5 http request limit
   const SAMPLE_SIZE = 4
   // The number of decimals the price in USD is formatted to
   const DECIMALS = 18
   // A block buffer to take into consideration the synchronization of the subgraph
   const GRAPH_BLOCK_BUFFER = 50
   const AVG_SECONDS_PER_BLOCK = 12

   // Token address
   const token = args[0].toLowerCase();
   // Pair address
   const pair = args[1]
   // Period in seconds
   const period = args[2]

   const blockRange = period / AVG_SECONDS_PER_BLOCK

   if (!secrets.rpc) {
     throw Error("\"rpc\" environment variable not set")
   }

   const blockNumberResponse = await Functions.makeHttpRequest({
     url: secrets.rpc,
     method: "POST",
     headers: {
       "Accept": "application/json",
       "Content-Type": "application/json",
     },
     data: JSON.stringify({
       jsonrpc: "2.0",
       method: "eth_blockNumber",
       params: [],
       id: "1",
     }),
   })

   if (blockNumberResponse.error) {
     throw Error("Unable to fetch current block number")
   }

   const blockNumber = parseInt(blockNumberResponse.data.result, 16) - GRAPH_BLOCK_BUFFER

   const fetchPrice = (blockNumber) => Functions.makeHttpRequest({
     url: "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2",
     method: "POST",
     data: {
       query: `{
         pair(id: "${pair}", block: {number: ${blockNumber}}) {
           token0 {
             id
           }
           token1 {
             id
           }
           reserve0
           reserve1
           reserveUSD
         }
       }`,
     },
   })

   const stringToBigInt = (str) => {
     const splitStr = str.split(".")
     const decimals = splitStr[1].slice(0, DECIMALS).padEnd(DECIMALS, "0")
     return BigInt(`${splitStr[0]}${decimals}`)
   }

   const getPrice = async (blockNumber) => {
     const {
       error,
         data: {
           errors,
           data,
         },
     } = await fetchPrice(blockNumber)
     if (error.error || errors) {
       throw Error("Unable to fetch price from subgraph")
     }
     const { pair: { token0: { id: token0 }, token1: { id: token1 }, reserve0, reserve1, reserveUSD } } = data
     const token0LC = token0.toLowerCase()
     const token1LC = token1.toLowerCase()
     if (token0LC !== token && token1LC !== token) {
       throw Error("Token not found as part of the pair")
     }
     const tokenReserveInUSD = stringToBigInt(reserveUSD) / 2n
     const tokenReserve = stringToBigInt(token0LC === token ? reserve0 : reserve1)
     return BigInt(10 ** DECIMALS) * tokenReserveInUSD / tokenReserve
   }

   const pickRandomBlock = () => {
     return blockNumber - Math.round(Math.random() * blockRange)
   }

   let prices = []
   for (let i = 0; i < SAMPLE_SIZE; i++) {
     const price = await getPrice(pickRandomBlock())
     prices.push(price)
   }

   const midpoint = SAMPLE_SIZE % 2 === 0 ? SAMPLE_SIZE / 2 : (SAMPLE_SIZE + 1) / 2
   const median = prices[midpoint]

   return Functions.encodeUint256(median)
   ',
           'moonthoon',
           'This function calculates the median price of a token that is on Uniswap V2. It works by sampling up to 4 prices over a given time period then chooses the median value'
       ),
       (
           13,
           'Twitter account verification with an Ethereum address',
           '2023-03-30',
           '// https://github.com/polar0/twitter-verifier-chainlink-functions/blob/main/implementation/twitter-verification/functions/Functions-request-source.js

   // Get the arguments from the request config
   const twitterUsername = args[0]; // e.g. ''TwitterDev''
   const ethereumAddress = args[1]; // e.g. ''0x1234567890123456789012345678901234567890''
   // The string that must be included in the latest tweets of the user for the verification to pass
   const requiredStringIncluded = `Verifying my Twitter account for ${ethereumAddress}`;
   // How many tweets to check (min 5, max 100)
   const MAX_RESULTS = 10;

   // Initialize the result to -1 (error)
   let result = -1;

   // Get the bearer token from the environment variables
   if (!secrets.apiKey) {
     throw Error(
       ''TWITTER_BEARER_TOKEN environment variable not set for Twitter API. Get a free one: https://developer.twitter.com/en/docs/authentication/oauth-2-0/bearer-tokens'',
     );
   }

   // Don''t even try if the username or address is empty
   if (!twitterUsername || !ethereumAddress) {
     throw Error(''Twitter username or Ethereum address is empty'');
   }

   // Prepare the API requests
   const twitterRequest = {
     // Get the user id from the provided username
     userIdByUsername: () =>
       Functions.makeHttpRequest({
         url: `https://api.twitter.com/2/users/by/username/${twitterUsername}`,
         headers: { Authorization: `Bearer ${secrets.apiKey}` },
       }),
     // Get the latest n tweets from the user (n = MAX_RESULTS)
     lastTweetsByUserId: (userId) =>
       Functions.makeHttpRequest({
         url: `https://api.twitter.com/2/users/${userId}/tweets?max_results=${MAX_RESULTS}`,
         headers: { Authorization: `Bearer ${secrets.apiKey}` },
       }),
   };

   // First, request the user id from their username
   const idRes = await new Promise((resolve, reject) => {
     twitterRequest.userIdByUsername().then((res) => {
       if (!res.error) {
         resolve(res);
       } else {
         reject(res);
       }
     });
   });

   if (idRes.error) {
     throw Error(''Twitter API request failed - coult not get user id'');
   }

   // Grab the user id
   const userId = idRes.data.data.id || null;

   // Let''s be extra careful and make sure the user id is not null
   if (!userId) {
     throw Error(''Twitter API request failed - user id is null'');
   }

   // Then, request the latest tweets
   const tweetsRes = await new Promise((resolve, reject) => {
     twitterRequest.lastTweetsByUserId(userId).then((res) => {
       if (!res.error) {
         resolve(res);
       } else {
         reject(res);
       }
     });
   });

   if (tweetsRes.error) {
     throw Error(''Twitter API request failed - coult not get tweets'');
   }

   // It''ll only get here if the request was successful
   const tweets = tweetsRes.data.data;
   const tweetTexts = tweets.map((tweet) => tweet.text);
   // Check if any of these tweets include the required string
   const res = tweetTexts.some((text) =>
     text.toLowerCase().includes(requiredStringIncluded.toLowerCase()),
   );
   // If it found the string, return 1, otherwise 0
   result = res ? 1 : 0;

   // `result` can either be:
   // - 1 (verified)
   // - 0 (not verified)
   // - -1 (if by any chance no error was thrown, yet it could not verify)

   // Return the result along with the username and address, which can be parsed and split
   return Functions.encodeString(
     `${result},${twitterUsername},${ethereumAddress}`,
   );

   ',
           'polarzero',
           'Check if a Twitter account belongs to a specific Ethereum address. This example uses the Twitter API to retrieve a user''s recent tweets, and checks if they tweeted a specific message containing their address. It provides the arguments and returns the result via Chainlink Functions, which allows for prior validation of the user''s ownership of the address via a signature or other method, thus performing a secure and non-intrusive verification.'
       ),
       (
           14,
           'Price data from multiple sources',
           '2023-03-30',
           'const coinMarketCapCoinId = args[0];
   const coinGeckoCoinId = args[1];
   const coinPaprikaCoinId = args[2];
   const badApiCoinId = args[3];

   const scalingFactor = parseInt(args[4]);

   if (!secrets.apiKey) {
     throw Error(''API_KEY environment variable not set for CoinMarketCap API.  Get a free key from https://coinmarketcap.com/api/'')
   }

   // OCR2DR.makeHttpRequest function parameters:
   // - url
   // - method (optional, defaults to ''GET'')
   // - headers: headers supplied as an object (optional)
   // - params: URL query parameters supplied as an object (optional)
   // - data: request body supplied as an object (optional)
   // - timeout: maximum request duration in ms (optional, defaults to 10000ms)
   // - responseType: expected response type (optional, defaults to ''json'')

   // Use multiple APIs & aggregate the results to enhance decentralization
   const coinMarketCapResponse = await OCR2DR.makeHttpRequest({
     url: `https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?convert=USD&id=${coinMarketCapCoinId}`,
     // Get a free API key from https://coinmarketcap.com/api/
     headers: { ''X-CMC_PRO_API_KEY'': secrets.apiKey }
   });
   const coinGeckoResponse = await OCR2DR.makeHttpRequest({
     url: `https://api.coingecko.com/api/v3/simple/price?ids=${coinGeckoCoinId}&vs_currencies=usd`,
   });
   const coinPaprikaResponse = await OCR2DR.makeHttpRequest({
     url: `https://api.coinpaprika.com/v1/tickers/${coinPaprikaCoinId}`
   });
   const badApiResponse = await OCR2DR.makeHttpRequest({
     url: `https://badapi.com/price/symbol/${badApiCoinId}`
   });

   const prices = [];

   if (!coinMarketCapResponse.error) {
     prices.push(coinMarketCapResponse.data.data[coinMarketCapCoinId].quote.USD.price);
   }
   else {
     console.log(''CoinMarketCap Error'');
     console.log({ ...coinMarketCapResponse });
   }
   if (!coinGeckoResponse.error) {
     prices.push(coinGeckoResponse.data[coinGeckoCoinId].usd);
   } else {
     console.log(''CoinGecko Error'');
     console.log({ ...coinGeckoResponse });
   }
   if (!coinPaprikaResponse.error) {
     prices.push(coinPaprikaResponse.data.quotes.USD.price);
   } else {
     console.log(''CoinPaprika Error'');
     console.log({ ...coinPaprikaResponse });
   }

   // A single failed API request does not cause the whole request to fail
   if (!badApiResponse.error) {
     prices.push(httpResponses[3].data.price.usd);
   } else {
     console.log(''Bad API request failed. (This message is expected and just for demonstration purposes.)'')
   }

   // At least 3 prices are needed to aggregate the median price
   if (prices.length < 3) {
     // If an error is thrown, it will be returned back to the smart contract
     throw Error(''More than 1 API failed'');
   }

   const medianPrice = prices.sort((a, b) => a - b)[Math.round(prices.length / 2)];
   console.log(`Median Bitcoin price: $${medianPrice.toFixed(2)}`);

   // Use the following functions to encode a single value:
   // - OCR2DR.encodeUint256
   // - OCR2DR.encodeInt256
   // - OCR2DR.encodeString
   // Or return a Buffer for a custom byte encoding
   return OCR2DR.encodeUint256(Math.round(medianPrice * 100));
   ',
           'Morgan Kuphal',
           'Retrieve the price of an asset from multiple API sources. Assets could be practially anything, incuding equities, crypto, or commodities. This example pulles from multiple different data providers (APIs) and derrives the median to return on chain via Chainlink Functions.'
       );