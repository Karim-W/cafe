# CAFE
Configuation Management for Environment Variables for Applications

## What is CAFE?
CAFE is a simple configuration management tool for environment variables for applications. It is designed to offer similar functionalities to [joi](https://joi.dev) and [yup](https://github.com/jquense/yup) but for environment variables in Golang Projects.

## Why CAFE?
CAFE is designed to be used in Golang projects that require environment variables to be validated and parsed. Cafe can Ensure that the Enviroment Variables are set and are of the correct type.

## How to use CAFE?
### Set-up
```go
// Setting Up The schema
config := cafe.New(
	cafe.Schema{
			"PORT":        cafe.Int("SERVER_PORT"),              // PORT is an integer that is set by the SERVER_PORT environment variable and is not required
			"DB_HOST":     cafe.String("DB_HOST").Require(),     // DB_HOST is a string that is required
			"DB_PORT":     cafe.Int("DB_PORT").Require(),        // DB_PORT is an integer that is required
			"DB_USER":     cafe.String("DB_USER").Require(),     // DB_USER is a string that is required
			"DB_PASSWORD": cafe.String("DB_PASSWORD").Require(), // DB_PASSWORD is a string that is required
			"DB_NAME":     cafe.String("DB_NAME").Require(),     // DB_NAME is a string that is required
		})
```
### Intializing and Validating
```go
 // err will be nil if all the environment variables are set and are of the correct type
 // if err is not nil, it will contain the error message check errs.go for more information
 err := s.Initialize()
```
### ALT: Set-up, Intialize and Validate
```go
config, err := New( // Creates a new schema and initializes it
		Schema{
			"PORT":        cafe.Int("SERVER_PORT"),              // PORT is an integer that is set by the SERVER_PORT environment variable and is not required
			"DB_HOST":     cafe.String("DB_HOST").Require(),     // DB_HOST is a string that is required
			"DB_PORT":     cafe.Int("DB_PORT").Require(),        // DB_PORT is an integer that is required
			"DB_USER":     cafe.String("DB_USER").Require(),     // DB_USER is a string that is required
			"DB_PASSWORD": cafe.String("DB_PASSWORD").Require(), // DB_PASSWORD is a string that is required
			"DB_NAME":     cafe.String("DB_NAME").Require(),     // DB_NAME is a string that is required
		},
	)
	if err != nil {
		t.Error(err)
	}
```
### Accessing the variables
```go
serverPort,err := config.Getcafe.Int("PORT") 
if err != nil {
	// handle error
}
```
## Roadmap
- [x] Support for Integers
- [x] Support for Strings
- [ ] Support for Floats
- [x] Support for Booleans
- [ ] Support for Arrays
- [ ] Support for Maps
- [ ] Support for Structs and nested objects
- [ ] Muliple Configurations Sources (JSON, YAML, TOML, etc)
