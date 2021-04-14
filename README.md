# Legato-Backend

### [Documentation](docs/README.md)
- How the APIs works.
- Default models

## Run Legato server

### Run the server for the first time
- Make sure you have installed the docker on your operating system.
- Clone the repository

    ```shell script
    $ git clone https://github.com/Ragnacodes/Legato-Backend
    ```
- Open terminal and enter this command to run server.

    ```shell script
    $ docker-compose -f docker-compose-dev.yml up --build
    ```
    > **NOTE:** Do not forget `sudo` if you are using linux base operating system.

### Run the server in developing mode
- Run it with just one command

    ```shell script
    $ docker-compose -f docker-compose-dev.yml up
    ```
    > **NOTE:** Do not forget `sudo` if you are using linux base operating system.


### Update server and run the server
- Pull the repository to get new files.

    ```shell script
    $ git pull origin master
    ```
- Run server with a flag `--build` for the first time.

    ```shell script
    $ docker-compose -f docker-compose-dev.yml up --build
    ```