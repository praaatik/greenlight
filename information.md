# Information

## Folder structure

- `bin` => all the compiled binaries are present here.
- `cmd/api` => all the application specific code is present here.
- `internal` => database interaction, data validation, sending emails and so on.
  > Go code in the `cmd/api` would import the ones from `internal` but never the other way around.
- `migrations` => SQL migrations
- `remote` => all the config files and scripts for the Prod server.

> Packages under the `internal` cannot be imported from outside the project.

## Health check endpoint

- The healthcheck endpoint has been added as a method in the application. This means that whenever we need to
  add any dependency, we can just add in the application and it'll be "injected" in the healthcheck.

## `httprouter`

- We have kept the routes in a separate file `routes.go`. A benifit to this is that we can easily access the
  routes during the tests.
- To get the parameters, we use the `ParamsFromContext()` function from the URL.
