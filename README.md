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


## API Endpoints

 ### Auth GET Requests
  **Refresh token**

|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `none` |             |         | Refresh your access token.                                                                    |

 **Response**

 ```json
{
  "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk3MDEwNTAsImlhdCI6MTY4OTcwMDE1MCwibmJmIjoxNjg5NzAwMTUwLCJzdWIiOiI2NGI1NTgwMWU0MGUyMTJiNzI4ODY3ZTcifQ.0dGiTh16sPWgmcC50NbQPzemP9V2xEXppH65eVh0yFeZbL9ry3rxNMplYgfjv5d5J2W5uQxHPiY7GBWc6YIyog",
  "status": "success"
}
```
**Login**
|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `email` |    required      |    string     | The email address linked to a registered account.                                                                    |
|     `password` |    required      |    string     | A valid password to login                                                                    |

 **Response**
```json
{
  "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODk3MDEwMjYsImlhdCI6MTY4OTcwMDEyNiwibmJmIjoxNjg5NzAwMTI2LCJzdWIiOiI2NGI1NTgwMWU0MGUyMTJiNzI4ODY3ZTcifQ.eV8ngCjVtF1jcyPMleAILmOsithnsBwrJtVdMTx873z-Lir098pg0PzeBgEWod7N74hbvZWcmcOzsiYaSqVghQ",
  "status": "success"
}
```

