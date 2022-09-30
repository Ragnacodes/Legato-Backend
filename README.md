# Legato-Backend

Legato is a workflow system management. You can define your routines in a scenario and automate them. 

> 😄 In music performance and notation, legato indicates that musical notes are played or sung smoothly and connected.

All you need to determine is a graph. The nodes are the services or the tasks that you need to be done. The edges specify the parallelism of a series of tasks or the priority of a job over another. 

## Documentation

- [Authentication](docs/api/auth.md)
- [Scenario](docs/api/scenario.md)
- [Service](docs/api/service.md)
- [Webhook](docs/api/webhook.md)
- [Spotify](docs/api/spotify.md)
- [Connection](docs/api/connection.md)
- [Logs](docs/api/logs.md)
- [Https](docs/api/http.md)
- [Gmail](docs/api/gmail.md)
- [Github](docs/api/github.md)
- [Discord](docs/api/discord.md)


## Run the Legato server

```bash
# Clone the repository 
git clone https://github.com/Ragnacodes/Legato-Backend

# Run the docker compose 
sudo docker-compose up --build
```

## Preview

### Services

At the moment, there are about eight available services. There are also some other toolkits like sleep and loop. The point is that you can easily define your own service and add to this framework. 

There are a bunch of services that you can use.

![docs/images/Untitled%203.png](docs/images/Untitled%203.png)

### Panel

You can  use the provided panel to

- keep an eye on your workflows
- see the logs or status of each workflow
- see some statistics on the first page of the panel.

![docs/images/Untitled%204.png](docs/images/Untitled%204.png)
