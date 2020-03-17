# ACM

---

ACM's Website Mono-Repository. Inside contains submodules for all micro-frontend and micro-backend apis

## Getting Started

1. `git submodule init`
2. `git submodule update`
3. The whole project runs on multiple docker containers on a shared network. You can find all off the development yaml files in the `/docker` folder

## Run the Project with Docker Compose

Below is a list of possible combinations to run the ACM Website locally

### Website with all API Services

1. Run the root yaml compose file `docker-compose up -d`
2. To create all tables to store data on run the bash script with this command `chmod +x database_setup.sh && ./database_setup.sh`

> Note: When running the project this way the Chapter Website opens on `http://127.0.0.1:4200/home`

To check if they are setup correctly:

1. Docker Compose

    ```bash
    $ docker-compose ps
         Name                    Command              State            Ports
    ---------------------------------------------------------------------------------
    Chapter-Website   yarn start                      Up      0.0.0.0:4200->4200/tcp
    Half-Dome         yarn start                      Up      0.0.0.0:4201->4201/tcp
    postgres          docker-entrypoint.sh postgres   Up      0.0.0.0:35432->5432/tcp
    ```

2. Docker

    ```bash
    $ docker ps
    CONTAINER ID        IMAGE                  COMMAND                  CREATED              STATUS              PORTS                     NAMES
    87f43900d759        chapter-website:v0.1   "yarn start"             About a minute ago   Up About a minute   0.0.0.0:4200->4200/tcp    Chapter-Website
    d4b918fab0bc        half-dome:v1.1         "yarn start"             About a minute ago   Up About a minute   0.0.0.0:4201->4201/tcp    Half-Dome
    f7b512d5f823        postgres:12.2-alpine   "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:35432->5432/tcp   postgres
    ```

### Debugging Issues

#### Containers not Updating with New Code

1. Recreate Containers

    ```bash
    $ docker-compose up --force-recreate
    Recreating postgres ... done
    Recreating Half-Dome ... done
    Recreating Chapter-Website ... done
    Attaching to postgres, Half-Dome, Chapter-Website
    ...
    ```

2. Rebuild Containers

    ```bash
    $ docker-compose up --build
    Building half-dome
    ...
    Successfully built ce4b7e44d100
    Successfully tagged docker_half-dome:latest
    Building chapter-website
    ...
    Successfully built 6955801be2ec
    Successfully tagged docker_chapter-website:latest
    Starting postgres ... done
    Recreating Half-Dome ... done
    Recreating Chapter-Website ... done
    Attaching to postgres, Half-Dome, Chapter-Website
    ```

3. netstat

    ```bash
    $ netstat -lna | grep '4200\|4201\|35432'
    tcp46      0      0  *.4200                 *.*                    LISTEN
    tcp46      0      0  *.4201                 *.*                    LISTEN
    tcp46      0      0  *.35432                *.*                    LISTEN
    ```

#### Manually Checking the Postgres Database

Run `docker exec -it postgres psql -U docker acm` after you have started up any of the compose files to sift through the database.

#### Cleaning Up When Done Developing or Trying to Reset Your Environment

```bash
$ docker system prune --volumes
WARNING! This will remove:
  - all stopped containers
  - all networks not used by at least one container
  - all volumes not used by at least one container
  - all dangling images
  - all dangling build cache

Are you sure you want to continue? [y/N] y
...
Total reclaimed space: __GB
```
