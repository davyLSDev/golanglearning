# Bug list

## Bugs found
* **B-1.** 4-Oct-2019. Without enough command line parameters, program crashes
* **B-2.** 4-Oct-2019. Note difference between entering "./calculator 1 2 3 4 5", and "./calculator a;lksdjflk ;aklsjdf ;laksjdf ;aksdjf ;adfj". The first yeilds the default from the switch statement, whereas the second yeilds -
> Assume interractive calculator then.
For testing purposes, add 1 2, which results in:  3
lksdjflk: command not found
aklsjdf: command not found
laksjdf: command not found
aksdjf: command not found
adfj: command not found


## Error Handling that should be done
* **EH-1.** avoid division by zero
* **EH-2.** avoid overflowing the result

## Improvements needed
* **I-1.** make this work with float64 instead of integer
* **I-2.** get operands from the command line to do the calculation
* **I-3.** get command line parameters to select which calculation to do
* **I-4.** refactor to make a library
* **I-5.** add a gui to make a four banger calculator
* **I-6** make the interractive calculator portion work