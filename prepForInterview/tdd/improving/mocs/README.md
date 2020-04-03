#Notes:

* The mock example does not work straight out of the box!
* See also the section on how to generate mock objects to save time
* For further study: [Advanced testing in go](https://tutorialedge.net/golang/advanced-go-testing-tutorial/)

##Generating mocs
So, in the above example we mocked out all of the various methods ourselves, but in real-life examples, this
may represent a h*** of a lot of different methods and functions to mock.

Thankfully, this is where the [vektra/mockery](https://github.com/vektra/mockery) package comes to our aide.

The mockery binary can take in the name of any interfaces you may have defined within your Go packages and it’ll automatically output the generated mocks to mocks/InterfaceName.go. This is seriously handy when you want to save yourself a tonne of time and it’s a tool I would highly recommend checking out!