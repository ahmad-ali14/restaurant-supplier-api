# restaurant-supplier-api

## Env Files

- make a `./config/config.go` file
- put this code in that file:

  ```go
        package config

        func GetMongoUrl() string {
          return "<your mongodb connection string>"
        }

  ```
