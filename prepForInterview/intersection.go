// findIntersection reads an array of strings containing two elements each
// of which is a list of comma-separated numbers sorted in ascending order.

// This function returns return a string of numbers that occur in both
// elements of the input array in sorted order. If there is no intersection,
// the function returns the string "false".

// An example
// ["1, 3, 4, 7, 13", "1, 2, 4, 13, 15"] the output string should be "1, 4, 13"
// because those numbers appear in both strings.
// The array given will not be empty, and each string inside the array will be the
// numbers sorted in ascending order. It may contain negative numbers.

// Improvement notes:
//      loop through the shortest string,
//      don't slurp the string into a string array
//      write a parser to fetch numbers from the array

// From https://dev.to/coderbyte/a-common-coding-interview-question-461f

package main

import (
    "fmt"
    "strings"
)

func main() {
    sortedNumbers := [2]string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"}

    fmt.Println("The intersection of the initial string arrays ", sortedNumbers[0], "and ", sortedNumbers[1], " is:")
    err, intersection := findIntersection(sortedNumbers)
    fmt.Println(intersection, " and the error returned is ", err)
}

func findIntersection(twoSortedNumberStrings [2]string) (err bool, result string) {
    matches := 0
    fieldSeparator := ", "
    sortedA := strings.Split(twoSortedNumberStrings[0], fieldSeparator)
    sortedB := strings.Split(twoSortedNumberStrings[1], fieldSeparator)
    var intersection string

    if len(sortedA) == 0 || len(sortedB) == 0 {
        err = true
    } else {
        err = false
// loop through string A
        lengthA := len(sortedA) - 1
        first := true
        last := false

        for idx, numberInA := range sortedA {
            if foundMatch(sortedB, numberInA) {
                matches++
                if matches != 1 {
                    first = false
                }
                if idx == lengthA {
                    last = true
                }
                intersection = stringBuilder(intersection, numberInA, first, last)
            }
        }
    }
    result = intersection

    if matches == 0 {
        result = "none!"
        err = true
    } else {
        result = intersection
    }
    return err, result
}

func foundMatch(stringToSearch []string, stringToFind string) (foundIt bool) {
    for _, oneString := range stringToSearch {
        if oneString == stringToFind {
            foundIt = true
        }
    }
    return
}

func stringBuilder(stringToBuild, stringItem string, first, last bool) (builtString string) {
    separator := ", "
    if first {
        builtString = stringItem + separator
    } else if last {
        builtString = stringToBuild + stringItem
    } else {
        builtString = stringToBuild + stringItem + separator
    }
    return builtString
}
