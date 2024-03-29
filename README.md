[![License: MIT](https://img.shields.io/badge/license-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Test Workflow](https://github.com/collatz-prefixes/collatz-prefixes-go/actions/workflows/tests.yml/badge.svg?branch=main)

# Collatz Prefixes

> A pattern among hailstone numbers. Read the [Gitbook](https://erhany96.gitbook.io/collatz-prefixes) for the theory.

## Usage

You can build the binary under `bin` via:

```sh
go build -o ./bin/collatz -ldflags "-s -w" ./cmd/main.go
```

or run it directly with:

```sh
go run ./cmd/main.go <function> <name>
```

The executable is A CLI that takes two arguments:

```sh
# Print the stopping time of a number
len <num>

# Print the sequence of a number
seq <num>

# Print the reduced sequence of a number
rdseq <num>

# Print Exponential Canonical Form of a number
ecf <num>

# Print the path of a number
path <num>

# Prints the mapping of this number to a prefix
map <num>

# Prints the mapping of the prefix of a number
pf-map <num>

# Print the prefix w.r.t RIPTree
pf-rip <num>

# Print the prefix w.r.t PIPTree
pf-pip <num>

# Find ECF iteratively via RIPTree + Prefix
ecf-pf-rip <num>

# Find ECF iteratively via PIPTree + Prefix
ecf-pf-pip <num>

# Find ECF iteratively via RIPTree + Path extensions
ecf-path-rip <num>

# Find ECF iteratively via PIPTree + Path extensions
ecf-path-pip <num>
```

## Testing

Run all tests with coverage:

```sh
go test ./... -cover
```
