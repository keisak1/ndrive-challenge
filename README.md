# Fullstack NDrive Challenge

## Installation
 - Create a new file on the api/config folder with the name **app.env** (Change any field as you wish)
   ```env
   PORT=
   MONGO_ROOT_USERNAME=
   MONGODB_URI=
   ACCESS_TOKEN_PRIVATE_KEY=
   ACCESS_TOKEN_PUBLIC_KEY=
   ACCESS_TOKEN_EXPIRED_IN=15m 
   ACCESS_TOKEN_MAXAGE=15
   REFRESH_TOKEN_PRIVATE_KEY=
   REFRESH_TOKEN_PUBLIC_KEY=
   REFRESH_TOKEN_EXPIRED_IN=60
   REFRESH_TOKEN_MAXAGE=60```
   
- Install the Golang viper package
  ```shell
  go get github.com/spf13/viper
  ```

- Run the command below to install the MongoDB Go driver and Gin Gonic packages.
  ```shell
  go get go.mongodb.org/mongo-driver/mongo github.com/gin-gonic/gin
  ```

