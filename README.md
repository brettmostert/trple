# trple - who know what it may be
trple

## Project Structure

Below is a brief description of the project folder structure. 

### /cmd

Contains the main executeable applications, namely the `cli` and `server`.

### /internal

Contains all the private code for this application to work.

### /scripts

Scripts to perform various build, install, analysis, etc. This allows us to keep our makefile in the root small.

## BootStrap

### https://golangci-lint.run/
`curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0`

for alpine
`wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.43.0`

# Giving Thanks
A Big Shout Out! To the people who worked on the following, the work you have done has aided in my learning of the go language and eco-system. 

- https://github.com/golang-standards/project-layout
- https://github.com/hashicorp/terraform
- https://github.com/kubernetes/kubernetes/

