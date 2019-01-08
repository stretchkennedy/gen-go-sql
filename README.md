go-gen-sql
==========

Generate boilerplate for interactive with a SQL database.

Installation
------------

```bash
go get -i github.com/stretchkennedy/go-gen-sql
```

Getting started
---------------

```go
// main.go

package main

//go:generate go-gen-sql Foo

type Foo struct {
  ID uint64
}

func main() {
    db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	r := data.NewFooRepo(db)
	err = r.Create(&data.Foo{})
    // ...
}
```

```bash
go generate
go run main.go
```

Running the example
-------------------

```bash
git clone https://github.com/stretchkennedy/go-gen-sql.git
cd go-gen-sql
go build -i
cd example
go generate ./...
go run main.go
```
