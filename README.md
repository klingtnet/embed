# embed

![CI](https://github.com/klingtnet/embed/workflows/CI/badge.svg)

- [Documentation](https://pkg.go.dev/github.com/klingtnet/embed)
- [Releases](https://github.com/klingtnet/embed/releases)

embed is a tool for embedding static content in your Go application.

It provides three methods, listing embedded files and getting their content as `[]byte` or `string`.  If you need a `io.Writer` just wrap the `[]byte` content in a `bytes.NewBuffer`.

The motivation for building yet another static file embedding tool for Go was that I am not satisified with any of the existing tools, they either have inconvenient APIs or did not support to include more than a single folder or file.

Please note that this tool, as well as most other static file embedding tools, will be redundant as soon as the proposal to [add support for embedded files](https://github.com/golang/go/issues/41191) lands in `go/cmd`.

## Usage

![a demo of embed](https://golangleipzig.space/372959.gif)

You can run the tool with `go run github.com/klingtnet/embed/cmd/embed` or by downloading a precompiled binary from the [releases page](https://github.com/klingtnet/embed/releases).

```sh
$ ./embed
NAME:
   embed - A new cli application

USAGE:
   embed [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --package value, -p value                    name of the package the generated Go file is associated to (default: "main")
   --destination value, --dest value, -d value  where to store the generated Go file (default: "embeds.go")
   --include value, -i value                    paths to embed, directories are stored recursively (can be used multiple times)
   --help, -h                                   show help (default: false)
```

Running `embed --include assets --include views` will create a file `embeds.go` (you can change the destination) that bundles all files from the assets and views directory.  In your application you can then use `embeds.File("assets/my-asset.png")` to get the contents of an embedded file.  For an example of such a generated file see [`internal/embeds.go`](https://github.com/klingtnet/embed/blob/master/internal/embeds.go).

## golang-migrate driver

The package also provides a migration source driver for [golang-migrate](https://github.com/golang-migrate/migrate).
For a usage example refer to [`examples/migrate/migrate.go`](https://github.com/klingtnet/embed/blob/master/examples/migrate/migrate.go).
