# SSH
- Type = `sshes`

### Data
Adding ssh to scenario. There is just one sub type yet.
- Data request to create:
    ```json
    {
        "parentId": null,
        "name": "hamintori",
        "type": "sshes",
        "subType": "username and password",
        "position": {
            "x": 100,
            "y": 100
        },
        "data": {
          "host" :"37.152.181.64",
          "username":"reza",
          "password":"sko192j3h",
          "commands":["ls","echo hello world"]
        }
    }
    ```

- Data response
    ```json
    {
        "message": "node is created successfully.",
        "node": {
            "id": 9,
            "parentId": null,
            "name": "hamintori",
            "type": "sshes",
            "subType": "username and password",
            "position": {
                "x": 100,
                "y": 100
            },
            "data": {
                "commands": [
                    "ls",
                    "echo hello world"
                ],
                "host": "37.152.181.64",
                "password": "sko192j3h",
                "username": "reza"
            }
        }
    }
    ```


### /api/users/:username/services/sshes `GET`
To get the all  connections of a user as list .

- Header
    - `Authorization` = `access_token`
    
 - Response
     ```json
     {
         "Sshes": [
            {
                "id": 1,
                "username": "reza",
                "host": "37.152.181.64",
                "password": "aaaaa",
                "sshkey": "",
                "connectionid": 25
            },
            {
                "id": 2,
                "username": "reza",
                "host": "37.152.181.64",
                "password": "aaaaaaa",
                "sshkey": "",
                "connectionid": 24
            },
            {
                "id": 3,
                "username": "reza",
                "host": "37.152.181.64",
                "password": "aaaaaaa",
                "sshkey": "",
                "connectionid": 0
            }
        ]
     }
     ```
    

### /api/check/ssh/:type `POST`

- Header
    - `Authorization` = `access_token`
 - Data request to create:
    ```json
    {
        "host":"37.152.181.64",
        "username":"reza",
        "password":"password"
    }
    ```
 - Response
     ```json
    {
       "massage": "OK"
    }
    ```