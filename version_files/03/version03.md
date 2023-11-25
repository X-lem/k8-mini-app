# Version 03

## Routes

| route         | http method | Request                 | Response                                  | Other info                                  |
| ------------- | ----------- | ----------------------- | ----------------------------------------- | ------------------------------------------- |
| /             | GET         |                         | _simple hello response_                   |                                             |
| /pod          | GET         |                         | _name of the k8 pod_                      |                                             |
| /secrets      | GET         |                         | _the data in the secrets yaml file_       | Prints out all the env variables to the log |
| /create-table | POST        |                         | _success message or error will be logged_ | Creates the `users` DB table                |
| /user         | POST        | { "username": "X-lem" } | _returns the created user_                | Creates a user with the given username      |
| /users        | GET         |                         | _returns all the created users_           |                                             |

## Section Requirments.

If you're running this on Windows WSL2 you'll need to install [docker desktop](https://docs.docker.com/desktop/install/windows-install/). Without it you cannot run the `docker-compose` commands.

## Explanation

In this version we're only going to be running things locally via Docker. This is because we're adding the `./client`. The client is a simple react application that allows you to make the DB requests via a UI.

To do this take a look at the `./docker-compose.yaml` file. This file has 3 docker services.

1. server
   - The server you should already be familiar with. It functions exactly the same as it has previously. Instead of building the server/Dockerfile directly we allow the docer-compose file to do it. We've added a `depends_on` to ensure that the DB service is up and running prior to the server. Because we've added the condition `condition: service_healthy` the db service needs the `healthcheck` parameters.
2. client
   - The client is new. Like the `./server/Dockerfile` the `./client/Dockerfile` is responsible for creating the Docker image. It builds the react application than uses nginx set it up so we can access it on port 80. The docker-file then forwards our local 3000 port to the docker port.
3. db
   - The DB is fairly simple. It's pulling the postgres image from [dockerhub] (https://hub.docker.com/_/postgres/) and creating the `k8-mini-app` database. Notice the `healthcheck`. Every 2 seconds the **test** is run. If after 20 seconds or 5 failed attempt the health check will be concidered `unhealth`. Feel free to change these if you're using a slower computer that needs more time to load up the db service. See [healthcheck](https://docs.docker.com/engine/reference/builder/#healthcheck) documentation for more information.

### Starting the services

There are two main commands we'll be using here. `docker-compose up` and `docker-compose down`. In the `./k8-mini-app` folder run the `docker-compose up` command. The first time you run this it will take a little while to build and pull the images.

You will get something like this followed by a bunch of logs

```
[+] Building 0.0s (0/0)                                                                                                                                                                                                                    docker:default
[+] Running 4/4
 ✔ Network k8-mini-app_default     Created
 ✔ Container k8-mini-app-client-c  Created
 ✔ Container k8-mini-app-db-c      Created
 ✔ Container k8-mini-app-server-c  Created
...
...
k8-mini-app-server-c  | 2023/11/21 16:38:35 Hello, From K8 Mini App DB
k8-mini-app-server-c  | 2023/11/21 16:38:35 k8-mini-app started
...
```

The terminal will remain connected to the docker-compose process and will show logs for all the services. You can use `ctrl + c` to stop the containers. The containers will still exist, you can valid this via `docker container ls -a`. Therefore you're DB will retail all of it's data.

If you make changes to the `./client` or `./server` you'll need to run `docker-compose up --build`.

### Validating Postgres

To connect to the psql DB that's hosted locally you can run: `docker exec -it <container_name> psql -U postgres`. You can get the container name via `docker ps`.

So in our case it would be: `docker exec -it k8-mini-app-db-c psql -U postgres`.

Like in version 2 you should get an output like: `postgres=# `

Type `\list` or `\l` to see a list of all the tables. You should see the `k8-mini-app` table there. `\connect <db_name>` or`\c <db_name>` will connect you to a given database. Typing `\q` to exit postgres.

You can also connect directly to the DB via: `docker exec -it k8-mini-app-db-c psql -U postgres k8-mini-app`.

### Terminating Containers

To delete the containers you can run `docker-compose down`. You'll get an output like this.

```
[+] Running 4/4
 ✔ Container k8-mini-app-client-c  Removed
 ✔ Container k8-mini-app-server-c  Removed
 ✔ Container k8-mini-app-db-c      Removed
 ✔ Network k8-mini-app_default     Removed
```

**WARNING**: This will delete the database along with all its data.

## Suggestions
*Easy*
- Add to the client application a way for the user to update/delete a user (hint, you may need to udpate your cors headers)

*Medium*
- In the `docker-compse.yaml` add a service that create an image via mysql. Update the `main.go` file to connect to this (either instead of postgres or in addition too).