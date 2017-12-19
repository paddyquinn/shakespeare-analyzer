package main

import (
  "fmt"
  "os"

  log "github.com/Sirupsen/logrus"
  "github.com/paddyquinn/shakespeare-analyzer/analyzer"
  "github.com/paddyquinn/shakespeare-analyzer/router"
)

const (
  consoleFlag = "-c"
  prompt      = "Enter link: "
  serverFlag  = "-s"
  usageString = "Usage: %s -c | -s\n\nOptions:\n\t-c\tconsole mode\n\t-s\tserver mode\n"
)

func main() {
  if len(os.Args) != 2 {
    usage(os.Args[0])
    return
  }

  if os.Args[1] == consoleFlag {
    fmt.Print(prompt)
    var link string
    fmt.Scanln(&link)

    a := analyzer.NewAnalyzer()
    characters, err := a.Analyze(link)
    if err != nil {
      log.Fatal(err)
    }

    characters.Print()
  } else if os.Args[1] == serverFlag {
    router.Run()
  } else {
    usage(os.Args[0])
  }
}

func usage(executable string) {
  fmt.Printf(usageString, executable)
}