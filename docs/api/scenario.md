# Scenario

### /api/users/:username/scenarios `GET`
To get all the user scenarios.
- Header
    - `Authorization` = `access_token`
- Response
    ```json
    {
        "scenarios": [
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
    }
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
        "message": "scenario is created successfully.",
        "scenario": {
          "id": 22,
          "name": "test title for scenario",
          "is_active": true,
          "nodes": []
        }
    }
    ```

### /api/users/:username/scenarios/:scenario_id `PATCH`
To start the scenarios.
- Header
    - `Authorization` = `access_token`
    
- Response
    ```json
    {
        "message": "scenario is started successfully"
    }
    ```

### /api/users/:username/scenarios/:scenario_id/schedule `POST`
To schedule the scenarios. `systemTime` is the user's current time.
`scheduledTime` is the time the user set to schedule the scenario.
It is needed to send both of them. Consider the case that the server time
is not as same as the user. With having both of these fields we calculate
the duration `scheduledTime - systemTime`.
The given time format should be RFC3339.
`interval` is the period time to repeat starting scenario. Value 0 meant to
start the scenario once. 
- Header
    - `Authorization` = `access_token`
    
- Response
    ```json
    {
        "systemTime": "2021-05-14T02:16:00+04:30",
        "scheduledTime": "2021-05-14T03:00:04+04:30",
        "interval": 2
    }
    ```
  
- Response
    ```json
    {
        "message": "scenario is scheduled successfully for 2021-05-14 03:00:04 +0430 +0430"
    }
    ```
  
### /api/users/:username/scenarios/:scenario_id `DELETE`
To delete specific scenarios.
- Response
    ```json
    {
        "message": "scenario is deleted successfully."
    }
    ```

### /api/users/:username/scenarios/:scenario_id `GET`
To get all the scenario details including the services list.
- Header
    - `Authorization` = `access_token`
 - Response
     ```json
    {
        "scenario": {
            "id": 1,
            "name": "my favorite scenario2",
            "is_active": true,
            "interval": 20,
            "services": [
                {
                    "id": 21,
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
                    "id": 22,
                    "parentId": 21,
                    "userId": 1,
                    "name": "My first http",
                    "type": "webhooks",
                    "position": {
                        "x": 0,
                        "y": 0
                    },
                    "data": {}
                },
                {
                    "id": 23,
                    "parentId": 21,
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
                    "id": 24,
                    "parentId": 22,
                    "userId": 1,
                    "name": "My second http",
                    "type": "webhooks",
                    "position": {
                        "x": 0,
                        "y": 0
                    },
                    "data": {}
                }
            ]
        }
    }
     ```
   
### /api/users/:username/scenarios/:scenario_id `PUT`
To update user scenario.
- Header
    - `Authorization` = `access_token`
- Request
    ```json
    {
        "name": "my test scenario",
        "is_active": true
    }
    ```
- Response
    ```json
    {
        "message": "scenario i updated successfully",
        "scenario": {
            "id": 16,
            "name": "my test scenario",
            "is_active": true,
            "interval": 20,
            "services": [
                {
                    "id": 21,
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
                    "id": 22,
                    "parentId": 21,
                    "userId": 1,
                    "name": "My first http",
                    "type": "webhooks",
                    "position": {
                        "x": 0,
                        "y": 0
                    },
                    "data": {}
                },
                {
                    "id": 23,
                    "parentId": 21,
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
                    "id": 24,
                    "parentId": 22,
                    "userId": 1,
                    "name": "My second http",
                    "type": "webhooks",
                    "position": {
                        "x": 0,
                        "y": 0
                    },
                    "data": {}
                }
            ]
        }
    }
    ```
> **Note:** At first the services list should be an empty array.

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
        "message": "scenario is updated successfully",
        "scenario": {
            "id": 16,
            "name": "another name for scenario",
            "is_active": true,
            "services": []
        }
    }
    ```
