# Collatz Prefixes

> A pattern among hailstone numbers. Read the [Gitbook](https://erhany96.gitbook.io/collatz-prefixes) for the theory.

## Usage

All functions operate on `*big.Int` pointers, but functions that may change the underlying value do NOT change them, a simple copy is made within the function. This is to make the interface compatible with other bigint functions, while also keeping things behave like pass-by-value. For example, this prevents your value to go to 1 when you compute the sequence.

## Testing

Run all tests with:

```sh
go test -v
```
