# datastruct/bench

This directory contains performance benchmarks comparing the internal `gosynutils/datastruct` Trie implementation against popular open-source Trie/Radix libraries:

- [`dghubble/trie`](https://github.com/dghubble/trie)
- [`sauerbraten/radix`](https://github.com/sauerbraten/radix)
- [`hashicorp/go-immutable-radix`](https://github.com/hashicorp/go-immutable-radix)

## Why a separate package?

To adhere to the zero-dependency philosophy of `gosynutils`, these benchmarks are isolated in their own `bench_test` package. This ensures that the main library remains completely free of external dependencies while still allowing rigorous performance testing of its components against community standards.

## Running the Benchmarks

Navigate to this directory and run the standard Go benchmark command:

```sh
go test -bench=. -benchmem
```
