# Deploying an app to k8s cluster with minikube

Assuming you have knowledge about docker, golang, and what is kubernetes.

## Stages

We assume you already have docker and golang.

### Install Kubectl

```bash
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
```

### Install And Start Minikube

```bash
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
  && chmod +x minikube && sudo mv minikube /usr/local/bin

minikube start
```

### Create Docker Image

```bash
docker build -t [DOCKER_HUB_USERNAME]/k8s-go-pg-example .
```

### Push Image To Docker Hub

```bash
docker push [DOCKER_HUB_USERNAME]/k8s-go-pg-example
```

### Create Secret File On Cluster

```bash
kubectl apply -f deployments/go-pg-secret.yaml
```

### Create Deployment And Service On Cluster For PG

```bash
kubectl apply -f pg-deployments.yaml
kubectl apply -f pg-service.yaml
```

### Create Deployment And Service On Cluster For API
```bash
kubectl apply -f api-deployment.yaml
kubectl apply -f api-service.yaml
```

### Get Service Ip From Minikube
```bash
minikube service api-service --url
```

## Usage

Open the url that minikube gives you. And you will see our application runs on k8s cluster.


Api have 1 endpoint with 2 http method.
1. /users [GET] -> Retrieves all users but only 100
2. /users [POST] -> Creates a new user. Create a post request with body to create a user.

**body example**
```json
{
    "nickname": "serkanerip",
    "email": "serkanerip@gmail.com",
    "password": "123456"
}
```

I hope this tuto help you to start to learn kubernetes. This tuto only covers k8s basics.
