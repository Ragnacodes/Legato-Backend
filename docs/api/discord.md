# Discord
- Type = `discords`

> **Note**: You should create an env file in /.
```dotenv
DISCORD_BOT_SECRET=<ele>
```

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
### /api/services/discord/guilds/:guildId/channels/text `GET`
Returns the text channels of a guild.
- Response
    ```json
    {
        "channels": [
            {
                "id": "844018666315710479",
                "guild_id": "844018666315710476",
                "name": "general",
                "topic": "",
                "type": 0,
                "last_message_id": "847540869434703932",
                "nsfw": false,
                "icon": "",
                "position": 0,
                "bitrate": 0,
                "user_limit": 0,
                "parent_id": "844018666315710477",
                "rate_limit_per_user": 0,
                "owner_id": "",
                "application_id": ""
            },
            {
                "id": "845633435011514369",
                "guild_id": "844018666315710476",
                "name": "groovy",
                "topic": "",
                "type": 0,
                "last_message_id": "846502029579780117",
                "nsfw": false,
                "icon": "",
                "position": 2,
                "bitrate": 0,
                "user_limit": 0,
                "parent_id": "844018666315710477",
                "rate_limit_per_user": 0,
                "owner_id": "",
                "application_id": ""
            },
            {
                "id": "846051939149807666",
                "guild_id": "844018666315710476",
                "name": "temp",
                "topic": "",
                "type": 0,
                "last_message_id": "847540868630315049",
                "nsfw": false,
                "icon": "",
                "position": 3,
                "bitrate": 0,
                "user_limit": 0,
                "parent_id": "844018666315710477",
                "rate_limit_per_user": 0,
                "owner_id": "",
                "application_id": ""
            }
        ]
    }
    ```