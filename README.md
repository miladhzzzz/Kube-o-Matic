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

#### GitHub Deployer Robot!
- theres a github robot you can use to check your manifests / automate other tasks right from the api no coding neded!

#### kubectl in a browser!
- so you can use the swagger UI to easily interact with the your kubernetes cluster

#### dynamic kubeconfig loading!
- you can import and inject your kube config in a bunch of ways that we'll discuss in depth. keep on reading!

#### monitoring!
- you can easily set up jobs and watches and see the results in JSON!(Come on now i am not a Frontend Developer!)

## Installation

**To install Kube-o-Matic please follow the instructions bellow:**

#### Build
using make you can build and run the project binary on linux.
```shell
# build go and output to /bin
make build
# run the binary
make run
```

#### Run in Docker

This script will build a docker image then run it and inject your kubeconfig with /hack/upload-kubeconfig.sh.

```shell
# Clone The repository
git clone https://github.com/miladhzzzz/Kube-o-Matic.git

# Change Directory into it
cd Kube-o-Matic

# Use Makefile to build and run a docker image
make docker

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

    2. webhook secret is basically preventing anyone else from activating your CD Pipeline! so you have to provide a good one and then use http://localhost:8555/webhook/secret/<YOUR SECRET HERE> to set the secret in your CD.

    3. make sure all of your manifests are in the root directory.

    4. Push Changes To your Repo and they will get deployed eachtime you update the code!

### Kubectl API

- you can use Postman , CURL, Browser REST Clients , or any other REST Client to communicate and monitor your kubernetes cluster.

### Endpoints

| Endpoint | HTTP Method | Description |
| --- | --- | --- |
| / | GET | Returns a success message and a value |
| /kube | GET | Returns a list of all Kubernetes clusters |
| /kube/:cluster_id | GET | Returns the details of a specific Kubernetes cluster |
| /kube/:cluster_id | DELETE | Deletes a specific Kubernetes cluster |
| /kube/:cluster_id/node | GET | Returns a list of all nodes in a specific Kubernetes cluster |
| /kube/:cluster_id/node/:node_id | GET | Returns the details of a specific node in a specific Kubernetes cluster |
| /kube/:cluster_id/pod | GET | Returns a list of all pods in a specific Kubernetes cluster |
| /kube/:cluster_id/pod/:pod_id | GET | Returns the details of a specific pod in a specific Kubernetes cluster |
| /kube/:cluster_id/service | GET | Returns a list of all services in a specific Kubernetes cluster |
| /kube/:cluster_id/service/:service_id | GET | Returns the details of a specific service in a specific Kubernetes cluster |
| /kube/:cluster_id/namespace | GET | Returns a list of all namespaces in a specific Kubernetes cluster |
| /kube/:cluster_id/namespace/:namespace_id | GET | Returns the details of a specific namespace in a specific Kubernetes cluster |
| /kube/:cluster_id/deployment | GET | Returns a list of all deployments in a specific Kubernetes cluster |
| /kube/:cluster_id/deployment/:deployment_id | GET | Returns the details of a specific deployment in a specific Kubernetes cluster |
| /kube/:cluster_id/replicaset | GET | Returns a list of all replicasets in a specific Kubernetes cluster |
| /kube/:cluster_id/replicaset/:replicaset_id | GET | Returns the details of a specific replicaset in a specific Kubernetes cluster |
| /kube/:cluster_id/statefulset | GET | Returns a list of all statefulsets in a specific Kubernetes cluster |
| /kube/:cluster_id/statefulset/:statefulset_id | GET | Returns the details of a specific statefulset in a specific Kubernetes cluster |
| /kube/:cluster_id/daemonset | GET | Returns a list of all daemonsets in a specific Kubernetes cluster |
| /kube/:cluster_id/daemonset/:daemonset_id | GET | Returns the details of a specific daemonset in a specific Kubernetes cluster |
| /kube/:cluster_id/job | GET | Returns a list of all jobs in a specific Kubernetes cluster |
| /kube/:cluster_id/job/:job_id | GET | Returns the details of a specific job in a specific Kubernetes cluster |
| /kube/:cluster_id/cronjob | GET | Returns a list of all cronjobs in a specific Kubernetes cluster |
| /kube/:cluster_id/cronjob/:cronjob_id | GET | Returns the details of a specific cronjob in a specific Kubernetes cluster |
| /kube/:cluster_id/configmap | GET | Returns a list of all configmaps in a specific Kubernetes cluster |
| /kube/:cluster_id/configmap/:configmap_id | GET | Returns the details of a specific configmap in a specific Kubernetes cluster |
| /kube/:cluster_id/secret | GET | Returns a list of all secrets in a specific Kubernetes cluster |
| /kube/:cluster_id/secret/:secret_id | GET | Returns the details of a specific secret in a specific Kubernetes cluster |
| /kube/:cluster_id/ingress | GET | Returns a list of all ingress in a specific Kubernetes cluster |
| /kube/:cluster_id/ingress/:ingress_id | GET | Returns the details of a specific ingress in a specific Kubernetes cluster |
| /kube/:cluster_id/networkpolicy | GET | Returns a list of all networkpolicies in a specific Kubernetes cluster |
| /kube/:cluster_id/networkpolicy/:networkpolicy_id | GET | Returns the details of a specific networkpolicy in a specific Kubernetes cluster |
| /kube/:cluster_id/podsecuritypolicy | GET | Returns a list of all podsecuritypolicies in a specific Kubernetes cluster |
| /kube/:cluster_id/podsecuritypolicy/:podsecuritypolicy_id | GET | Returns the details of a specific podsecuritypolicy in a specific Kubernetes cluster |
| /kube/:cluster_id/storageclass | GET | Returns a list of all storageclasses in a specific Kubernetes cluster |
| /kube/:cluster_id/storageclass/:storageclass_id | GET