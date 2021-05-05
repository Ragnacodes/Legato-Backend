# Spotify
- Type = `spotifies`

### SubTypes & Data
When adding spotify node
- ```addToPlaylist```
    - Data request to create:
    
        ```json
        {
            "TrackId" : "ersgrt53mSr32Qq0srgya",  
            "PlaylistId" : "5vC3aemdQGzV8Dq0eroyDa"
        }
        ``` 
    
- ```getTopTracks```
   

### /api/services/spotify/track/:trackid `GET`
Get Spotify catalog information for a single track identified by its unique Spotify ID. 
- Header
    not required

- Response
    200 ok
    ```json
    {
        "artists": [
        {
            "name": "Archive",
            "id": "1Q776wzj2mrtXrNu3iH6nk",
            "uri": "spotify:artist:1Q776wzj2mrtXrNu3iH6nk",
            "href": "https://api.spotify.com/v1/artists/1Q776wzj2mrtXrNu3iH6nk",
            "external_urls": {
                "spotify": "https://open.spotify.com/artist/1Q776wzj2mrtXrNu3iH6nk"
            }
        }
    ],
    "available_markets": [
        "AD",
        "...",
        "ZW"
    ],
    "disc_number": 1,
    "duration_ms": 978626,
    "explicit": false,
    "external_urls": {
        "spotify": "https://open.spotify.com/track/25XKxXQy5AhBZ2uOizLQGw"
    },
    "href": "https://api.spotify.com/v1/tracks/25XKxXQy5AhBZ2uOizLQGw",
    "id": "25XKxXQy5AhBZ2uOizLQGw",
    "name": "Again",
    "preview_url": "https://p.scdn.co/mp3-preview/23fa1f92d96d6b14d73f8fbe966ba4ddb5210d1f?cid=c42e107d3ae641e4af9e08e7d7a55b9b",
    "track_number": 1,
    "uri": "spotify:track:25XKxXQy5AhBZ2uOizLQGw",
    "album": {
        "name": "You All Look the Same to Me",
        "artists": [
            {
                "name": "Archive",
                "id": "1Q776wzj2mrtXrNu3iH6nk",
                "uri": "spotify:artist:1Q776wzj2mrtXrNu3iH6nk",
                "href": "https://api.spotify.com/v1/artists/1Q776wzj2mrtXrNu3iH6nk",
                "external_urls": {
                    "spotify": "https://open.spotify.com/artist/1Q776wzj2mrtXrNu3iH6nk"
                }
            }
        ],
        "album_group": "",
        "album_type": "album",
        "id": "5YYiF8dLVo9jt0qBY2zyjz",
        "uri": "spotify:album:5YYiF8dLVo9jt0qBY2zyjz",
        "available_markets": [
            "AD",
            "...",
            "ZW"
        ],
        "href": "https://api.spotify.com/v1/albums/5YYiF8dLVo9jt0qBY2zyjz",
        "images": [
            {
                "height": 640,
                "width": 640,
                "url": "https://i.scdn.co/image/ab67616d0000b2735bceb0c8a1059da0f4257956"
            },
            {
                "height": 300,
                "width": 300,
                "url": "https://i.scdn.co/image/ab67616d00001e025bceb0c8a1059da0f4257956"
            },
            {
                "height": 64,
                "width": 64,
                "url": "https://i.scdn.co/image/ab67616d000048515bceb0c8a1059da0f4257956"
            }
        ],
        "external_urls": {
            "spotify": "https://open.spotify.com/album/5YYiF8dLVo9jt0qBY2zyjz"
        },
        "release_date": "2002-01-01",
        "release_date_precision": "day"
    },
    "external_ids": {
        "isrc": "GBEYA0100002"
    },
    "popularity": 42,
    "is_playable": null,
    "linked_from": null
    }
    ```

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
    }
    ```
  


### /api/users/:username/spotify/playlists `GET`
****To get list of all user playlists. 
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
