package docs

// repository
/*
	Go programs are organized into packages.

	A Go repository typically contains only one module, located at the root of the repository.
*/

// Packages
/*
	A Go package can be split into multiple files, all residing within the same directory,
and all the files in the directory declare the package directive.
	At least one file with the .go extension must be present in a directory for it to be considered a package.
*/

// modules
/*
	Sometimes we want to group related packages together
and release them as a single unit because many programs that depend on one package might also use the other packages.
	A module is a collection of packages that are released, versioned, and distributed together.
	The "net" is a standard go module, and it groups different packages together - "http", "mail", "rpc", "url" and others.

	If we want to import packages from a module,
we must always reference the package import path with reference to the module path of the module that has the package.
	A module is identified by a module path,
which is declared in a go.mod file, together with information about the module’s dependencies.
*/

/*
	https://go.dev/doc/modules/layout
	https://go.dev/doc/modules/developing
	https://go.dev/wiki/Learn
*/
