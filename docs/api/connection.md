### /api/users/:username/add/connection `POST`
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
        "message": "connection added"
    }
    ```
### /api/users/:username/get/connection/:id `GET`
To get the user connection with id.

- Response
    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxlZ2F0by"
    }
    ```


### /api/user/connection/access/token/:service `GET`
To get the access url for apps that  you would like to connect.
- Response
    ```json
    {
        "spotify_url": "https://discord.com/api/oauth2/authorize?access_type=online&client_id=830463353079988314&redirect_uri=http://localhost:8080/callback&response_type=code&scope=identify+email&state=h8EecvhXJqHsG5EQ3K0gei4EUrWpaFj_HqH3WNZdrzrX1BX1COQRsTUv3-yGi3WmHQbw0EHJ58Rx1UOkvwip-Q%3D%3D"
    }
    ```
  
> `token_type` are spotify, google, git, discord.

### /api/users/:username/get/connections `GET`
To get the all  connections of a user as list .

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

### /api/users/:username/update/connection/name/:id `PUT`
To update the name of a connection with id.
- Request
    ```json
    {
       "name": "git1"
    }
    ```
- Response
    ```json
    {
        "message":  "updated connection successfully"
    }
    ```

### /api/"users/:username/check/connection/:id `GET`
To get the id of a connection and check if there is a connection with this id for a user return correct token
- Response
    ```json
    {
        "message": "correct connection"
    }
    ```

### /api/users/:username/update/connection/token/:id `PUT`
To update field token in connection with id.
- Request
    ```json
    {
        "token" : "asdpjbfipdsafhbcop;bfdhsbocp;x"
    }
    ```
- Response
    ```json
    {
        "message":  "updated connection successfully"
    }
    ```


### /api/users/:username/connection/delete/:id `DELETE`
To delete the connection with id.
- Response
    ```json
    {
        "message":  "deleted connection successfully"
    }
    ```