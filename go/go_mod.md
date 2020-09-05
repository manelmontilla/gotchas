# Go modules cheat sheet

`go get <pkg>`: pkg@latest (respecting transitive deps).

`go get -u <pkg>`: pkg@latest + last minor version of transitive deps.

`go get -u=patch <pkg>`: @latest + last patch version of transitive deps.

The flag `-t` would also add test dependencies.

Of course, you can specify a specific version of the package:

`go get -u <pkg>@<v1.2.3>`: pkg@v1.2.3 + last minor version of transitive deps.

On the other hand, when using the reserved path name "all":

`go get -u all`: @latest of ALL packages transitively imported + last minor
version of transitive deps.

From "go help packages":

> When using modules, "all" expands to all packages in the main module and
> their dependencies, including dependencies needed by tests of any of those.
