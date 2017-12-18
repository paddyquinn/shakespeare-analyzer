package analyzer

import (
  "encoding/xml"

  "github.com/paddyquinn/shakespeare-analyzer/analyzer/dao"
)

// Analyzer is a struct that uses its data access object to analyze the number of lines characters have in a play
type Analyzer struct {
  Dao dao.Interface
}

// NewAnalyzer returns an analyzer with a real data access object
func NewAnalyzer() *Analyzer {
  return &Analyzer{Dao: &dao.Dao{}}
}

// Analyze takes a link to an xml version of a play and returns a list of characters sorted by lines spoken
func (a *Analyzer) Analyze(link string) (Characters, error) {
  body, err := a.Dao.GetXMLStream(link)
  if err != nil {
    return nil, err
  }
  defer body.Close()

  decoder := xml.NewDecoder(body)
  play := &play{}
  err = decoder.Decode(play)
  if err != nil {
    return nil, err
  }

  characters := NewCharacters(play)
  return characters, nil
}
