# Services

### /api/users/:username/scenarios/:scenario_id/nodes `POST`
To add a service (node) to the scenario.
- Header
    - `Authorization` = `access_token`
    
- Request
    ```json
    {
        "parentId": 90,
        "name": "hamintori",
        "type": "webhooks",
        "subType": "",
        "position": {
            "x": 40,
            "y": 40
        },
        "data": {}
    }
    ```
        
- Response
    ```json
    {
        "message": "node is created successfully.",
        "node": {
            "id": 131,
            "parentId": 90,
            "name": "hamintori",
            "type": "webhooks",
            "subType": "",
            "position": {
                "x": 40,
                "y": 40
            },
            "data": {
                "url": "http://localhost:8080/api/services/webhook/1d96b814-2b6a-4e61-8360-d0580bbc332a",
                "active": false
            }
        }
    }
    ```

> Here are some services that can be created in scenarios.

#### Available Services
Data field for each one of services is different. See the service documentation for more details.
- [Webhook](webhook.md) 
- [Http](http.md)
- [Telegram](telegram.md)
- [Spotify](spotify.md)
- [SSH](ssh.md)
- [Discord](discord.md)
- [Tool Box](tool_box.md)
- [Github](github.md)

### /api/users/:username/scenarios/:scenario_id/nodes `GET`
To get all services (nodes) in that user scenarios.
- Header
    - `Authorization` = `access_token`
    
- Response
    ```json
    {
        "nodes": [
            {
                "id": 90,
                "parentId": 88,
                "userId": 1,
                "name": "another http",
                "type": "webhooks",
                "position": {
                    "x": 0,
                    "y": 0
                },
                "data": {}
            },
            {
                "id": 88,
                "parentId": null,
                "userId": 1,
                "name": "My starter webhook",
                "type": "webhooks",
                "position": {
                    "x": 0,
                    "y": 0
                },
                "data": {}
            },
            {
                "id": 91,
                "parentId": 88,
                "userId": 1,
                "name": "updated",
                "type": "webhooks",
                "position": {
                    "x": 40,
                    "y": 40
                },
                "data": {}
            },
            {
                "id": 97,
                "parentId": 92,
                "userId": 1,
                "name": "New http added 22333",
                "type": "webhooks",
                "position": {
                    "x": 123,
                    "y": 200
                },
                "data": {}
            },
            {
                "id": 92,
                "parentId": 88,
                "userId": 1,
                "name": "updated",
                "type": "webhooks",
                "position": {
                    "x": 40,
                    "y": 40
                },
                "data": {}
            }
        ]
    }
    ```

### /api/users/:username/scenarios/:scenario_id/nodes/:node_id `GET`
To get details about a single service (node) in that user scenario.
- Header
    - `Authorization` = `access_token`
    
- Response
    ```json
    {
        "node": {
            "id": 88,
            "parentId": null,
            "name": "My starter webhook",
            "type": "webhooks",
            "position": {
                "x": 0,
                "y": 0
            },
            "data": null
        }
    }
    ```

### /api/users/:username/scenarios/:scenario_id/nodes/:node_id `PUT`
To update a specific service (node) in that user scenarios.
> Only the given field will be changed.

- Header
    - `Authorization` = `access_token`
    
- Request
    ```json
    {
        "parentId": 88,
        "name": "updated",
        "position": {
            "x": 40,
            "y": 40
        }
    }
    ```
    
- Response
    ```json
    {
        "message": "node is updated successfully.",
        "node": {
            "id": 91,
            "parentId": 88,
            "name": "updated",
            "type": "webhooks",
            "position": {
                "x": 40,
                "y": 40
            },
            "data": null
        }
    }
    ```

### /api/users/:username/scenarios/scenario_id/nodes/:node_id `DELETE`
To delete a specific service (node) in that user scenarios.
> After deleting a node in the scenario, The parent of all its children would become the parent of deleted node.
> etc. If A was connected to B and B was connected to C after deleting the B, A would be connected to c, automatically.
- Header
    - `Authorization` = `access_token`
    
- Response
    ```json
    {
        "message": "node is deleted successfully"
    }
    ```