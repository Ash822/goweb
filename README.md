# HTTP Server
## Getting started

Prerequisites: `git` `go 1.16.3` `docker` `docker hub account`

### Running using Makefile
Commands required for testing, building the project and for serving swagger client as added to the `Makefile`.

Commands | Tasks| 
--- | --- | 
`make test` | To execute all the test in the project. | 
`make run` | To run the server. The server will be listening to port `8000`. | 
`swagger-gen` | To build swagger spec yaml and serve it. | 

### Running using Docker
Alternatively, the application can also be run using `docker`.

Use the following command to pull the docker image.

```aidl
docker image pull ash822/goweb
```

Run the application using the following command.

```aidl
docker run -p 8000:8000 ash822/goweb:latest
```

## About
This repository showcases a simple HTTP server serving REST endpoints, built using the Golang `net/http` package and `gorilla/mux` as http router.

### Endpoints
It supports the following endpoints.

![endpoints](.README_images/endpoints.png)

### Architecture

The data flow is designed to better manage the dependencies as follows.

![arch](.README_images/arch.png)

The router links the path to the handler functions defined in the controller. The service controls the business logic and persist the data using the repository.

### Testing

`Gomega` is used as the matcher library for assertions. In order to test the service, `GoMock` is used for mocking the repository.

The repository interface is decorated with `mockgen` annotations. The mock repository code can be generated using `make mock-gen` command.

The REST endpoints are tested using `net/http/httptest` package.

To execute the tests, run `make test` command.

#### API documentation

To generate the swagger file and to serve the swagger client, run `make swagger-run` command. 

[swagger.yaml](swagger.yaml)

### CI/CD

The project is setup with Continuous Integration and Continuous Deployment (to Docker hub) using Github Actions.

On every pull request to `master` branch and on every push to `master` branch, the workflow gets executed. The workflow checks out the latest, build and run tests. 

As the last step, the application builds and tags a docker image, and pushed the image to Docker hub registry.

Visit [Actions](https://github.com/Ash822/goweb/actions/workflows/build-deploy-goweb.yaml) tab to check the jobs.

