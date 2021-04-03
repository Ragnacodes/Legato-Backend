# Scenario

### /api/users/:username/scenarios `GET`
To get all the user scenarios.
- Header
    - `Authorization` = `access_token`
- Response
    ```json
    [
        {
            "id": 1,
            "name": "my favorite scenario2",
            "is_active": true,
            "nodes": []
        },
        {
            "id": 3,
            "name": "amin scenario",
            "is_active": false,
            "nodes": []
        },
        {
            "id": 14,
            "name": "my favorite scenario",
            "is_active": true,
            "nodes": []
        }
    ]
    ```

### /api/users/:username/scenarios `POST`
To create new scenarios.
- Request
    ```json
    {
        "name": "test title for scenario",
        "is_active": true
    }
    ```
- Response
    ```json
    {
        "message": "scenario created successfully.",
        "scenario": {
          "id": 22,
          "name": "test title for scenario",
          "is_active": true,
          "nodes": []
        }
    }
    ```

### /api/users/:username/scenarios/:scenario_id `GET`
To get all the scenario details including the service tree.
- Header
    - `Authorization` = `access_token`
 - Response
     ```json
    {
        "id": 1,
        "name": "my favorite scenario2",
        "is_active": true,
        "graph": {
            "name": "My initial webhook",
            "type": "webhook",
            "children": [
                {
                    "name": "Event 1",
                    "type": "http",
                    "children": [],
                    "data": {}
                },
                {
                    "name": "Event 2",
                    "type": "http",
                    "children": [
                        {
                            "name": "First Http child",
                            "type": "http",
                            "children": [],
                            "data": {}
                        },
                        {
                            "name": "Event 4",
                            "type": "http",
                            "children": [],
                            "data": {}
                        },
                        {
                            "name": "Event 5",
                            "type": "http",
                            "children": [],
                            "data": {}
                        }
                    ],
                    "data": {}
                }
            ],
            "data": {}
        }
    }
     ```
   
### /api/users/:username/scenarios/:scenario_id `POST`
To update user scenario.
- Header
    - `Authorization` = `access_token`
- Request
    ```json
    {
        "name": "my test scenario",
        "is_active": true,
        "graph": {
            "name": "Webhook",
            "type": "webhook",
            "children": [
                {
                    "name": "Event 1",
                    "type": "http",
                    "children": [],
                    "data": {}
                },
                {
                    "name": "Event 2",
                    "type": "http",
                    "children": [],
                    "data": {}
                }
            ],
            "data": {}
        }
    }
    ```
- Response
    ```json
    {
        "message": "update scenario successfully",
        "scenario": {
            "id": 16,
            "name": "my test scenario",
            "is_active": true,
            "graph": {
                "name": "Webhook",
                "type": "webhook",
                "children": [
                    {
                        "name": "Event 1",
                        "type": "http",
                        "children": [],
                        "data": {}
                    },
                    {
                        "name": "Event 2",
                        "type": "http",
                        "children": [],
                        "data": {}
                    }
                ],
                "data": {}
            }
        }
    }
    ```

> **Note:** When a scenario has been created, the `graph` field should be null. 

Just pass the changes. For example if you just want to update scenario name the json should look like:
- Request
    ```json
    {
        "name": "another name for scenario"
    }
    ```
- Response
    ```json
    {
        "message": "update scenario successfully",
        "scenario": {
            "id": 16,
            "name": "another name for scenario",
            "is_active": true,
            "graph": null
        }
    }
    ```
> **Note:** You are not allowed to pass `"graph": null`. It doesn't make any changes.
