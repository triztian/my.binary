# golang/go: issue #32886

This project reproduces the issue mentioned here in a public way:

 * https://github.com/golang/go/issues/32886

What this project aims to achieve if the issue were not present is to switch the implementation of a function based on a build constraint:

When perfoming a build normally:

```bash
go build
```

It should build the `cmd/func.go` file

When performing a build with the `some.constraint.with-dash` tag it should build the `cmd/func_stub.go` file.

```bash
go build -tags some.constraint.with-dash
```


