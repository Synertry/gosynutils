# GoSynUtils

A collection of common functions and structs I use for my Golang projects.

## Motivation

This is a detached fork of my old repo [GoSysUtils](https://github.com/Synertry/GoSysUtils), mainly because of naming reasons.

Originally it was a collection of utilities for terminal and file handling, but it evolved into a more general-purpose library.
That's why I wanted to remove the "Sys" from the name. The best way with Go's package manager for this is to create a new repository.
Also, now I can adhere to Go's naming conventions for the repo and package names, because I was stuck with CamelCase from PowerShell.

Some of those functions like str.Reverse() sound trivial and should have been part of the standard library, but are not.
Others are workarounds and methods, which are refined and encapsulated enough to be exportable,
instead of being a one-time use in a project.

Additionally, with the new repository I can have a clean slate of dependencies and no mismatch of licensing.
My aim is 0 dependencies, even for testing.

## Package Overview

### Current

- [datastruct](./datastruct/): Structs and functions to handle complex data structures, like heaps and tries
- [enc](./enc/): Functions to handle serialization and deserialization of data, like base64, hex, and JSON
- [file](./file/): Functions to handle file operations, like copying, moving, and deleting files and directories
- [fspath](./fspath/): Functions to handle file system paths, like joining, splitting, and validating paths
- [gen](./gen/): Generator package to create random strings, numbers, slices, and other data types
- [math](./math/): Functions to handle mathematical operations, which are not covered by the standard library
- [self](./self/): Special package to handle the current executable, like getting its path and name
- [slice](./slice/): Functions to handle slice operations
- [str](./str/): Functions to handle string operations, like string building, reversing, and validation

### Planned

I will the rest of the packages from my previous library [GoSysUtils](https://github.com/Synertry/GoSysUtils) soon.
I could add them right now for the functionality, but I would still need to create unit tests for each.
Until it has complete feature parity with the old library, I will not stay below v1.0.0 for this current library.

## License

This repository is licensed under the Boost Software License 1.0. See [LICENSE](https://github.com/Synertry/GoSysUtils/blob/main/LICENSE)
