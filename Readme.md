# Example of db integration tests
Code sample of solution that help you to run your test cases locally without any effort

### Testing

**Every pushed commit trigger github action that run tests for application and show result if tests were broken.** We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests locally:


```shell
go test -v ./... -> to run all tests
```
