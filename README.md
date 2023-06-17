# kube-o-Matic!

"a robot butler serving kubernets on a platter!"

Open Source solution to automate use of kubectl!

> this project is under active development!

## What?
as everyone knows by now! in order to communicate with a kubernetes cluster you need to use a tool called "Kubectl".theres a bunch of commands that you can use to interact with your cluster!

**Kube-o-Matic!** provides a layer on top of k8s.io/client-go to automate "kubectl"!
we expose an HTTP REST API to bring kubectl to your browser!


## Why?
we use our browsers to interact with alot of different systems daily so why not kubernetes! ?

### Back Story!
as i was doing a system analysis on kubernetes to understand how it works and i was knee deep in Kubernetes repo :) i had Epiphany!!

kubectl is basically a wrapper for k8s.io/client-go. so client-go is the actuall client for kubernetes so then i used that knowledge to write my own wrapper that takes away little bit of weight from the back of developers / devops / sre teams!

### Features!

#### GitOps / CD
- as the first feature that came to mind was parsing webhooks from github to watch a set of manifests and if changed deploy them!

#### Multiple Cluster Management / Automation !
- you can have multiple clusters setup and add automation to deploy accross all of them!

#### GitHub Deployer Robot!
- theres a github robot you can use to check your manifests / automate other tasks right from the api no coding neded!

#### kubectl in a browser!
- so you can use the swagger UI / Any REST Client to easily interact with your kubernetes cluster.

#### dynamic kubeconfig loading!
- you can import and inject your kube config in a bunch of ways that we'll discuss in depth. keep on reading!

#### monitoring!
- you can easily set up jobs and watches and see the results in JSON!(Come on now i am not a Frontend Developer!)

## Installation

**To install Kube-o-Matic please follow the instructions bellow:**

#### Build
using make you can build and run the project binaries.

```shell
# build go and output to /bin
make build
# run the binary
make run
```

#### Run in Docker

- Download from **Github packages** and run in docker.

```shell
# If you have the repository cloned:
make PullAndRun

# IF you dont have the repository localy run :
docker pull docker pull ghcr.io/miladhzzzz/kube-o-matic
docker run -p 8555:8555 -d --name kube-o-matic ghcr.io/miladhzzzz/kube-o-matic

# DONT FORGET TO UPLOAD YOUR KUBECONFIG!

```

- This script will build a docker image then run it and inject your kubeconfig with /hack/upload-kubeconfig.sh.

```shell
# Clone The repository
git clone https://github.com/miladhzzzz/Kube-o-Matic.git

# Change Directory into it
cd Kube-o-Matic

# Use Makefile to build and run a docker image
make BuildAndRun

## INJECTING CONFIG

# after we are finished with building and running the image we need to upload kubeconfig
cd hack && chmod +x ./upload-kubeconfig.sh && ./upload-kubeconfig.sh -c <example> -a <http://localhost:8555>

# Example ./upload-kubeconfig.sh -c kind -a http://localhost8555

```

#### Manual Config injection

- Theres a upload-kubeconfig.sh file in /hack directory which exports your kubernetes configuration based on the context you provided otherwise it will use default and uploads the config file to Kube-o-Matic.

- if your are on windows try this:
```shell
curl -X POST -H "Content-Type: multipart/form-data" -F "file=@/path/to/kubeconfig" http://localhost/upload
```

## Usage

### Setting up GitOps for kubeomatic
- To setup our GitOps CD pipeline you need to head to your github repository : Settings > Webhook > Add Webhook and provide the following parameters:
    1. in order for github to be able to reach our CD we need a "Public IP Adrress". if you are running this localy use (Ngrok)[https://github.com/ngrok] and make sure you have it installed. provide the Ngrok URL in the webhook settings.

    2. webhook secret is basically preventing anyone else from activating your CD Pipeline! so you have to provide a good one and then use http://localhost:8555/webhook/secret/"<YOUR SECRET HERE>" to set the secret in your CD.

    3. make sure all of your manifests are in the root directory.

    4. Push Changes To your Repo and they will get deployed eachtime you update the code!

### Kubectl API

- you can use Postman , CURL, Browser REST Clients , or any other REST Client to communicate and monitor your kubernetes cluster.

### Endpoints

| Endpoint | Method | Description | Example |
| --- | --- | --- | --- |
| /kube/:cluster/deployments | GET | Returns a list of all deployments in a specific Kubernetes cluster | `curl http://localhost:8555/kube/my-cluster/deployments` |
| /kube/:cluster/replicasets | GET | Returns a list of all replicasets in a specific Kubernetes cluster | `curl http://localhost:8555/kube/my-cluster/replicasets` |
| /kube/:cluster/nodes | GET | Returns a list of all nodes in a specific Kubernetes cluster | `curl http://localhost:8555/kube/my-cluster/nodes` |
| /kube/:cluster/services | GET | Returns a list of all services in a specific Kubernetes cluster | `curl http://localhost:8555/kube/my-cluster/services` |
| /kube/:cluster/events | GET | Returns a list of all events in a specific Kubernetes cluster | `curl http://localhost:8555/kube/my-cluster/events` |
| /hooks | POST | Receives a webhook and processes it | `curl -X POST -H "Content-Type: application/json" -d '{"cluster": "my-cluster", "pod": "my-pod", "namespace": "my-namespace"}' http://localhost:8555/hooks` |