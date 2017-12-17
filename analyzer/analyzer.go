package analyzer

import (
  "encoding/xml"
  "net/http"
)

// Analyze takes a link to an xml version of a play and returns a list of characters sorted by lines spoken
func Analyze(link string) (Characters, error) {
  response, err := http.Get(link)
  if err != nil {
    return nil, err
  }
  // TODO: if an error is returned, is the response.Body always nil?
  defer response.Body.Close()

  decoder := xml.NewDecoder(response.Body)
  p := &play{}
  decoder.Decode(p)

  characters := NewCharacters(p)
  return characters, nil
}
