# restaurant-supplier-api

### This is my first attempt to work with Go, this project is done quickly and not optimized 

## Live

the api is live on: https://restaurant-supplier-api--ahmadali5.repl.co/

live with the front end: https://restaurant-supplier.netlify.app/

### Info
- 4 modules for data `users, restaurants, suppliers, orders`
- users : stores authentication data from both restaurants and suppliers.
- orders: stores orders data with their status, messages.
- Not much actions done with orders yet.


## Test it

- the main route is not guarded and ots response is public:

```json
  {
  "_welcome ": {
  "info": "resturant-supplier-api endpoints by Ahmad Ali"
  },
  "route / ": {
  "Available Methods": "GET only",
  "more info": "no authentication yet",
  "this endpoint": "will give you more info about the Available routes by this api"
  },
  "route /order ": {
  "Available Methods": "CRUD",
  "this endpoint": "CRUD on orders"
  },
  "route /restaurant ": {
  "Available Methods": "CRUD",
  "this endpoint": "CRUD on restaurnts"
  },
  "route /supplier ": {
  "Available Methods": "CRUD",
  "this endpoint": "CRUD on suppliers"
  }
  }
```

## Test it As Restaurant

1. make post request to: https://restaurant-supplier-api--ahmadali5.repl.co/login with body:
  ```json
  {
    "email":"mc@london.uk",
    "password":"london"
  }
  ```
  
  and you will get this result:
  
  ```json
    {
    "role": "restaurant",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjbGllbnQiOiJFbGxpb3QgRm9yYmVzIiwiZXhwIjoxNTk3MDE4NzUwfQ.0PPK73nbjSDNwRKc3OJBo_PCPo61ZGgV8CEUpXSEZC0",
    "userAccount": [
        {
            "Key": "_id",
            "Value": "5f1ed376741bfdfe5ba916ad"
        },
        {
            "Key": "name",
            "Value": "MC Donalds London 1"
        },
        {
            "Key": "address",
            "Value": "nw"
        },
        {
            "Key": "email",
            "Value": "mc@london.uk"
        },
        {
            "Key": "role",
            "Value": "restaurant"
        }
    ],
    "userId": "5f1ed376741bfdfe5ba916ad"
}
```

2. add a new header called `Token` to your request so you will be authorized.


## Test it As Supplier

1. make post request to: https://restaurant-supplier-api--ahmadali5.repl.co/login with body:
  ```json
  {
    "email":"mc3@london.uk",
    "password":"lodon"
  }
  ```
  
  and you will get this result:
  
  ```json
   {
    "role": "supplier",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjbGllbnQiOiJFbGxpb3QgRm9yYmVzIiwiZXhwIjoxNTk3MDE4OTM0fQ.Up8sKj6RF5OSjam1KCFWTAiSmp_a0BKKujeiyGeHFCs",
    "userAccount": [
        {
            "Key": "_id",
            "Value": "5f1f3ee4cfedf9af9b6bdd08"
        },
        {
            "Key": "email",
            "Value": "mc3@london.uk"
        },
        {
            "Key": "address",
            "Value": "mc1"
        },
        {
            "Key": "phone",
            "Value": "+123435"
        },
        {
            "Key": "role",
            "Value": "supplier"
        },
        {
            "Key": "products",
            "Value": [
                [
                    {
                        "Key": "productName",
                        "Value": "pr 1"
                    },
                    {
                        "Key": "productPrice",
                        "Value": "1"
                    }
                ],
                [
                    {
                        "Key": "productName",
                        "Value": "pr 2"
                    },
                    {
                        "Key": "productPrice",
                        "Value": "2"
                    }
                ],
                [
                    {
                        "Key": "productName",
                        "Value": "pr 3"
                    },
                    {
                        "Key": "productPrice",
                        "Value": "3"
                    }
                ]
            ]
        }
    ],
    "userId": "5f1f3ee4cfedf9af9b6bdd08"
}
```

2. add a new header called `Token` to your request so you will be authorized.



## Test Locally

- make a `./config/config.go` file
- put this code in that file:

  ```go
        package config

        func GetMongoUrl() string {
          return "<your mongodb connection string>"
        }

  ```

  ## live
  [![Run on Repl.it](https://repl.it/badge/github/ahmad-ali14/restaurant-supplier-api)](https://repl.it/github/ahmad-ali14/restaurant-supplier-api)
