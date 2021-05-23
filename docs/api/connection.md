### /api/users/:username/add/connection `POST`
To create a new connection.

- Request
    ```json
    {
        "data": {
            "host": "37.152.181.64",
            "password": "sko192j3h",
            "username": "reza"
         },
        "name" :"ssh5"
    }
    ```
- Response
    ```json
    {
        "Data": {
            "host": "37.152.181.64",
            "password": "sko192j3h",
            "username": "reza"
        },
        "Id": 26,
        "Name": "ssh5"
    }
    ```
### /api/users/:username/get/connection/:id `GET`
To get the user connection with id.

- Response
    ```json
    {
        "Data": {
            "host": "37.152.181.64",
            "password": "sko192j3h",
            "username": "reza"
        },
        "Name": "ssh5"
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
    "connections": [
        {
            "id": 32,
            "data": {
                "host": "37.152.181.64",
                "password": "aaaaaaaa",
                "username": "reza"
            },
            "name": "ssh8",
            "type": "sshes"
        },
        {
            "id": 34,
            "data": {
                "host": "37.152.181.64",
                "password": "ddddddddddddddddddd",
                "username": "reza"
            },
            "name": "ssh9",
            "type": "sshes"
        },
        {
            "id": 35,
            "data": {
                "host": "37.152.181.64",
                "sshKey": "sko192j3h",
                "username": "reza"
            },
            "name": "ssh10",
            "type": "sshes"
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

### /api/users/:username/check/connection/:id `GET`
To get the id of a connection and check if there is a connection with this id for a user return true or false
- Response
    ```json
    {
        "message": "true"
    }
    ```

### /api/users/:username/update/connection/token/:id `PUT`
To update field data in connection with id. send data as a string with this fromat 
- Request
    ```json
    {
        "data" : {"data":{"host":"37.152.181.64","password":"sko192j3h","username":"reza"},"id":2,"name":"ssh1"}
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