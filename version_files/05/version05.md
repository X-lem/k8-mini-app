## Routes

| route         | http method | Request                 | Response                                  | Other info                                  |
| ------------- | ----------- | ----------------------- | ----------------------------------------- | ------------------------------------------- |
| /             | GET         |                         | _simple hello response_                   |                                             |
| /pod          | GET         |                         | _name of the k8 pod_                      |                                             |
| /secrets      | GET         |                         | _the data in the secrets yaml file_       | Prints out all the env variables to the log |
| /create-table | POST        |                         | _success message or error will be logged_ | Creates the `users` DB table                |
| /user         | POST        | { "username": "X-lem" } | _returns the created user_                | Creates a user with the given username      |
| /users        | GET         |                         | _returns all the created users_           |                                             |

## Section Requirements.

In the `client/Docker file` comment out the line `COPY nginx.conf /etc/nginx/nginx.conf`. Instead of using this we're going to use the `ingress.yaml` to route the traffic.


## Setup

Instead of using nginx in the client, we're going to create an ingres which will handle directing traffic to out application. In order to do this we will need an ingress controller. There are a couple of ways to set this up.

1. Using yaml files
2. Using helm

For simplicity we're going to use helm.

### What is helm?

Helm is a package manager for Kubernetes. It's like what `npm` is for node or `apt` is for linux. People can create and write yaml file configurations and share them on the [https://helm.sh](heml website). So instead of you having to write a bunch of yaml files to create some kubernetes objects (like we're doing in our postgres.yaml) we can use an already created set of configurations in a helm package (called charts). Check out their website for getting helm installed on your computer.

### Setting up ingress nginx

We're going to be using [this github](https://github.com/kubernetes/ingress-nginx) repo to accomplish this. It's important to take note of which version of Kubernetes you're on as not all versions are compatible with all ingress-nginx versions (see *Supported Versions table* on github link). You'll then want to take note of the **Ingress-NGINX version** that matches up with your **k8 supported version**. There will probably be several that match up so note the more recent version.

1. To determine your kubernetes version
    ```
    $ kubectl version --short
    Flag --short has been deprecated, and will be removed in the future. The --short output will become the default.
    Client Version: v1.25.0
    Kustomize Version: v4.5.7
    ```
2. We need to add (install) this chart:
    ```
    $ helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
    "ingress-nginx" has been added to your repositories
    ```
3. Then get an output of all the versions:
    ```
    $ helm search repo ingress-nginx --versions
    NAME                            CHART VERSION   APP VERSION     DESCRIPTION                                       
    ingress-nginx/ingress-nginx     4.8.4           1.9.4           Ingress controller for Kubernetes using NGINX a...
    ingress-nginx/ingress-nginx     4.8.3           1.9.4           Ingress controller for Kubernetes using NGINX a...
    ingress-nginx/ingress-nginx     4.8.2           1.9.3           Ingress controller for Kubernetes using NGINX a...
    ```
4. You'll want to find the **APP VERSION** that matches the version you noted before. You can then replace the `${...}` with the appropriate value. This will create the yaml files you need for your ingress controller.
    ```
    helm template ingress-nginx ingress-nginx \
    --repo https://kubernetes.github.io/ingress-nginx \
    --version ${CHART_VERSION} \
    --namespace ingress-nginx \
    > ./version_files/05/nginx-ingress.${APP_VERSION}.yaml
    ```
You should now see a new `./version_files/05/nginx-ingress.${APP_VERSION}.yaml`

5. Create a namespace for the ingress (this namespace is standard for ingress-nginx): `kubectl create namespace ingress-nginx`
6. As usual, to deploy everything, use: `kubectl apply -f version_files/05`. Reminder that the `ingress.yaml` file will take a few minutes to properly assign an IP address.


To get your ingress pods you can do the following:
`kubectl get po -n ingress-nginx`. Remember you're looking for the controller (should say Running)
To watch the logs of your pod you can do the following: `kubectl -n ingress-nginx logs -f <pod_name>`.
This will help you troubleshoot if needed.

Use the following commands
1. `kubectl get ingress`
2. `kubectl get svc -n ingress-nginx`

You'll notice that the **ADDRESS** and **EXTERNAL-IP** are the same. This is the IP you want to assign your DNS record for the IP you listed in host (in `ingress.yaml`). If you didn't list an domain name there you should be able to go to the IP directly.

## Troubleshooting

If you are getting this error in gCloud on your ingress `Missing one or more resources. If resource creation takes longer than expected, you might have an invalid configuration.` it's probably because you're nginx controller was set up correctly.