# FileJanitor

FileJanitor is a Go library for automated file management based on user-defined policies. It helps to clean up files by scanning directories, filtering based on extensions, and deleting older files beyond a specified retention period.

## Features

- Scheduled file scanning and cleanup
- Configurable file retention policies
- Support for file extension filtering
- Simple API to start and stop the service

## Installation

To install FileJanitor, use `go get`:

```sh
go get github.com/toxyl/filejanitor
```

## Usage

For an usage example, have a look at `example/main.go`. Run it with `go run -C example . `.

## License

This project is licensed under the UNLICENSE - see the [LICENSE](LICENSE) file for details.
