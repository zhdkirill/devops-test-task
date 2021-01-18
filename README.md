# Ataccama DevOps Engineer Technical Challenge

Hello! If you're reading this, you've applied to work with us as a DevOps Engineer. That's great!
As a first round of our hiring process, we want to see your technical ability. Please, read this README carefully, as it contains everything that you neeed to complete the task.

## Goal

This repository contains a little Golang application in `app` folder. You need to set up a complete end-to-end CI/CD pipeline, that will build & deploy an application on every commit to `master` branch.

You can read more about the application and it's dependencies in `app/README.md` file.

Please, clone this repo to your public GitHub, GitLab, BitBucket or other git service and go on from there. After you're done, please send us link to your repository.

You can use any open-source tools to implement this CI/CD pipeline, as well as any way to package and run the application. If you'd like to have a VM or two to work on - shoot us an email, we'll deploy two instances in AWS for you.

## Evaluation criteria

The main evaluation criteria is that the CI/CD pipeline works. Application has to run somewhere and be accessible via web browser. When a change in `master` branch of repository occurs, it should automatically propagate to application.

Apart from that we will evaluate:
* tools that you've picked
* correct usage of said tools
* code readability
* solution flexibility & extensibility

## To sum up

This task is loosely-defined on purpose. With DevOps being a set of practices at best and a buzzword at worst, we want to leave you some creative freedom in how you approach this task. There is no wrong solution, apart from a solution that outright doesn't work, so don't worry about it.

Good luck, and hope to hear back from you soon!

_Ataccama Cloud Solutions team_

# Solution description
My solution utilizes only GitHub and Azure services and does not require any other services, servers or local commands to run.

## Prerequisites
The pipeline requires a few pre-confugured Azure resources:
- resource group to group the services
- container registry to store application images
- kubernetes service to run the application
- service principal for authentication
These resources need to be created only once.

## Architecture
The pipeline is executed by GitHub Actions on each `push` to the repository.

### Step 1: Testing
Golang code is being verified by a linter: https://github.com/golangci/golangci-lint
If the code fails to pass the test, the application would not be built and published.
There is also an additional action with the same linter that is triggered on pull requests and does not build a package.

### Step 2: Packaging
The application will be distributed as a docker image.
Fist, the application is being built via golang container. After that the resulting binary is being copied to another alpine-based container without go runtime. Such strategy reduces the size of the resulting image.

### Step 3: Publishing
Container image of the application is published in an Azure container registry repository. The images are tagged with the commit hash.

### Step 4: Deploying
The application is being deployed as a microservice to Azure Kubernetes Service.
The application requires redis database, so there are two deployments and services: ervcp and redis. Redis is exposed as a service inside the cluster, while ERVCP application service is exposed via ingress loadbalancer to the internet. The application is available over FQDN provided by Azure: http://ervcp.germanywestcentral.cloudapp.azure.com/
