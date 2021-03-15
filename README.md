# Legato-Backend
## Run the server for the first time
1. Make sure you have installed the docker on your operating system.
2. Clone the repository
    ```shell script
    $ git clone https://github.com/Ragnacodes/Legato-Backend
    ```
3. Open terminal and enter this command to run server.
    ```shell script
    $ docker-compose -f docker-compose-dev.yml up --build
    ```
    > **NOTE:** Do not forget `sudo` if you are using linux base operating system.

## Run the server in developing mode
Run it with just one command
```shell script
$ docker-compose -f docker-compose-dev.yml up
```
> **NOTE:** Do not forget `sudo` if you are using linux base operating system.


## Update server and run the server
1. Pull the repository to get new files.
    ```shell script
    $ git pull origin master
    ```
2. Run server with a flag `--build` for the first time.
    ```shell script
    $ docker-compose -f docker-compose-dev.yml up --build
    ```