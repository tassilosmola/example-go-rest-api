This project is a test project which contains basic backend functionalities and a clean code approach. 

## Structure
To provide a clean and consistent design, every package is defined in the 
following structure

```
├── package
│   ├── interfaces.go
│   ├── impl1.go
│   ├── impl1_test.go
│   ├── impl2.go
|   ├── impl2_test.go
|   ├── mock_Interface.go
```
## main.go
This is the main entry point of the application. The [Composition Root](https://freecontent.manning.com/dependency-injection-in-net-2nd-edition-understanding-the-composition-root/) is defined by the variable declarations. Those can be replaced with other interface implementations to satisfy `SOLID` principles.
You can run the programm by executing 
````
go run main.go
````

## REST
Routing is implemented with [Mux](https://github.com/gorilla/mux) to provide the REST endpoints. 
Under `/resources` you will find a [Postman](https://www.postman.com/postman/) collection with stored `CRUD` operations for testing the API. 

## Repository
For simplicity reasons, the repository is just an in memory collection of structs. It can be replaced by a database or any other storage.

## Service Layer
All business logic is in the service layer before sending operations to the repository layer. 

## Unit Tests
Unit Tests are defined directly in the package with the `_test` declaration to identify test files. This is a recommended method to keep packages self contained and independend of other packages. 

### Mockery
To generate mocks of the implemented interfaces, you can use [Mockery](https://github.com/vektra/mockery).
This useful tool is designed for auto generating the mocks of all interfaces with this command: 

```
go run github.com/vektra/mockery/v2 --all --inpackage
```

## Swagger Documentation
To document the REST API endpoints, the `code first` design strategy is used by this project. All documentation is handled within the code and can be modified if any changes in the codebase are necessary. 

This is a great advantage because you don´t have to update the swagger.yaml specification by hand and also have no huge amount of generated boilerplate code as if you specify the swagger file and generate go code with the `spec first` strategy. 

If a `spec first` design approach is desired, you can design your swagger.yaml and paste the components of the different endpoints as comments in the controller layer during implementation. 

For more information, please refer to [Spec from Source](https://goswagger.io/generate/spec.html) and [Source from Spec](https://goswagger.io/generate/requirements.html)

If you want to add a new endpoint in the controller, use the swagger comments to specify the yaml spec.

In the /api subfolder, the file `docs.go` is just for documenting the api and has no productive meaning. Inside the file there is swagger specific metadata for auto-generation and wrapper functions for documenting the return and input types. 

### Generating swagger.yaml
To generate a swagger yaml specification, a `Makefile` is provided in the root folder, which can be called with 
```
make swagger-generate
```

It is also possible with the included command
```
swagger generate spec -o ./swagger.yaml --scan-models
```

If you now run the program, the Swagger UI documentation is available in `/docs`



