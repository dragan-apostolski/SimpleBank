### How to run this app

Build docker image:
`docker build -t account-service .`

Run server
`docker run -p 8080:8080 account-service`

### Interacting via the REST API

The database has 4 prepopulated accounts (IDs 1-4). 
We can check the balance of the accounts using the /balance endpoint:

`curl -X GET "http://localhost:8080/balance?account_id=1"`

`Account balance: 200`

We can open another savings account with id 5 with 500 euros balance

`curl -X POST -d "account_id=5&amount_to_deposit=500" http://localhost:8080/open-savings-account`

And check its balance:
`curl -X GET "http://localhost:8080/balance?account_id=5"`

`Account balance: 500`

We can then withdraw from this account 200 euros:
`curl -X POST -d "account_id=5&amount_to_withdraw=200" http://localhost:8080/withdraw`


Check balance again, should be 500-200=300
`curl -X GET "http://localhost:8080/balance?account_id=5"`

`Account balance: 300`

If we try to withdraw another 250, there should be an info message that we can only withdraw up to 200 euros:
`curl -X POST -d "account_id=5&amount_to_withdraw=250" http://localhost:8080/withdraw`

`You can only withdraw up to 200 euros`