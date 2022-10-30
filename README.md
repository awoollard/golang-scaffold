# Golang Scaffold

The purpose of this repository is to provide a sample Golang application which utilises standard technologies and patterns such that developers may re-use this repository when spinning up their own microservices.

## Dependencies
* Golang 1.9
* [github.com/gorilla/mux](https://github.com/gorilla/mux)
* [github.com/stretchr/testify](https://github.com/stretchr/testify)

## Endpoints
* `/` - Simple "Hello World" response
* `/status` - Returns a JSON structure containing application description, version and SHA (current Git revision).

## Usage
Compile build locally:

`go build`

Run unit tests:

`go test`

Build using Docker:

`docker build -t golang-scaffold .`

Run Docker container:

`docker run --publish 80:80 golang-scaffold`

Connect into a shell of the running container:

`docker exec -it [container-id] /bin/sh`

## CI/CD
Upon pushing a git commit to this repository, a Github workflow will execute using Github Actions.

The application will be built, tested, and published to GHCR.

You may see past runs at the [Github Actions](https://github.com/awoollard/golang-scaffold/actions) page.
## Project Structure
```
├─── .github
│   └─── workflows
│      └─── go.yml
├─── model
│   └─── status.go
├─── Dockerfile
├─── go.mod
├─── handlers.go
├─── handlers_test.go
├─── meta.json
├─── utils.go
└─── utils_test.go
```

### .github/workflows
Contains YAML files specifying Github workflows.
Workflows contain jobs which are comprised of steps defining actions which get executed sequentially within a virtual environment.

The workflow defined in `go.yml` performs the following actions:
1. Checks out this git repository
2. Installs Golang version 1.19
3. Compiles the application from source code
4. Runs unit tests
5. Logs into GHCR (Github Packages Repository)
6. Extract metadata (tags, labels) for Docker
7. Builds the Docker image
8. Publishes the Docker image to GHCR at [ghcr.io/awoollard/golang-scaffold](https://ghcr.io/awoollard/golang-scaffold)

### Dockerfile
The base image for the dockerfile is `golang:1.19`, which contains the `git` binary.
This is important because after copying the source code to the target Docker container file-system,
and moving the `meta.json` to the root of the filesystem, we utilise a git command, `git rev-parse HEAD > /git.sha`, which prints the current git revision of this repository to a file located also in the root filesystem of the Docker container.

This allows the API runtime to simply parse the `/meta.json` and `/git.sha` files when writing a response for the HTTP `GET` `/status` endpoint.

### model/
Directory containing domain models for the application. There is currently only one file within this directory containing a few structs used by the endpoint handlers.

### go.mod
Contains Golang dependencies used by this project.

### handlers.go
Contains two handlers routed via HTTP using Mux.

### handlers_test.go
Unit tests for the handlers.

### meta.json
Simple JSON structure containing values for the application description and version.

### utils.go
Utility functions which take a file-handle, reads the contents and parses the data.

### utils_test.go
Unit tests for the utility functions.

## Improvements, Limitations and Risks
There are a variety of ways this repository can be improved. Here's a few:
* Any user or application that has write-access to the SHA file may overwrite the contents of it, which would have the consequence of the `/status` endpoint displaying whatever contents it's overwritten with within its JSON response payload.
* The Github workflow builds the application twice. First in the Github Action steps, and then within Docker. This is intentional for now because performing a quick Golang build of the application takes ~2 seconds, and allows us to instantly test the code without first spinning up a 1 GB Docker image, bootstrapping the environment, etc. It allows for developers to near-instantly receive feedback in case their code breaks tests.
* [Docker BuildX](https://docs.docker.com/engine/reference/commandline/buildx/) custom build hooks were investigated for their usage in injecting the SHA hash to the filesystem but this proved heavy-handed and in the end a simple `RUN` command in the `Dockerfile` proved sufficient.
* More unit tests - could abstract out the file structs using interfaces such that the `/status` endpoint handler could be properly tested.
* The artefact published to GHCR could have more tags. For example, it could be possible to tag artefacts with a `pull-request` tag or a `release` tag depending on which Git branch the code was published from.

## Output
![Sample output](https://i.imgur.com/Snnuril.png "Output")