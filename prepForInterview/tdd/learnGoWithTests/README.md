# Begin learning Go with tests

[URL](https://github.com/quii/learn-go-with-tests)

## Some other resources

[tips](https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742)
[basics](https://blog.alexellis.io/golang-writing-unit-tests/])
[simplicied framework](https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c)
[subtests & benchmarks](https://blog.golang.org/subtests)
[test coverage](https://blog.golang.org/cover)


## Test coverage metrics

[caveat emptor](https://stackoverflow.com/questions/90002/what-is-a-reasonable-code-coverage-for-unit-tests-and-why)

So to run the tests and check the coverage, make sure you are in the code directory and type
 
 ```
 go test -cover
 ```

TO view the results type

```
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
```

But more interesting, and perhaps more useful visual representation type

```
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```




