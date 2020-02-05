package main
/* saw this: https://www.dotnetperls.com/parseint-go
and this : https://stackoverflow.com/questions/17690776/how-to-add-pause-to-a-go-program */

import (
    "fmt"
    "time"
)

func main() {
	const PAUSE = 3.0
	t0 := time.Now()
	wait := time.Second
	time.Sleep(wait*PAUSE)
	t1 := time.Now()
    fmt.Println("Time elapsed for ", t1.Sub(t0))
}