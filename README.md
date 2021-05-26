# The App

This web app consists of three API endpoints:

1. `GET /` returns "hello world"
2. `GET /health` returns the following status code:
   - 204 -> The app is functioning as expected
   - 503 -> The app is under maintenance
3. `GET /metadata` returns a json blob containing some properties of this app


# Development
The APIs are developed with a swagger-first approach, in which you describe API endpoints in OpenAPI 3 specification, i.e. Swagger, in `./openapi.yaml`

The codebase is in Golang, hence we use an OAS 3-compatible [code-generator](https://github.com/deepmap/oapi-codegen) to generate Go types/interfaces from the openapi specs we write.

To run the app, execute the following command 
```
make run-local
```
This will start a local server on `http://localhost:8888`. `run-local` will first generate all the code, and the build the project into a binary including the latest commit hash, and then executes the binary in a docker container. The port that the app is running on will be exposed to the localhost on port 8888.

# CI
The project is using Github Actions for continuous integration. There are two CI workflows defined:

- Test that runs on each commit for all branches bar the master branch, as well as for pull requests made to the master branch.
- Test and then deploy that runs only for master branch.
  
# Test
Two main tests are involved:
- `make test-swagger`: test swagger yaml definition file
- `make test`: unit test custom go code.

# Risks
Depending on where the app will be deployed to for production, there maybe further API integration required. When `-local` flag is specified during the execution of the binary like in `make run-local`, it's using a built-in http server, but the app behaviour may vary depending on the production runtime setup.