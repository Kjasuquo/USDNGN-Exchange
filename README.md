# cmd
This houses the router, server and main.go needed to start the app

# config 
This houses the configuration variables gotten from which ever environment the app is running or being deployed

# internal
This houses all the logic implementation for the app to work. It includes the controllers, the database, middleware, models, services, port, etc.
The models, controllers and even the database collections are divided into 2 just to keep it simple
- user
- transaction

# Getting Started
To get started, you can run the dockerized application using ``` make up```. It will automatically start mongodb and run the service on localhost ```http://localhost:8085```.

The routes are divided mainly into:
### 1. Open routes which has:
- /api/v1/ping - for testing connection
- /api/v1/signup - to create user [POST]
```bash
sample payload
{
  "full_name": "Joseph Asuquo",
  "phone_number": "08133477843",
  "email": "okoasuquo@yahoo.com",
  "password": "okoasuquo"
}
  
  sample response
{
    "data": null,
    "errors": "",
    "message": "user created",
    "status": "Created"
}
```
- /api/v1/login - to login a user [POST]
```bash
sample payload
{
  "email": "okoasuquo@yahoo.com",
  "password": "okoasuquo"
}
  
  sample response
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im9rb2FzdXF1b0B5YWhvby5jb20iLCJleHAiOjE2NzczMzIwMTN9.pHjCT3HFlGqPTbQKKtc5aH15LQ8GkCTNuNcdhdCTpds",
        "user": {
            "ID": "74f3588e298cf515dec298ba",
            "FullName": "Joseph Asuquo",
            "PhoneNumber": "08133477843",
            "Email": "okoasuquo@yahoo.com",
            "PasswordHash": "573a39574bd9c043994128aa",
            "Salt": "707850798748855153",
            "USDBalance": 100,
            "NGNBalance": 0,
            "CreatedAt": "2023-02-24T13:30:47.592Z"
        }
    },
    "errors": "",
    "message": "login successful",
    "status": "OK"
}
```
Note that the token is used as Bearer Token for authenticated routes!

### 2. Authenticated routes which has
A user must be logged in (i.e. provide a token inorder to access these routes)
Authenticated routes are further subdivided into 2.
1. user:
- /api/v1/user/profile - to view a user profile [GET]
```bash
  sample response
{
    "data": {
        "ID": "000000000000000000000000",
        "FullName": "Joseph Asuquo",
        "PhoneNumber": "08133477843",
        "Email": "okoasuquo@yahoo.com",
        "PasswordHash": "",
        "Salt": "",
        "USDBalance": 100,
        "NGNBalance": 0,
        "CreatedAt": "2023-02-24T12:54:59.825Z"
    },
    "errors": "",
    "message": "successful",
    "status": "OK"
}
```

- /api/v1/user/balances - to view a user balance in USD and NGN [GET].
Assuming a transaction had been done by the user, his/her balance are (depending on the transaction done).
```bash
  sample response
{
    "data": {
        "USD": 76.4375413086583,
        "NGN": 16825
    },
    "errors": "",
    "message": "successful",
    "status": "OK"
}
```

2. Transaction:
- /api/v1/transaction/sellUSD - to change USD to NGN [POST]
```bash
  sample payload
{
    "amount": 50
}
  
  
  sample response
{
    "data": null,
    "errors": "",
    "message": "transaction successful",
    "status": "OK"
}
```

- /api/v1/transaction/buyUSD - to change NGN to USD [POST]
```bash
  sample payload
{
    "amount": 20000
}
  
  
  sample response
{
    "data": null,
    "errors": "",
    "message": "transaction successful",
    "status": "OK"
}
```

- /api/v1/transaction/transactions - used to get the record of transactions made by the logged-in user [GET]
```bash
  sample response
{
    "data": [
        {
            "ID": "63f8c2a0cdbccbc48e91bcb3",
            "UserId": "96c9e6dfd9de40d8622a78af",
            "Type": "sell_usd",
            "Rate": 736.49,
            "RequestCurrency": "USD",
            "RequestAmount": 10,
            "ReceivedCurrency": "NGN",
            "ReceivedAmount": 7364.9,
            "CreatedAt": "2023-02-24T13:58:56.626Z"
        },
        {
            "ID": "63f8b4c7cdbccbc48e91bcb2",
            "UserId": "96c9e6dfd9de40d8622a78af",
            "Type": "buy_usd",
            "Rate": 756.5,
            "RequestCurrency": "NGN",
            "RequestAmount": 20000,
            "ReceivedCurrency": "USD",
            "ReceivedAmount": 26.437541308658293,
            "CreatedAt": "2023-02-24T12:59:51.482Z"
        },
        {
            "ID": "63f8b4adcdbccbc48e91bcb1",
            "UserId": "96c9e6dfd9de40d8622a78af",
            "Type": "sell_usd",
            "Rate": 736.5,
            "RequestCurrency": "USD",
            "RequestAmount": 50,
            "ReceivedCurrency": "NGN",
            "ReceivedAmount": 36825,
            "CreatedAt": "2023-02-24T12:59:25.555Z"
        }
    ],
    "errors": "",
    "message": "successfully found",
    "status": "OK"
}
```

# Shutting Down
To shut down the app, after terminating, simply run the command ```make down```

# MOCKING and TESTING
To mock the port/database for integration testing, simply use the command ```make mock```.
To run the test files, use the command ```make test```
