# Golang-from-Crash

A personal repository for learning **Go (Golang)** from the fundamentals through practical programming exercises.

This repository contains small Go programs, exercises, tests, and shell scripting tasks created while learning Go syntax, programming concepts, problem solving, and command-line development.

## About

This repository is my learning journey with Go.

The exercises start with basic programming concepts and gradually introduce more Go-specific features and problem-solving techniques.

The main goal is to:

* Learn Go syntax and conventions
* Practice problem solving
* Understand Go functions and data types
* Work with strings, integers, and floating-point values
* Practice loops and conditionals
* Learn slices and arrays
* Work with files and command-line tools
* Write unit tests
* Practice shell scripting
* Build a stronger foundation for larger Go projects

## Tech Stack

* **Go:** 1.27.1
* **Package:** `sprint`
* **Shell:** Bash
* **Testing:** Go testing package

The current Go module is:

```text
module sprint
```

## Repository Structure

The repository contains individual exercises rather than one large application.

```text
Golang-from-Crash/
│
├── *.go                 # Go exercises
├── *_test.go            # Go tests
├── *.sh                 # Shell scripting exercises
│
└── go.mod               # Go module definition
```

Some of the exercises currently included are:

```text
abacus.go
accumulate.go
alphabetmastery.go
alphanumber.go
balanceout.go
betweenlimits.go
bulkatoi.go
casting.go
combinations.go
countdivisible.go
countdown.go
doop.go
fibonacci.go
filterbysum.go
finddividend.go
generaterange.go
intvsfloat.go
isleapyear.go
isnegative.go
isprime.go
mean.go
pairs.go
pointdistance.go
removeelementsinrange.go
reversealphabet.go
reversealphabetvalue.go
runechecksum.go
season.go
shiftby.go
simplestrtoint.go
sortintegertable.go
strconcat.go
strtoint.go
timeconverter.go
```

There are also shell exercises such as:

```text
count.sh
files.sh
find.sh
introduction.sh
loop.sh
needles.sh
theDirectory.sh
theListing.sh
```

## Topics Practiced

### Go Fundamentals

The repository covers fundamental Go concepts including:

* Variables
* Constants
* Data types
* Type conversion
* Functions
* Parameters and return values
* Conditional statements
* `switch`
* `for` loops
* Integer operations
* Floating-point operations
* Strings
* Runes
* Arrays and slices

### Functions

The exercises use small, focused functions to practice solving individual problems.

For example, `Doop` implements basic arithmetic operations:

```go
func Doop(a int, op string, b int) int {
    switch op {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        if b == 0 {
            return 0
        }
        return a / b
    case "%":
        if b == 0 {
            return 0
        }
        return a % b
    default:
        return 0
    }
}
```

This exercise provides practice with:

* Functions
* `switch`
* Arithmetic operators
* Conditional logic
* Division by zero handling

### Strings and Runes

Several exercises work with strings, characters, and runes.

Examples include:

* String concatenation
* String-to-integer conversion
* Integer-to-string conversion
* Rune processing
* Alphabet manipulation
* Checksums

### Mathematical Problems

The repository also contains exercises involving:

* Fibonacci sequences
* Prime numbers
* Factorial-style calculations
* Averages
* Ranges
* Divisibility
* Distances
* Integer operations

### Testing

Some exercises include corresponding Go test files.

For example:

```text
abacus.go
abacus_test.go
```

Tests can be executed using Go's built-in testing framework.

## Running the Exercises

Clone the repository:

```bash
git clone https://github.com/Tin280/Golang-from-Crash.git
```

Move into the project:

```bash
cd Golang-from-Crash
```

Run an individual Go file:

```bash
go run filename.go
```

For example:

```bash
go run doop.go
```

## Running Tests

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

Run tests for a specific package:

```bash
go test
```

## Building

Build the current Go package with:

```bash
go build
```

You can also build a specific Go file:

```bash
go build doop.go
```

## Shell Exercises

The repository also contains Bash exercises created while learning command-line and shell scripting concepts.

Examples include:

```text
count.sh
files.sh
find.sh
introduction.sh
loop.sh
needles.sh
theDirectory.sh
theListing.sh
```

These exercises provide practice with:

* Shell variables
* Conditions
* Loops
* File operations
* Directory operations
* `case`
* Environment variables
* Command-line utilities

## Learning Progression

The repository represents a progression from simple exercises toward more complex programming tasks.

```text
Go Syntax
    ↓
Variables & Types
    ↓
Functions
    ↓
Conditions & Switch
    ↓
Loops
    ↓
Strings & Runes
    ↓
Arrays & Slices
    ↓
Algorithms & Problem Solving
    ↓
File & CLI Operations
    ↓
Testing
    ↓
Larger Go Projects
```

## Why This Repository Exists

This repository is primarily a **learning workspace**.

Instead of only reading Go documentation, I use small programming problems to practice concepts and understand how Go works through implementation.

The exercises also serve as a reference when moving from basic Go syntax toward larger projects involving:

* CLI applications
* Backend development
* APIs
* Testing
* Security
* System programming

## Current Focus

The long-term goal is to move from small exercises to complete Go applications with:

* Clean project structure
* Multiple packages
* Error handling
* File I/O
* Automated testing
* CLI interfaces
* APIs
* Concurrency
* Secure coding practices

## Projects Beyond the Exercises

This repository is the foundation for moving from individual Go exercises into larger projects.

Examples of future development areas include:

```text
Basic Go Exercises
        ↓
CLI Applications
        ↓
Data & File Processing
        ↓
HTTP / REST APIs
        ↓
Backend Services
        ↓
Testing & CI/CD
        ↓
Secure Go Applications
```

## Disclaimer

This repository contains learning exercises and experiments.

The code is primarily written for practicing Go concepts and understanding programming fundamentals rather than for production use.

## Author

**Tin Nguyen**

Learning Go one exercise at a time.

GitHub: [Tin280](https://github.com/Tin280)
