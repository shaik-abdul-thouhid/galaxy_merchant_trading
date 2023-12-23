# GALAXY MERCHANT TRADING

To read setting up environment, go to [Setup](#setup) section

## Table of Content

- [Problem](#problem)
- [Assumptions/Prerequisites](#assumptionsprerequisites)  
  - [Roman Numeral Aliasing](#roman-numeral-aliasing)
  - [Metal Credits](#metal-credits)
- [Input](#input)
- [Output](#output)
- [Setup](#setup)
- [Running The Application](#running-the-application)
- [Application Design](#application-design)
- [Testing](#testing)

## Problem

You decided to give up on earth after the latest financial collapse left 99.99% of the earth's population with 0.01% of the wealth. Luckily, with the scant sum of money that is left in your account, you are able to afford to rent a spaceship, leave earth, and fly all over the galaxy to sell common metals and dirt (which apparently is worth a lot). Buying and selling over the galaxy requires you to convert numbers and units, and you decided to write a program to help you.The numbers used for intergalactic transactions follows similar convention to the roman numerals and you have painstakingly collected the appropriate translation between them. Roman numerals are based on seven symbols:

Symbol Value\
I1\
V5\
X 10\
L 50\
C 100\
D 500\
M 1,000\
\
Numbers are formed by combining symbols together and adding the values. For example, MMVI is 1000 + 1000 + 5 + 1 = 2006. Generally, symbols are placed in order of value, starting with the largest values. When smaller values precede larger values, the smaller values are subtracted from the larger values, and the result is added to the total. For example MCMXLIV = 1000 + (1000 − 100) + (50 − 10) + (5 − 1) = 1944.
The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more. (They may appear four times if the third and fourth are separated by a smaller value, such as XXXIX.)

"D", "L", and "V" can never be repeated.
"I" can be subtracted from "V" and "X" only. "X" can be subtracted from "L" and "C" only. "C" can be subtracted from "D" and "M" only. "V", "L", and "D" can never be subtracted.
Only one small-value symbol may be subtracted from any large-value symbol.
A number written in Arabic numerals can be broken into digits. For example, 1903 is composed of 1, 9, 0, and 3. To write the Roman numeral, each of the non-zero digits should be treated separately. In the above example, 1,000 = M, 900 = CM, and 3 = III. Therefore, 1903 = MCMIII.

## Assumptions/Prerequisites

### Roman Numeral Aliasing

For the program to run, we have to first of all populate aliases along with the corresponding **VALID ROMAN NUMERALS** and *METALS* with **VALID ALIAS** stored.

ex:

```text
glob is I
prok is V
pish is X
tegj is L
```

Here, `glob` `prok` `pish` `tegj` are aliases for the corresponding Roman numerals.

Also, the assignments should be passed to a single line with the format `XXX is YYY` where `XXX` and `YYY` are case-sensitive.

### Metal Credits

And also we need to assign credits to *METALS* with corresponding Roman numeral aliases

ex:

```text
glob glob Silver is 34 Credits
glob prok Gold is 57800 Credits
pish pish Iron is 3910 Credits
```

Also, the assignments should be passed to a single line with the format `...XXX METAL_NAME is YYY Credits` where `XXX` `METAL_NAME` `YYY` are case-sensitive and the entire line **SHOULD** follow same pattern to

And also we need to assign credits to *METALS* with corresponding Roman numeral aliases

Here, the `glob glob Silver is 34 Credits` is parsed like this,

```text
glob = I // in roman numeral

glob glob = II // here the space between the glob is ignored.
```

So, the `glob glob` is equal to `II` in Roman, so `II Silver = 34 Credits`.

Since `II` is equal to `2 units`, then `1 unit` of Silver is,

```text
2 units of Silver = 34 Credits
1 uint of Silver = 34 / 2 Credits = 17 Credits
```

Similar logic is applied for other metals as well.

| Given Units (In Roman) | Value in Integer | Metal       | Total Credits | Credit Per Unit |
|------------------------|:----------------:|-------------|---------------|----------------:|
|glob glob = "II"        | 2                | **Silver**  | 34            | 17              |
|glob prok = "IV"        | 4                | **Gold**    | 57800         | 14450           |
|pish pish = "XX"        | 20               | **Iron**    | 3910          | 195.5           |

## Input

```text
glob is I
prok is V
pish is X
tegj is L
glob glob Silver is 34 Credits
glob prok Gold is 57800 Credits
pish pish Iron is 3910 Credits
how much is pish tegj glob glob ?
how many Credits is glob prok Silver ?
how many Credits is glob glob Gold ?
how many Credits is glob glob glob glob glob glob Gold ?
how many Credits is pish tegj glob Iron ?
Does pish tegj glob glob Iron has more Credits than glob glob Gold ?
Does glob glob Gold has less Credits than pish tegj glob glob Iron?
Is glob prok larger than pish pish?
Is tegj glob glob smaller than glob prok?
how much wood could a woodchuck chuck if a woodchuck could chuck wood ?
```

## Output

following is the output printed on the terminal after reading the queries from `sample_input.txt` file

```zsh
coffee_backend_assignment on  main [!?] via 🐹 v1.21.5
❯ go run ./main.go sample_input.txt

pish tegj glob glob is 42
glob prok Silver is 68 Credits
glob glob Gold is 28900 Credits
Requested number is in invalid format
pish tegj glob Iron is 8015.5 Credits
glob glob Gold has more Credits than pish tegj glob glob Iron   
glob prok is not larger than pish pish
tegj glob glob is larger than glob prok
I have no idea what you are talking about
```

If you want to run the application using an executable then,

```bash
go build -o ./destination/path/executable_name.go ./source/path/
```

and

```zsh
./destination/path/executable_name.go ./path/to/input/file.txt
```

This will emit the previous result

## Setup

To run the application, you need to have go compiler (version >= 1.21.5) setup in your local system and appended to the path.

To check whether your system has goc installed, run the command which should output result,

```text
❯ go version
go version go1.21.5 windows/amd64
```

## Running The Application

To run the application you need to setup the application in your local machine, for instruction [follow](#running-the-application)...

In the working directory make sure there is a file with input queries or anywhere in your local machine. Make sure to remember the relative path to the input file from the current working directory.

Run,

```zsh
go run ./main.go ./relative/path/to/input_file.txt
```

Replace `./relative/path/to/input_file.txt` with your file. Else if you want to use of existing sample inputs, run,

```zsh
go run ./main.go ./sample_input.txt
```

## Application Design

The application and code logic is split into different packages. **Object-Oriented** design pattern and paradigm is used.

### package main

where the `main` function is present. The file path passed to in the terminal command is parsed [here](./main.go#21)

### package utils

this is where additional functions for miscellaneous operations are present for reusability

For example: `RomanToInt64`, `Int64ToRoman` functions

### package parse

This is where the [`parseQueries`](./parseStrings/parse.go#9) struct is present along with its associated methods are present with private visibility

Only one method [`func (p *parseQueries) ParseLines(input string) (string, error)`](./parseStrings/parse.go#97) is exposed and rest all the logic is abstracted from the consumer side for object oriented approach.

## Testing

All possible test cases are covered for parsing valid and invalid inputs.

And tests fails even if a valid string is passed before assigning aliases and metal credits.

```zsh
coffee_backend_assignment on  main [!?] via 🐹 v1.21.5 
❯ go test -v ./parseStrings

=== RUN   TestParseString
    parse_test.go:46: Test Name: Parse Valid Roman Alias
    parse_test.go:46: Test Name: Parse InValid Roman Alias
    parse_test.go:46: Test Name: Parse Valid Roman Alias
    parse_test.go:46: Test Name: Parse Valid Roman Alias
    parse_test.go:46: Test Name: Parse Valid Roman Alias
    parse_test.go:46: Test Name: Calculating Credits before assigning
    parse_test.go:46: Test Name: Parse Valid Metal Credits with Roman Aliases
    parse_test.go:46: Test Name: Parse Valid Metal Credits with Roman Aliases
    parse_test.go:46: Test Name: Parse Valid Metal Credits with Roman Aliases
    parse_test.go:46: Test Name: Query valid metal credits
    parse_test.go:46: Test Name: Query invalid metal credits
    parse_test.go:46: Test Name: Query valid roman alias
    parse_test.go:46: Test Name: Query comparison of two metal's credits
    parse_test.go:46: Test Name: Invalid Input
--- PASS: TestParseString (0.00s)
PASS
ok      coffee/assessment/parseStrings  (cached)
```
