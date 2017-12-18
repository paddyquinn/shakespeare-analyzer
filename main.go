package main

import (
  log "github.com/Sirupsen/logrus"
  "github.com/paddyquinn/shakespeare-analyzer/analyzer"
)

const link = "http://www.ibiblio.org/xml/examples/shakespeare/macbeth.xml"

func main() {
  a := analyzer.NewAnalyzer()
  characters, err := a.Analyze(link)
  if err != nil {
    log.Fatal(err)
  }

  characters.Print()
}
