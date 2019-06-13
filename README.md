# Pointer
[![Build Status](https://travis-ci.org/onrik/pointer.svg?branch=master)](https://travis-ci.org/onrik/pointer)
[![Coverage Status](https://coveralls.io/repos/github/onrik/pointer/badge.svg?branch=master)](https://coveralls.io/github/onrik/pointer?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/onrik/pointer)](https://goreportcard.com/report/github.com/onrik/pointer)
[![GoDoc](https://godoc.org/github.com/onrik/pointer?status.svg)](https://godoc.org/github.com/onrik/pointer)

Helpers for working with pointers.


## Examples

```go
package main

import (
    "time"

    "github.com/onrik/pointer"
)

type User struct {
    ID        int
    Name      string
    Age       *int
    UpdatedAt *time.Time
}

func main() {
  // Set helpers
  user := User{
      ID:        1,
      Name:      "John",
      Age:       pointer.Int(27),
      UpdatedAt: pointer.Time(time.Now()),
  }
  
  // Get helpers
  // instead of
  age := -1
  if user.Age != nil {
      age := *user.Age
  }
  
  // Use
  age := pointer.GetInt(user.Age, -1)
}
```
