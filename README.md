# ACM

---

ACM's Website Mono-Repository. Inside contains submodules for all micro-frontend and micro-backend apis

## Prequisities

1. Docker version

   ```bash
   $ docker -v
   Docker version 19.03.8, build afacb8b
   ```

2. Docker Compose version

   ```bash
   $ docker-compose -v
   docker-compose version 1.25.4, build 8d51620a
   ```

## Getting Started

1. `git submodule init`
2. `git submodule update`
3. The whole project runs on multiple docker containers on a shared network. You can find all off the development yaml files in the `/docker` folder

## Run the Project with Docker Compose

Below is a list of possible combinations to run the ACM Website locally

### Website with all API Services

1. Run the root yaml compose file `docker-compose up -d`
2. To create all tables to store data on run the bash script with this command `chmod +x database_setup.sh && ./database_setup.sh`

> Note: When running the project this way all applications opens on `http://127.0.0.1:<portNumber>`

To check if they are setup correctly:

1. Docker Compose

   ```bash
   $ docker-compose ps
        Name                    Command              State            Ports
   ---------------------------------------------------------------------------------
   Chapter-Website   docker-entrypoint.sh yarn  ...   Up      0.0.0.0:4200->4200/tcp
   Half-Dome         docker-entrypoint.sh yarn  ...   Up      0.0.0.0:4201->4201/tcp
   Tuolumne          docker-php-entrypoint php  ...   Up      0.0.0.0:4202->4202/tcp
   pgadmin           /entrypoint.sh                   Up      443/tcp, 0.0.0.0:8080->80/tcp
   postgres          docker-entrypoint.sh postgres    Up      0.0.0.0:35432->5432/tcp
   ```

2. Docker

   ```bash
   $ docker ps
   CONTAINER ID        IMAGE                  COMMAND                  CREATED             STATUS              PORTS                           NAMES
   bcdd8cab6353        chapter-website:v0.1   "docker-entrypoint.s…"   3 minutes ago       Up 3 minutes        0.0.0.0:4200->4200/tcp          Chapter-Website
   29ace3f0ea26        tuolumne:v1.0          "docker-php-entrypoi…"   21 minutes ago      Up 21 minutes       0.0.0.0:4202->4202/tcp          Tuolumne
   36a54a443c24        half-dome:v1.1         "docker-entrypoint.s…"   21 minutes ago      Up 21 minutes       0.0.0.0:4201->4201/tcp          Half-Dome
   2b7e5a24c88a        dpage/pgadmin4         "/entrypoint.sh"         21 minutes ago      Up 21 minutes       443/tcp, 0.0.0.0:8080->80/tcp   pgadmin
   2d3f4c37df54        postgres:12.2-alpine   "docker-entrypoint.s…"   21 minutes ago      Up 21 minutes       0.0.0.0:35432->5432/tcp         postgres
   ```

### Debugging Issues

#### Containers not Updating with New Code

1. Recreate Containers

   ```bash
   $ docker-compose up -d --force-recreate
   Recreating pgadmin         ... done
   Recreating postgres        ... done
   Recreating Half-Dome       ... done
   Recreating Tuolumne        ... done
   Recreating Chapter-Website ... done
   Attaching to postgres, Half-Dome, Chapter-Website
   ...
   ```

2. Rebuild Containers

   ```bash
   $ docker-compose up -d --build
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

   > Rebuilding a specific container run `docker-compose up -d --force-recreate --build --no-deps <service-name>`

3. netstat

   ```bash
   $ netstat -lna | grep '4200\|4201\|35432'
   tcp46      0      0  *.4202                 *.*                    LISTEN
   tcp46      0      0  *.8080                 *.*                    LISTEN
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
