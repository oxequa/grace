<p align="center">
  <img src="https://i.imgur.com/8b6ngnu.png" width="125px">
</p>
<p align="center">
  <a href="https://travis-ci.org/oxequa/grace"><img src="https://img.shields.io/travis/oxequa/grace.svg?style=flat-square" alt="Build status"></a>
  <a href="https://goreportcard.com/report/github.com/oxequa/grace"><img src="https://goreportcard.com/badge/github.com/oxequa/grace?style=flat-square" alt="GoReport"></a>
  <a href="http://godoc.org/github.com/oxequa/grace"><img src="http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square" alt="GoDoc"></a>
  <a href="https://raw.githubusercontent.com/oxequa/grace/master/LICENSE"><img src="https://img.shields.io/aur/license/yaourt.svg?style=flat-square" alt="License"></a>
  <a href="https://gitter.im/oxequa/grace?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge"><img src="https://img.shields.io/gitter/room/oxequa/grace.svg?style=flat-square" alt="Gitter"></a>
</p>
<hr>
<h3 align="center">Handle recover, panic and errors in a graceful way</h3>
<hr>

## Quickstart

```
go get github.com/oxequa/grace
```

The following is a simple example that handles a panic and returns the error without the program crash.

```go
package main

import (
  "fmt"
  "github.com/oxequa/grace"
)

func example() (e error){
  defer grace.Recover(&e) // save recover error and stack trace to e
  numbers := []int{1, 2}
  fmt.Println(numbers[3]) // panic out of index
  return
}

func main() {
  err := example() // no panic occur
  fmt.Println(err)
  fmt.Println("End")
}
```

## Documentation

You can read the full documentation of Grace [here](https://grace.oxequa.com).

## Contributing

Please read our guideline [here](CONTRIBUTING.md).

## License

Grace is licensed under the [GNU GENERAL PUBLIC LICENSE V3](LICENSE).
