# Webhook

### /users/:username/services/webhook `POST`
To create a webhook.
- Header
    - `Authorization` = `access_token`
- Request
    {
        "name" : "my first webhook"
    }
   
- Response
    ```json
    {
        "enable" : false,
        "name" : "my first webhook",
        "url": "http://127.0.0.1:8080/api/services/webhook/8bb33bc6-6957-4ded-b448-f12a52e613de",
    }
    ```

### /api/services/webhook/:uuid `POST`
To send request to webhook endpoint.
### Notice that you should enable your webhook before using this api
- Request
    any arbitrary data

- Response
    200 ok

### /users/:username/services/webhook/:webhookid `PATCH`
to make your webhook active or change its name, this is the api to update your webhook.
- Header
    - `Authorization` = `access_token`

- Request
   ```json
    {
        "enable": "true",
        "name" : "af"
    }
    ```
- Response
    200 ok
    ```json
    {
       "message": "updated successfully"
    }
    ```
    
### /users/:username/services/webhook/
to get list of all user webhooks 
- Header
    - `Authorization` = `access_token`

- Response
    200 ok
    ```json
    [
        {
            "url": "http://localhost:8080/api/services/webhook/4e688e25-51b0-437b-b467-e271646ca83e",
            "name": "my first webhook",
            "active": false
        },
        {
            "url": "http://localhost:8080/api/services/webhook/d029ea0c-75b1-4be7-ac23-f275d01bb11a",
            "name": "my second webhook",
            "active": false
        },
        {
            "url": "http://localhost:8080/api/services/webhook/ab1aa571-052a-4bba-990a-855d2e43acdd",
            "name": "my third webhook",
            "active": true
        },
    ]
    ```