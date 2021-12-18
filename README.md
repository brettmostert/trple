# trple - who know what it may be
trple

## cmd = for the monolith using all the stuff in the microservices
## product for the microservice
## product / cmd for its microservice etc

## Project Structure

Below is a brief description of the project folder structure. 

### /cmd

Contains the main executeable applications, namely the `cli` and `server`.

### /internal

Contains all the private code for this application to work.

### /scripts

Scripts to perform various build, install, analysis, etc. This allows us to keep our makefile in the root small.

## Proto Workflow

- Mono Repo
- Project for proto
- generate project for go modules to be shared (look how to do this)
    this is so that we can share the code between projects
    if we were using multiple programming languages i would put proto in its own repo
    and generate libs from there into their own repos
- the folder under proto ie product is per component
- .protolangs is for the languages we are generating for... if its not go it goes into its own repo (weird i know but i wanted to play with as if we had multiple repos in java or go or nodejs etc)

https://www.bugsnag.com/blog/libraries-for-grpc-services - mention of go modules aka dep
https://medium.com/namely-labs/how-we-build-grpc-services-at-namely-52a3ae9e7c35
https://github.com/namely/chief-of-state
https://github.com/namely/chief-of-state-protos
https://github.com/namely/docker-protoc
https://golangbyexample.com/package-vs-module-golang/

## BootStrap

### https://golangci-lint.run/
`curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0`

for alpine
`wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.43.0`

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
`sudo apt-get install protobuf-compiler`
# Giving Thanks
A Big Shout Out! To the people who worked on the following, the work you have done has aided in my learning of the go language and eco-system. 

- https://github.com/golang-standards/project-layout
- https://github.com/hashicorp/terraform
- https://github.com/kubernetes/kubernetes/

install buf (linting and breaking changes)
go generators
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
install protoc


PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip
7z x -y -o/home/$USER/.local/protoc ./protoc-3.19.1-linux-x86_64.zip
