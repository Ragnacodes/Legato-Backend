# Logs
 
### sse stream channel > /api/events/:scid 


### /api/:username/logs/:scenario_id/ `GET`
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



### /api/:username/logs/:scenario_id/histories/:history_id `GET`
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
                },
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