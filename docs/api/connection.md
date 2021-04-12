

### /api/users/:username/connections/addtoken`POST`
To create a new connection.
- Request
    ```json
    {
    "name" :"spotify3",
    "token_type" : "spotify",
    "token" : "1afcd222222sdcfaxfdfa52aafkkkbvdj"


}
    ```
- Response
    ```json
    {
        "message": "token added"
    }
    ```
### /api/users/:username/connection/gettoken `POST`
get the connection token.
- Request
    ```json
    {
        "name": "git1",
        "token_type": "github"
    }
    ```
- Response
    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxlZ2F0by"
    }
    ```


### /api/user/connection/access/token/urls `POST`
get you access url for apps you shod send token_type to back .if token_type = "spotify" redirect to spotify access url . other app is git discord google
- Response
    ```json
    {
        "spotify_url": "https://discord.com/api/oauth2/authorize?access_type=online&client_id=830463353079988314&redirect_uri=http://localhost:8080/callback&response_type=code&scope=identify+email&state=h8EecvhXJqHsG5EQ3K0gei4EUrWpaFj_HqH3WNZdrzrX1BX1COQRsTUv3-yGi3WmHQbw0EHJ58Rx1UOkvwip-Q%3D%3D"
    }


### /api/users/:username/connection/gettokens `GET`
get the connection token.
- Header
    - `Authorization` = `access_token`
 - Response
     ```json
    [
        {
            "ID": 10,
            "CreatedAt": "2021-04-11T20:23:44.487065Z",
            "UpdatedAt": "2021-04-11T20:23:44.490836Z",
            "DeletedAt": null,
            "Token": "1afcd222222sdcfaxsdfsaz",
            "Token_type": "spotify",
            "UserID": 3,
            "Name": "spotify4"
        },
        {
            "ID": 11,
            "CreatedAt": "2021-04-11T20:25:22.047339Z",
            "UpdatedAt": "2021-04-11T20:25:22.056967Z",
            "DeletedAt": null,
            "Token": "1afcd222222sdcfaxsdfsaz",
            "Token_type": "spotify",
            "UserID": 3,
            "Name": ""
        },
        {
            "ID": 12,
            "CreatedAt": "2021-04-12T14:48:46.605771Z",
            "UpdatedAt": "2021-04-12T14:48:46.626229Z",
            "DeletedAt": null,
            "Token": "dpsfihio514356dfsc5sdf5c1252dcsx52",
            "Token_type": "git",
            "UserID": 3,
            "Name": "git1"
        }
    ]
    ```
### /api/users/connection/update/token/name `PUT`
update name of a token with id
Request
    ```json
    {
        "name": "git1",
        "id": 1
    }
    ```
- Response
    ```json
    {
         "message": "update token successfully"
    }
    ```
### /api/users/:username/connection/check/token `POST`
get a id and check if there is a token with this id for a user return correct token
- Request
    ```json
    {
        "id": 1,
        
    }
    ```
- Response
    ```json
    {
        "message": "correct token"
    }
    ```

### /api/users/:username/connection/update/token/token `put`
get the connection token.
- Request
    ```json
    {
        "id" : 2,
        "token" : "asdpjbfipdsafhbcop;bfdhsbocp;x"
    ```
- Response
    ```json
    {
        "message":  "updated token successfully"
    }
    ```


### /api/users/:username/connection/update/token/token `put`
get the connection token.
- Request
    ```json
    {
        "id": 1,
        
    }
    ```
- Response
    ```json
    {
        "message":  "deleted token successfully"
    }
    ```