# E-Wallet Proof of Concept

## Table of Contents

- [About](#about)
- [System Design](#system_design)
- [Scratch Schema](#scratch_schema)
- [Usage](#usage)

## About <a name = "about"></a>

This project is only for PoC purpose. All data will be stored locally without any database connection. All design and schema will explain below.

## System Design <a name = "system_design"></a>

![ScratchDesign](/scratch_design.png)

This is the scratch system design of E-Wallet service.

The technologies stack:

- ewallet-service: the core service that will handling all transaction of ewallet
- redis: to store latest wallet data of organization. so if service need to update the data, it can get the data from redis first (caching) then update the latest data to database. this method uses to reduce database read transaction
- SQL: this relational database will be used to store organization information detail, wallet information detail, and relation between organization and wallet. since one organization can have many wallet relation
- No SQL: this no-sql technology will be use to store transactional data records. since the transaction rate is high, I believe it is better to use no-sql to records the transactional. About data consistency, user is rare to see the web, so no-sql will not have problem about it.
- 3rd Party Money Transaction: this is client to send money from organization to recipients. Ewallet service need to provides callback endpoint that accessible from 3rd Party Payment. So, it will get latest data and align the data between local system and 3rd party system
- Scheduler service: incase 3rd party will not inform expired transaction, it is necessary to add scheduler service with expired date as publish date. so wallet service can check whether the transaction is still ACTIVE or not
- 3rd Party Email Notification: since the user is rare to open web page, it is better to integrate email notification if there is an error with transaction (insufficient balance, expired/cancel transaction, etc)

## Scratch Schema <a name = "scratch_schema"></a>

![ScratchSchema](/scratch_schema.png)
This is the scratch schema for ewallet service with SQL schema (blue) and NoSQL schema (green).

Terminologies:
in transaction history, there are status and type that will be a CONSTANT

type:

- ADDING to store balance
- SENDING to send transaction

status:

- ACTIVE
- SUCCESS
- EXPIRED
- CANCEL

## Usage <a name = "usage"></a>

Before you run this code, you must run these commands below

```
go mod download && go mod tidy
```

To test this code, you need to have Go in your local and run command below

```
go run main.go
```

there will be 2 endpoints to test this code, example cURLs are

**POST transaction**

```
curl -X POST \
  'http://localhost:8080/ewallet/v1/payment/organization0' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "wallet_id": 1,
  "transaction_type": "SENDING",
  "total_amount": 1000,
  "data": [
    {
      "wallet_id": 1,
      "recipient":"c1",
      "type": "SENDING",
      "amount": 500,
      "status": "SUCCESS"
    },
    {
      "wallet_id": 1,
      "recipient":"c2",
      "type": "SENDING",
      "amount": 500,
      "status": "SUCCESS"
    }
  ]
}'
```

**GET organization detail**

```
curl -X GET \
  'http://localhost:8080/ewallet/v1/payment/organization0?wallet_id=all' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'
```
