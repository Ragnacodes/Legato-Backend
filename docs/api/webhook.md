# Webhook

### /users/:username/services/webhooks `POST`
To create a new separate webhook.
- Header
    - `Authorization` = `access_token`
    
- Request
    ```json
    {
        "name" : "my first webhook"
    }
    ```

- Response
    ```json
    {
        "message": "webhook is added successfully",
        "webhook": {
            "url": "http://localhost:8080/api/services/webhook/d8fe28b9-738c-4d00-9630-a4492db87719",
            "name": "my lovely webhook",
            "active": false
        }
    }
    ```

### /api/services/webhook/:uuid `POST`
To send request to webhook endpoint.
> Notice that you should enable your webhook before using this api
- Request
    any arbitrary data

- Response
    200 ok

### /api/users/:username/services/webhooks/:webhook_id `PUT`
To update the webhook.
- Header
    - `Authorization` = `access_token`

- Request
   ```json
    {
        "name": "my webhook",
        "isEnable": false
    }
    ```
  
- Response
    ```json
    {
        "message": "webhook is updated successfully",
        "webhook": {
            "url": "http://localhost:8080/api/services/webhook/255d24ba-eef5-4968-bd38-2c35fd5cdaec",
            "name": "my webhook",
            "isEnable": false
        }
    }
    ```

### /api/users/:username/services/webhooks/:webhook_id `GET`
To get list of all user webhooks. 
- Header
    - `Authorization` = `access_token`
    
- Response 
    ```json
    {
        "webhook": {
            "url": "http://localhost:8080/api/services/webhook/255d24ba-eef5-4968-bd38-2c35fd5cdaec",
            "name": "my lovely webhook",
            "isEnable": false
        }
    }
    ```

### /api/users/:username/services/webhooks `GET`
To get list of all user webhooks. 
- Header
    - `Authorization` = `access_token`

- Response
    ```json
    {
        "webhooks": [
            {
                "url": "http://localhost:8080/api/services/webhook/0d7ecca0-55b7-402e-90e3-279a148ecd19",
                "name": "another http",
                "active": false
            },
            {
                "url": "http://localhost:8080/api/services/webhook/cee118b5-8999-4347-a025-0065e9fa9cd7",
                "name": "another http",
                "active": false
            },
            {
                "url": "http://localhost:8080/api/services/webhook/a1f42e4a-06f0-4967-8fde-20c11ab8f714",
                "name": "My second http",
                "active": false
            },
            {
                "url": "http://localhost:8080/api/services/webhook/83796a75-e1f1-4579-b716-6d020caf3845",
                "name": "My starter webhook",
                "active": false
            },
            {
                "url": "http://localhost:8080/api/services/webhook/a4ed49bc-62e3-48f0-8b0c-6c503ebb1538",
                "name": "My first http",
                "active": false
            },
            {
                "url": "http://localhost:8080/api/services/webhook/7bc5ffb1-0630-46d2-b2b2-07a4c9f795b0",
                "name": "another http",
                "active": false
            },
            {
                "url": "http://localhost:8080/api/services/webhook/1048ba09-d9e1-4c15-ac5d-69e66ef7c4ac",
                "name": "My second http",
                "active": false
            },
            {
                "url": "http://localhost:8080/api/services/webhook/54fb2628-73ce-460f-98b2-eb8e9636c484",
                "name": "My starter webhook",
                "active": false
            },
            {
                "url": "http://localhost:8080/api/services/webhook/b47eab40-b2d9-48cc-b0e6-c9b8250aa5d1",
                "name": "My first http",
                "active": false
            },
            {
                "url": "http://localhost:8080/api/services/webhook/3cc5eae1-344a-40c1-844b-3c3047c8431e",
                "name": "another http",
                "active": false
            }
        ]
    }
    ```