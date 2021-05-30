# Githubs
- Type = `githubs`

### SubTypes & Data
Adding githubs node to the scenario.
- `createPullRequest`
    - Data request to create:
    
        ```json
        {
            "parentId": null,
            "name": "mygithub",
            "type": "githubs",
            "subType": "createPullRequest",
            "position": {
            "x": 100,
                "y": 100
            },
            "data": {
            "repositoryName":"eeeee",
            "title":"merge to develop",
            "base":"main",
            "head":"develop",
            "owner":"rezamnkh79",
            "body":"develop"
            }
	    }
        ```
    
    - Data response
        ```json
        {
            "message": "node is created successfully.",
            "node": {
                "id": 81,
                "parentId": null,
                "name": "mygithub",
                "type": "githubs",
                "subType": "createPullRequest",
                "position": {
                    "x": 100,
                    "y": 100
                },
                "data": {
                    "base": "main",
                    "body": "develop",
                    "head": "develop",
                    "owner": "rezamnkh79",
                    "repositoryName": "eeeee",
                    "title": "merge to develop",
                    "token": {
                        "access_token": "--------------------------------",
                        "expiry": "0001-01-01T00:00:00Z",
                        "token_type": "bearer"
                    }
                }
            }
        }
        ```
  
- `createIssue`
    - Data request to create:
        ```json
        {
            "parentId": null,
            "name": "mygit",
            "type": "githubs",
            "subType": "createIssue",
            "position": {
            "x": 100,
                "y": 100
            },
            "data": {
                "repositoryName":"eeeee",
                "body":"hello",
                "title":"ffffff",
                "owner":"rezamnkh79",
                "labels" :["bug","invalid"],
                "assignee" :["rezamnkh79"],
                "state":"open"
            }
        }
        ```
    
    - Data response
        ```json
        {
            "message": "node is created successfully.",
            "node": {
                "id": 84,
                "parentId": null,
                "name": "gititititit",
                "type": "githubs",
                "subType": "createIssue",
                "position": {
                    "x": 100,
                    "y": 100
                },
                "data": {
                    "assignee": [
                        "rezamnkh79"
                    ],
                    "body": "hello",
                    "labels": [
                        "bug",
                        "invalid"
                    ],
                    "owner": "rezamnkh79",
                    "repositoryName": "eeeee",
                    "state": "open",
                    "title": "ffffff",
                    "token": {
                        "access_token": "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",
                        "expiry": "0001-01-01T00:00:00Z",
                        "token_type": "bearer"
                    }
                }
            }
        }
        ```

### /api/users/:username/services/github/branches `POST`
Return list of branches in a repository.

- Request
    ```json
    {
    "connectionId" :1,
    "repositoryName" :"rezamnkh79/eeeee"

}   
    ```
- Response
    ```json
    {
        "branches_name": [
            "develop",
            "main"
        ]
    }
    ```

### /api/users/:username/services/github/repositories `POST`
Return list of branches in a repository.

- Request
    ```json
    {
    "connectionId" :1,  
    ```
- Response
    ```json
    {
        "repositoriesName": [
            "armanheydari/UefaChampionsLeague_DB",
            "Cypherspark/TunePal",
            "Ragnacodes/Legato-Backend",
            "Ragnacodes/Legato-Frontend",
        ]
    }   
    ```