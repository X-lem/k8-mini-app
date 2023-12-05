# K8 Mini App

This is a simple golang/react application that can be deployed to any site that supports Kubernetes or locally via minikube. It is meant to be a simple example and that can help you get your feet wet in the world of Docker and Kubernetes. Feel free to fork it. Making changes and tinkering with this simple example will help you to learn and understand Docker/Kubernetes.

You'll notice a `personal.md` is listed as untracked in the `.gitignore`. You can create/use this file to keep all your commands for easy copy/paste so you don't accidentally commit them.

## Versions

You'll notice this repo has several folders in `./verion_files`. Each version folder builds upon each other. Start with version `01` to get a simple Docker/K8 application up and ready. Each of the version folders relies on the same `main.go` and `Dockerfile` so deploying them will be simple and can be done so by just deploying the desired yaml files. The version folders hold the yaml files that will be needed to deploy that versions functionality (if any).

In each of the versions folder there will be a `versionXX.md` file that will show the available routes that version allows as well as any instructions on the specific version. It is recommended to delete all the k8 objects created from one version before deploying the yaml files from another. This will ensure a totally clean deployment.

Some of the versions will have suggestions at the end on things you can add to help you learn.

### Version Overviews

1. Locally build a golang application (server) and run it via docker/kubernetes with minikube.
2. Launch the golang application to gCloud kuberbetes
3. Add a DB
4. Adding a front end (client). Running locally via docker compose and launching to gCloud
5. Using ingress to handle traffic (WIP)

## Resources

- K8 yaml guides
  - https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/
  - https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#workloads-apis (has examples)
