# Webhook

### /api/services/webhook `POST`
To create a webhook.
- Header
    - `Authorization` = `access_token`
- Request
    {}
   
- Response
    ```json
    {
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

### /api/services/webhook/:uuid `PATCH`
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

