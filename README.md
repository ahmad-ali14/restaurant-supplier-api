# restaurant-supplier-api

### This is my first attempt to work with Go, this project is done quickly and not optimized 

## Live

the api is live on: https://restaurant-supplier-api--ahmadali5.repl.co/

live with the front end: https://restaurant-supplier.netlify.app/

## Test it 

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
            "Key": "password",
            "Value": "$2a$14$YqfqDPQ3r59RRtzjG4iwReizQlDDhZkx2oIiTJRXcngvx6fFvpXCC"
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



## Env Files

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
