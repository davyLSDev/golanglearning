# Hold the fort

*[this might be the best](https://github.com/quii/learn-go-with-tests)
*[daily katas](http://www.peterprovost.org/blog/2012/05/02/kata-the-only-way-to-learn-tdd/)
*[some katas](https://kata-log.rocks/starter)

## Another tutorial on TDD in golang

* this one might be better than the previous one I went through: [intro to TDD](https://tutorialedge.net/golang/intro-testing-in-go/)

## On using "testify"

* Testify helps you to simplify the way you write assertions within your test cases.
* Testify can also be used to mock objects within your testing framework to ensure you aren’t calling production endpoints whenever you test.

## Notes

### about golang

```
Go does not let you use equality operators with slices. The tutorial author suggests
that I could write a function to iterate over each got and want slice and check their
values, but for convenience, we can use "reflect.DeepEqual" which is useful for seeing
if any two variables are the same.
if got != want {
```

* reflect.DeepEqual is *NOT* type safe, the code will still compile even if you do s/th
silly. For instance comparing a slice to a string will still compile; this makes no sense.

* make allows to create a slice with a starting capacity

### godoc

Very useful utility which serves up go documentation for my system

godoc -http :8000
then navigate to http://localhost:8000/pkg/testing ... or wherever

### commit whenever you can get tests to pass

In case you mess something up, but maybe do not push to master right away.

* You can add [Examples](https://blog.golang.org/examples) in your test code which is also testable (or optionally not).package tdd
The example will not be executed it the "// Output: 6", for example is missing
* type "godoc -http:8000" if your code is within the $GOPATH/src/.github.-com/{your_id} your example documentation will be published to localhost:8000/pkg/