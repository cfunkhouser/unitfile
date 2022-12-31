# SystemD Unit File Encoding for Go

The `unitfile` package provides utilities for encoding and decoding (certain) Go
structs with semantics as similar to `encoding/json` as possible.

This is a personal project I am using for some other hacking, and as an excuse
to learn ANTLR.

Currently, only decoding is supported.

## Using

1. Install the module using `go get funkhouse.rs/unitfile`
2. Read the documentation at <https://pkg.go.dev/funkhouse.rs/unitfile>
3. Create a `struct` and `Unmarshal` some Units

### Decoding

Though the goal is to be as similar to
[`encoding/json`](https://pkg.go.dev/encoding/json) as possible, the structure
of SystemD Unit files does not lend them to good generic encoding. I do not
recommend you use this library for that purpose. Instead, this library concerns
itself exclusively with at-most singly-nested structs. The usual structure of a
destination struct will contain the top-level Unit section options as struct
fields, and each subsequent section will be a struct.

For some well-known Unit and section types, see
[`wellknown.go`](./wellknown.go). For a simplified example, consider:

```go
type SimpleUnit struct {
  Description string
  Requires    []string
  Service struct {
    StartCommand string `unit:"ExecStart"`
    StopCommand  string `unit:"ExecStop"`
  }
}
```

Given this struct and [`ssh.service`](./testdata/ubuntu-22.04.1-ssh.service) (in
this case, the one distributed with Ubuntu 22.04.1), the result of unmarshalling

```go
var dest SimpleUnit
unitfile.Unmarshal(must(os.ReadFile("ssh.service")), &dest)
```

... will be a value like:

```go
SimpleUnit{
  Description: "OpenBSD Secure Shell server",
  Requires:    []string(nil),
  Service: struct {
    StartCommand string "unit:\"ExecStart\""
    StopCommand  string "unit:\"ExecStop\""
  }{
    StartCommand: "/usr/sbin/sshd -D $SSHD_OPTS",
    StopCommand:  "",
  },
}
```

## Developing

The details of the parser are hidden in
[`./internal/parser`](./internal/parser). Much of that package is generated by
ANTLR. If you plan to hack on the parser code, ensure you have ANTLR 4
installed, and then run `go generate ./...` from the root of this repository.

There are a number of tests, both for the decoder in the top-level package, and
in the parser package. Be sure that they all pass!
