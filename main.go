package main

import (
  "go-web/my"

  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
  my.Migrate()
}
