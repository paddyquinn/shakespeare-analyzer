package main

import (
  "fmt"

  log "github.com/Sirupsen/logrus"
  "github.com/paddyquinn/shakespeare-analyzer/analyzer"
)

const link = "http://www.ibiblio.org/xml/examples/shakespeare/macbeth.xml"

func main() {
  play, err := analyzer.Analyze(link)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(play)
}
