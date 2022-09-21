# pg-commands

[![codecov](https://codecov.io/gh/habx/pg-commands/branch/dev/graph/badge.svg?token=YTMXFOJDCZ)](https://codecov.io/gh/habx/pg-commands)
[![Release](https://img.shields.io/github/v/release/habx/pg-commands)](https://github.com/habx/pg-commands/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/habx/pg-commands/dev)](https://golang.org/doc/devel/release.html)
[![CircleCI](https://img.shields.io/circleci/build/github/habx/pg-commands/dev)](https://app.circleci.com/pipelines/github/habx/pg-commands)
[![License](https://img.shields.io/github/license/habx/pg-commands)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/habx/pg-commands.svg)](https://pkg.go.dev/github.com/habx/pg-commands)

## install

```bash
go get -t github.com/habx/pg-commands
```

## Example

### Code


```go
dump, _ := pg.NewDump(&pg.Postgres{
    Host:     "localhost",
    Port:     5432,
    DB:       "dev_example",
    Username: "example",
    Password: "example",
})
dumpExec := dump.Exec(pg.ExecOptions{StreamPrint: false})
if dumpExec.Error != nil {
    fmt.Println(dumpExec.Error.Err)
    fmt.Println(dumpExec.Output)

} else {
    fmt.Println("Dump success")
    fmt.Println(dumpExec.Output)
}

restore, _ := pg.NewRestore(&pg.Postgres{
    Host:     "localhost",
    Port:     5432,
    DB:       "dev_example",
    Username: "example",
    Password: "example",
})
restoreExec := restore.Exec(dumpExec.File, pg.ExecOptions{StreamPrint: false})
if restoreExec.Error != nil {
    fmt.Println(restoreExec.Error.Err)
    fmt.Println(restoreExec.Output)

} else {
    fmt.Println("Restore success")
    fmt.Println(restoreExec.Output)

}
```

### Lab

```
$ cd examples
$ docker-compose up -d
$ cd ..
$ POSTGRES_USER=example POSTGRES_PASSWORD=example POSTGRES_DB=postgres  go run tests/fixtures/scripts/init-database/init-database.go

$ go run main.go
Dump success

Restore success

```
