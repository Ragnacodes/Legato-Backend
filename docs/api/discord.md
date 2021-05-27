# Telegram
- Type = `discords`

### Connection
```json
{
    "type": "discord",
    "data": {
        "guildId": "844018666315711474"
    }
}
```

### SubTypes & Data
Adding discord node to the scenario.
- `sendMessage` (Send a message to the user)
    - Data request to create:
    
        ```json
        {
            "data": {
              "guildId": "12381902380192",
              "channelId": "128373",
              "content": "this is the message!"
            }
        }
        ```
    
    - Data response
        ```json
        {
            "data": {
              "guildId": "12381902380192",
              "channelId": "128373",
              "content": "this is the message!"
            }
        }
        ```
