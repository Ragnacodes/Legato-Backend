# Logs
 
### sse stream channel > /api/events/:scid 


### /api/users/:username/logs/:scenario_id/ `GET`
get list of histories of a scenario 
- Header
    - `Authorization` = `access_token`

- Response
    200 ok
    ```json
    {
    "histories": [
            {
                "id": 1,
                "created_at": "2021-05-28T21:52:32+0000"
            },
            {
                "id": 2,
                "created_at": "2021-05-28T22:36:41+0000"
            },
            {
                "id": 3,
                "created_at": "2021-05-28T22:38:38+0000"
            }
        ]
    }
    ```


### /api/users/:username/logs `GET`
Get a list of recent histories with their scenario object.
- Header
    - `Authorization` = `access_token`

- Response
    200 ok
    ```json
    {
      "histories": [
        {
          "scenario": {
            "id": 3,
            "name": "scenario 22 ",
            "interval": 0,
            "lastScheduledTime": "0001-01-01T00:00:00Z",
            "isActive": true,
            "nodes": [
              "https"
            ]
          },
          "history": {
            "id": 2,
            "created_at": "2021-07-11T04:50:58+0000",
            "scenario_id": 3
          }
        },
        {
          "scenario": {
            "id": 1,
            "name": "first scenario",
            "interval": 14,
            "lastScheduledTime": "2021-05-13T22:09:41Z",
            "isActive": true,
            "nodes": [
              "https",
              "spotifies",
              "telegrams",
              "webhooks"
            ]
          },
          "history": {
            "id": 1,
            "created_at": "2021-07-11T04:30:11+0000",
            "scenario_id": 1
          }
        }
      ]
    }

    ```


### /api/users/:username/logs/:scenario_id/histories/:history_id `GET`
get list of log messages of a history log 
- Header
    - `Authorization` = `access_token`

- Response
    200 ok
    ```json
    {
        "logs": [
            {
                "id": 7,
                "Messages": [
                    {
                        "context": "Make http request"
                    },
                    {
                        "context": "\nurl: https://abstergo.ir/api/ping\nmethod: get\nbody:\n{}\n"
                    },
                    {
                        "context": "Response from http request is : \n{\"message\":\"pong\"}\n"
                    }
                ],
                "Service": {
                    "id": 15,
                    "parentId": 12,
                    "name": "hamintori",
                    "type": "https",
                    "subType": "",
                    "position": {
                        "x": 40,
                        "y": 40
                    },
                    "data": {
                        "body": {},
                        "method": "get",
                        "url": "https://abstergo.ir/api/ping"
                    }
                }
            },
            {
                "id": 8,
                "Messages": [
                    {
                        "context": "webhook with id 3fbe716d-9f38-489e-88e3-ab96066719de got payload:"
                    },
                    {
                        "context": "{}"
                    },
                    {
                        "context": "Executing \"fuck\" Children \n"
                    },
                    {
                        "context": "*******End of \"fuck\"*******"
                    }
                ],
                "Service": {
                    "id": 12,
                    "parentId": null,
                    "name": "fuck",
                    "type": "webhooks",
                    "subType": "",
                    "position": {
                        "x": 0,
                        "y": 0
                    },
                    "data": {
                        "id": 6,
                        "isEnable": true,
                        "url": "3fbe716d-9f38-489e-88e3-ab96066719de"
                    }
                }
            }
        ]
    }
    ```