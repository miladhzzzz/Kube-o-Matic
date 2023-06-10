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

## Usage
**To Use Kube-o-Matic please follow the instructions bellow:**

```shell
# Clone The repository
git clone https://github.com/miladhzzzz/Kube-o-Matic.git

# Change Directory into it
cd Kube-o-Matic

# Use Makefile to build and run a docker image
make docker

# after we are finished with building and running the image we need to upload kubeconfig
cd hack && chmod +x ./upload-kubeconfig.sh && ./upload-kubeconfig.sh -c <example> -a <http://localhost:8555>
```
Theres a upload-kubeconfig.sh file in /hack directory which exports your kubernetes configuration based on the context you provided otherwise it will use default and uploads the config file to Kube-o-Matic.

**Please note that you need to change -c persys to provide your own context and api address**