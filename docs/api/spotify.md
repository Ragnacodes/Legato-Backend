# Spotify

### SubType
When adding spotify node
- "addToPlaylist"
   Data:
    ```json
    {
        "TrackId" : "ersgrt53mSr32Qq0srgya",  
        "PlaylistId" : "5vC3aemdQGzV8Dq0eroyDa"
    }
    ``` 
    
- "getTopTracks"
    Data: empty
   


### api/users/:username/spotify/ `GET`
To login to your spotify account
> after you send this request a link to spotify, would be in server log and
     request will be pending until you compelete it, click on link and after 
     getting permission, request would be completed.

- Response
    200 ok
    ```json
    {
    "message": "You are logged in as: spotifyUsername"
    }```


### /api/users/:username/spotify/playlists `GET`
To get list of all user webhooks. 
- Header
    - `Authorization` = `access_token`

- Response
    ```json
    [
        {
            "collaborative": false,
            "external_urls": {
                "spotify": "https://open.spotify.com/playlist/5vC3aemdQGzV8Dq0eroyDa"
            },
            "href": "https://api.spotify.com/v1/playlists/5vC3aemdQGzV8Dq0eroyDa",
            "id": "5vC3aemdQGzV8Dq0eroyDa",
            "images": [],
            "name": "The Bitter Truth",
            "owner": {
                "display_name": "persianrobo4",
                "external_urls": {
                    "spotify": "https://open.spotify.com/user/persianrobo4"
                },
                "followers": {
                    "total": 0,
                    "href": ""
                },
                "href": "https://api.spotify.com/v1/users/persianrobo4",
                "id": "persianrobo4",
                "images": null,
                "uri": "spotify:user:persianrobo4"
            },
            "public": true,
            "snapshot_id": "MixlOTAzNWYwNTQ1M2EwZTA1MTI4Mjk2NDZmMmMwNGFhYjk0NzNlMTFm",
            "tracks": {
                "href": "https://api.spotify.com/v1/playlists/5vC3aemdQGzV8Dq0eroyDa/tracks",
                "total": 12
            },
            "uri": "spotify:playlist:5vC3aemdQGzV8Dq0eroyDa"
        },
        {
            "collaborative": false,
            "external_urls": {
                "spotify": "https://open.spotify.com/playlist/3dZAxWxyihbk2e2mNFTR7T"
            },
            "href": "https://api.spotify.com/v1/playlists/3dZAxWxyihbk2e2mNFTR7T",
            "id": "3dZAxWxyihbk2e2mNFTR7T",
            "images": [
                {
                    "height": 640,
                    "width": 640,
                    "url": "https://i.scdn.co/image/ab67616d0000b273d9ac65941679b76106866c87"
                }
            ],
            "name": "Gg",
            "owner": {
                "display_name": "persianrobo4",
                "external_urls": {
                    "spotify": "https://open.spotify.com/user/persianrobo4"
                },
                "followers": {
                    "total": 0,
                    "href": ""
                },
                "href": "https://api.spotify.com/v1/users/persianrobo4",
                "id": "persianrobo4",
                "images": null,
                "uri": "spotify:user:persianrobo4"
            },
            "public": true,
            "snapshot_id": "MixlNTAxN2NjMmJlYTM4ZTA2ODNmNWRhZTM0ZmEyZDMzNDcwOTlhMGRh",
            "tracks": {
                "href": "https://api.spotify.com/v1/playlists/3dZAxWxyihbk2e2mNFTR7T/tracks",
                "total": 1
            },
            "uri": "spotify:playlist:3dZAxWxyihbk2e2mNFTR7T"
        },
        {
            "collaborative": false,
            "external_urls": {
                "spotify": "https://open.spotify.com/playlist/39busjRx7Hfu5PLHBqlE40"
            },
            "href": "https://api.spotify.com/v1/playlists/39busjRx7Hfu5PLHBqlE40",
            "id": "39busjRx7Hfu5PLHBqlE40",
            "images": [
                {
                    "height": 640,
                    "width": 640,
                    "url": "https://mosaic.scdn.co/640/ab67616d0000b273459d675aa0b6f3b211357370ab67616d0000b273574860379dd3dd615ec3bb7bab67616d0000b2737b8aabae10ab5bbe7c7f11c5ab67616d0000b273c052bad6a067197de0fb95a1"
                },
                {
                    "height": 300,
                    "width": 300,
                    "url": "https://mosaic.scdn.co/300/ab67616d0000b273459d675aa0b6f3b211357370ab67616d0000b273574860379dd3dd615ec3bb7bab67616d0000b2737b8aabae10ab5bbe7c7f11c5ab67616d0000b273c052bad6a067197de0fb95a1"
                },
                {
                    "height": 60,
                    "width": 60,
                    "url": "https://mosaic.scdn.co/60/ab67616d0000b273459d675aa0b6f3b211357370ab67616d0000b273574860379dd3dd615ec3bb7bab67616d0000b2737b8aabae10ab5bbe7c7f11c5ab67616d0000b273c052bad6a067197de0fb95a1"
                }
            ],
            "name": "emtehana",
            "owner": {
                "display_name": "persianrobo4",
                "external_urls": {
                    "spotify": "https://open.spotify.com/user/persianrobo4"
                },
                "followers": {
                    "total": 0,
                    "href": ""
                },
                "href": "https://api.spotify.com/v1/users/persianrobo4",
                "id": "persianrobo4",
                "images": null,
                "uri": "spotify:user:persianrobo4"
            },
            "public": true,
            "snapshot_id": "MTIsNjU0MmE0NTYxOTlmNGM1ODQ2ODE0N2MxMGYwNDk2ZDFjYmRjMjBmYg==",
            "tracks": {
                "href": "https://api.spotify.com/v1/playlists/39busjRx7Hfu5PLHBqlE40/tracks",
                "total": 11
            },
            "uri": "spotify:playlist:39busjRx7Hfu5PLHBqlE40"
        }
    ]
```