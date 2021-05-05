# Telegram
- Type = `telegrams`

### SubTypes & Data
Adding telegram node to the scenario.
- `sendMessage` (Send a message to the user)
    - Data request to create:
    
        ```json
        {
            "data": {
              "key": "17312061423:AAHkpaaUswee",
              "chat_id": "128373",
              "text": "this is the message!"
            }
        }
        ```
    
    - Data response
        ```json
        {
            "data": {
              "key": "17312061423:AAHkpaaUswee",
              "chat_id": "128373",
              "text": "this is the message!"
            }
        }
        ```
  
- `getChatMember` (Get user info)
    - Data request to create:
        ```json
        {
            "data": {
              "chat_id": "128373",
              "user_id": "128373"
            }
        }
        ```
    
    - Data response
        ```json
        {
            "data": {
              "chat_id": "128373",
              "user_id": "128373"
            }
        }
        ```