# atime

`atime` is an internal package used by the `file` package to handle file access times across different operating systems.

## Motivation

Instead of pulling in an external dependency, I extracted the necessary logic here to adhere to the zero-dependency goal of this repository. I also wrote dedicated tests for it so our overall test coverage doesn't drop.

## Attribution

Huge thanks to [djherbis/atime](https://github.com/djherbis/atime) for the original code. The underlying files and system interactions from Google are rock solid, so this package shouldn't need much updating going forward.