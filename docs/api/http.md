# Http
- Type = `https`

### SubTypes & Data
Adding http to scenario. There is not any sub type yet.
- Data request to create:
    ```json
    {
        "data": {
          "url": "http://localhost:8080/api/ping",
          "method": "post",
          "body": {
            "test": "test"
          }
        }
    }
    ```

- Data response
    ```json
    {
        "data": {
          "id": 23,
          "url": "http://localhost:8080/api/ping",
          "method": "post",
          "body": {
            "test": "test"
          }
        }
    }
    ```