# Internal Packages

Code in the `internal/` directory is restricted and cannot be imported by external projects.

While it is perfectly fine to use internal packages for private helper functions, this repository is fundamentally a utility library. If an internal function solves a common problem or could be useful in multiple projects, you are strongly encouraged to promote it to a public, exportable package.