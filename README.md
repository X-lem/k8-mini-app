# K8 Mini App

This is a simple golang application that can be deployed to any site that supports Kubernetes or locally via minikube. It is meant to be a simple example and that can help you get your feet wet in the world of Docker and Kubernetes. Feel free to fork it. Making changes and tinkering with this simple example will help you to learn and understand Kubernetes.

You'll notice a `personal.md` is listed as untracked in the `.gitignore`. You can create/use this file to keep all your commands for easy copy/paste so you don't accidentally commit them.

### Resources

- Sign up for gCloud
  - https://cloud.google.com
- Get gCloud CLI working locally
  - https://cloud.google.com/sdk/docs/install-sdk
  - Some step by step guides
    - mac: https://github.com/stacksimplify/google-kubernetes-engine/blob/main/03-gcloud-cli-install-macos/README.md
    - windows: https://github.com/stacksimplify/google-kubernetes-engine/blob/main/04-gcloud-cli-install-windowsos/README.md
- Install Docker
  - https://docs.docker.com/get-docker/
- K8 yaml guides
  - https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/
  - https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#workloads-apis (has examples)

## gCloud

1. Create a gCloud account and set up a billing account. You should be given $300 worth of credits (to be used over 3 months).
2. In gCloud go to the Kubernetes Engine and enable it
3. Once it's enabled gollow this guide to create a standard k8 cluster:
   - https://github.com/stacksimplify/google-kubernetes-engine/tree/main/02-Create-GKE-Cluster#step-02-create-standard-gke-cluster
   - Note: you'll need to active a billing account in order to select **regional** as the location type. It might work keeping it as zonal, but I have not tested this.
   - After following this my estimated monthly cost for the cluster came to $87.01 USD
   - To access it locally use the following command: `gcloud container clusters get-credentials <CLUSTER-NAME> --region <REGION> --project <PROJECT-NAME>`
4. (opt). Set up a budget alert so your CC isn't accidentally charged
   - Under credits **unselect** the check boxes _Discounts_ & _Promotions and others_. Otherwise the budget alert won't actually happen until after your credits are completely used up.

## Server

The `Dockerfile` is what is used to generate the docker image of the golang code. This assumes you have docker installed and it's running. It also assume you're in the `/server` directory.

1. `docker build --tag k8-mini-app .` This will generate a docker file from the `main.go` file
2. `docker image ls` should show you a list of images. You should see k8-mini-app there.
3. `docker run -p 8080:8080 k8-mini-app` This command runs the Docker container locally and maps port 8080 from the container to port 8080 on your host machine (**[host_port]:[container_port]**). Now you can access the app via http://localhost:8080 (via your browser, postman, curl, etc).
4. `docker ps` will get you information about the container including the name and id
5. `docker logs --follow <container-id>` will follow the logs of the container
6. `docker stop <container-name>` to shut down the container

## Launching to gCloud

This section assumes you've got google cli working locally

1. Create a gCloud Artifact Registry
   - Click “Create Repository”
   - Pick a **name** and **region** (you’ll need these later). The rest can be default (Docker, Standard, Google-managed, Dry run).
2. You'll need to set up your local env to be able to push to this
   - `gcloud auth configure-docker <registry-region>-docker.pkg.dev`
   ```
   Adding credentials for: us-west2-docker.pkg.dev
   After update, the following will be written to your Docker config file located at [/home/<username>/.docker/config.json]:
   {
     "credHelpers": {
       "us-west2-docker.pkg.dev": "gcloud"
     }
   }
   ```
3. Tag your docker image
   - `docker tag <source-image> <region>-docker.pkg.dev/<project_id>/<repository_id>/<image-name>`
   - Eg: `docker tag k8-mini-app us-west2-docker.pkg.dev/my_project/my_repository/k8-mini-app`
   - In the future you can just use this image name when you build your docker image (step 1 of server)
   - **For simplicity in the future I'll refer to the full gCloud image name (_`<region>-docker.pkg.dev/<project_id>/<repository_id>/<image-name>`_) as `<gcloud-image_name>`**
4. Push to gCloud
   - `docker push <gcloud-image_name>`
   - Eg: `docker push us-west2-docker.pkg.dev/my_project/my_repository/k8-mini-app`
5. Verify it was created
   - `gcloud artifacts docker images list <region>-docker.pkg.dev/<project_id>/<repository_id>`
   - Eg: `gcloud artifacts docker images list us-west2-docker.pkg.dev/my_project/my_repository`
   - You can also view it online in the Artifact Registry

## Create Deployment

