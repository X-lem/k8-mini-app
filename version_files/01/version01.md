# Version 01

Creating a simple kubernetes application and deploying it locally with dock and then minikube.

## Routes

| route    | http method | Request | Response                            | Other info                                  |
| -------- | ----------- | ------- | ----------------------------------- | ------------------------------------------- |
| /        | GET         |         | _simple hello response_             |                                             |
| /pod     | GET         |         | _name of the k8 pod_                |                                             |
| /secrets | GET         |         | _the data in the secrets yaml file_ | Prints out all the env variables to the log |

## Section Requirments.

- You'll need to install [docker desktop](https://docs.docker.com/get-docker/) - don't forget to enable kubernetes.
- You'll also need to install [minikube](https://minikube.sigs.k8s.io/docs/start/) to run k8 locally

If the installation was sucessful you should see `minikube` show up under **Containers** in Docker Desktop. Note, you'll need to run `minikube start`

## Resources

- K8 yaml guides
  - https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/
  - https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#workloads-apis (has examples)

## Running via Docker

The `Dockerfile` is what is used to generate the docker image of the golang code. Take a look through the code/Dockerfile The following instructions

1. `docker build --tag k8-mini-app /server` (or `docker build --tag k8-mini-app .` if you're in the `/server` directory already) This will generate a docker image from the `main.go` file. We're naming the image `k8-mini-app`.
2. `docker image ls` should show you a list of images. You should see k8-mini-app there.
3. `docker run -p 8080:8080 k8-mini-app` This command runs the Docker container locally and maps port 8080 from the container to port 8080 on your host machine (**[host_port]:[container_port]**). Now you can access the app via http://localhost:8080 (via your browser, postman, curl, etc).
4. `docker ps` will get you information about the container including the name and id
5. `docker logs --follow <container-id>` will follow the logs of the container
6. `docker stop <container-name>` to shut down the container.

## Running via Minikube

YAML files are used to configure Kubernetes. Take a look at all of the yaml files in this directory; there should be three. Above in the references are links to the k8 documentation that will help you to understand how they work. All k8 commands start with `kubectl`.

1. `kubectl config get-contexts` will display all the different k8 contexts that exist on your device. It should show the current context.
2. `kubectl config use-context <context-name>` to switch your context if nessisary (you'll be looking for `docker-desktop` or `minikube`).
3. There are two ways to deploy the yaml files to minikube.
   - Individually via `kubectl apply -f version_files/01/<file-name>.yaml`
   - The entire folder via `kubectl apply -f version_files/01`
   - Note, you can run the same command after updating the yaml files. K8 will update any changes it finds.
4. After deploying the files you can run `kubectl get pods` or `kubectl get po` to get a list of all the pods that were created. There should be 5 if you didn't change the `server_deployment.yaml` file.
5. `kubectl get all` will show all the k8 objects that exist.
6. Run `kubectl get service` (or `kubectl get svc`). You should see two object. `kubernetes` is one that exists apart from anything we did. You should also see a `k8-mini-app-loadbalancer` service. This is the one that was created from the `server_loadbalancer.yaml` file. This is what allows us to connect to our application (without it we couldn't). The load balancer is what directs traffic to the pods (ie. it splits the requests between the pods). Take a look at the _EXTERNAL-IP_ column. This is the IP address that we can use to access our application (in this case it should be **localhost**).
7. Access the application using curl. When you access the `/pod` route it'll return the pod that the load balancer directed your request too. These should match up with the pods displayed when you ran `kubectl get po`. Using curl is best to demonstrate this as the browser/postman can be continually directed to the same pod (probably of caching).
8. Finally to delete the k8 object you created you can run one of the following:
  - `kubectl delete -f version_files/01/<file-name>.yaml`
  - `kubectl delete -f version_files/01`

   ```
   $ curl http://35.233.156.20
   Hello, From K8 Mini App ouo <3
   $ curl http://35.233.156.20/pod
   k8-mini-app-server-594f55947b-xlb2q
   $ curl http://35.233.156.20/pod
   k8-mini-app-server-594f55947b-psvbh
   ...
   $ curl localhost/secrets
   {"nestedSecret":"MyNestedSecret","secret":"MySuperDuperSecret"}
   ```
8. To see the logs of the pods you can use the following command `kubectl logs <pod-name>` or `kubectl logs -f <pod-name>` to watch the logs live

### Troubleshooting

If your running Windows and WSL2 and you're getting this error while trying to build a docker image:

```
ERROR: failed to solve: golang:latest: error getting credentials - err: fork/exec /usr/bin/docker-credential-desktop.exe: exec format error, out: ``
```

Delete the line with `credsStore` from `~/.docker/config.json`. See this [stack overflow](https://stackoverflow.com/questions/65896681/exec-docker-credential-desktop-exe-executable-file-not-found-in-path) thread.

## Suggestions

_Easy_

- Update the `server_deployment.yaml` to create a single pod. Then follow the logs of that pod. Finally go to `/secrets` and you should see a list of all the enviroment variables that exist. 

_Medium_

- Update the deployment to use a specific port (ie. `localhost:8080`)