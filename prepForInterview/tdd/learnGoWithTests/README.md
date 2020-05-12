# Begin learning Go with tests

[URL](https://github.com/quii/learn-go-with-tests)

## Left off at

* the book is an epub in the directory
* the last page was 99 (finished up on arrays)
* where to pick up next. Review slices on the blog page first.

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

[abstraction from main for better testing](https://pace.dev/blog/2020/02/12/why-you-shouldnt-use-func-main-in-golang-by-mat-ryer.html)

But, it is not valid to write something like

```
func main() error {
	if err := Main(); err != nil {
		return err
	}
	return nil
}
```

The "main()" function in golang must not have arguments or return values.

[here is a way to do that](https://stackoverflow.com/questions/4278293/how-do-i-return-from-func-main-in-go)

```
func main() { os.Exit(mainReturnWithCode()) }

func mainReturnWithCode() int {
    // do stuff, defer functions, etc.
    return exitcode // a suitable exit code
}
```

However, this will only work if you ensure that rest of the code does not call os.Exit() anywhere, like flag.ExitOnError, log.Fatalf(), etc.

## Run specific tests

See structsMethodsInterface/shapes_test.go for more information about how to set that up. To run, for instance, only the rectangle test, type

```
go test -run TestArea/Rectangle
```