1. In `server/k8_files/01/server_deployment.yaml` you'll need to update the image with the one you just created/pushed to gCloud.
2. Deploy the yaml files in server/k8_files/01
   ```
   kubectl apply -f server/k8_files/01/server_deployment.yaml
   kubectl apply -f server/k8_files/01/server_loadbalancer.yaml
   kubectl apply -f server/k8_files/01/server_secrets.yaml
   ```
   - Optionally you can deploy the whole folder via: `kubectl apply -f server/k8_files/01`
3. To later delete the k8 objects in gCloud you can run the same commands but change `apply` -> `delete`

### Explanation

- The `server_deployment.yaml` is what will actually create + deploy the k8-mini-app in gCloud. It determines how many pods will be created, what image and port is to be used, and any non-sensitive environment variables you want. Feel free to add more.
- The `server_loadbalancer.yaml` is what makes the deployment publicly accessible.
- The `server_secrets.yaml` is where you would normally store sensitive data (db password, api keys, etc). Feel free to add more.

## Test the deployment

1. `kubectl get all` will get all the objects you've created. You'll see something like this

```
$kubectl get all
NAME                                      READY   STATUS    RESTARTS   AGE
pod/k8-mini-app-server-594f55947b-xlb2q   1/1     Running   0          6m20s
pod/k8-mini-app-server-594f55947b-psvbh   1/1     Running   0          6m20s

NAME                               TYPE           CLUSTER-IP     EXTERNAL-IP     PORT(S)          AGE
service/k8-mini-app-loadbalancer   LoadBalancer   10.108.8.200   35.233.156.20   8080:32719/TCP   24m
service/kubernetes                 ClusterIP      10.108.0.1     <none>          443/TCP          7d

NAME                                 READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/k8-mini-app-server   2/2     2            2           25m

NAME                                            DESIRED   CURRENT   READY   AGE
replicaset.apps/k8-mini-app-server-594f55947b   2         2         2       6m25s
```

2. The number of pods created will depend on how many you set in the `server/k8_files/01/server_deployment.yaml`
3. `service/kubernetes` is a default object that will exist pre/post creating/deleting everything
4. You’ll notice on the service (`service/k8-mini-app-loadbalancer`) that an EXTERNAL-IP was created (takes a few moments). You can now use this to access your application (don’t forget to add the port). You’ll notice if you go to the `/pod` route it will return the name one of the pods! Using `curl` is the best for this as your browser might cache the request and thus return the same response.

```
$ curl http://35.233.156.20:8080
Hello, From K8 Mini App ouo <3
$ curl http://35.233.156.20:8080/pod
k8-mini-app-server-594f55947b-xlb2q
$ curl http://35.233.156.20:8080/pod
k8-mini-app-server-594f55947b-psvbh
```

5. Going to `/secrets` will return the values you put in the `server_secrets.yaml`. It will also log all of the environment variables in the pod. To see these logs you can use the following:
   - `kubectl logs <pod-name>` or `kubectl logs -f <pod-name>` to watch the logs live

## Making Changes

Here is how you can update and add to the `main.go` file.

1. Make changes to the Go app
2. Build docker image
   - `docker build --tag <gcloud-image_name> .`
   - eg. `docker build --tag us-west2-docker.pkg.dev/my_project/my_repository/k8-mini-app .`
3. Push to gCloud
   - eg. `docker push <gcloud-image_name>`
   - eg. `docker push us-west2-docker.pkg.dev/my_project/my_repository/k8-mini-app`
4. Update deployment image
   - `kubectl set image deployment <deployment_name> <container_name>=<gcloud-image_name>:<tag>`
   - eg. `kubectl set image deployment k8-mini-app-server server=us-west2-docker.pkg.dev/my_project/my_repository/k8-mini-app:latest`
   - Using `latest` for the tag here is the important part. This will make sure it looks at the newest image.
   - The `<deployment_name>` and `<container_name>` are what we specified in the `server_deployment.yaml` file.
   - This doesn't seem to be 100% consistent (probably because the tag isn't changing). If it doesn't work you can restart the deployment: `kubectl rollout restart deployment k8-mini-app-server`
5. It will be a few seconds to update the pods

## Sections

You'll notice this repo has several folders in `/k8_files`. Each version folder builds upon each other. The guide above is for getting version `01` up and ready. Each of the version folders relies on the same `main.go` and `Dockerfile` so deploying them will be simple and can be done so by just deploying the desired yaml files. The version folders hold the yaml files that will be needed to deploy that versions functionality. In each of the versions folder there will be a `routesXX.md` file that will show the available routes that version allows. It is likely best to delete all the k8 objects created from one version before deploying the yaml files from another. This will ensure a totally clean deployment.
