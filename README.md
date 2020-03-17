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

### Manually Checking the Postgres Database

Run `docker exec -it postgres psql -U docker acm` after you have started up any of the compose files to sift through the database.
