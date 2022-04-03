### Testing

**Every pushed commit trigger github action that run tests for application and show result if tests were broken.** We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests locally:


```shell
go test -v ./... -> to run all tests
```
# Generating mocks
To generate mock you need to download generator library named mockery and type in the terminal command(ExampleService is name of the class Interface).

```shell
// 1st step download module
go get github.com/vektra/mockery/.../
// 2nd generate mock --name should be equal to Interface name
mockery -name ExampleService
```
