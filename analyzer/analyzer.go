package analyzer

import (
  "encoding/xml"
  "net/http"
)

// The following 2 structs are unexported because they are only used intermediately in this package to go from XML to
// the Characters struct. The elements within each struct need to be exported for `decoder.Decode` to work.

type play struct {
  SpeechElements []speechElement `xml:"ACT>SCENE>SPEECH"`
}

type speechElement struct {
  Speaker string   `xml:"SPEAKER"`
  Lines   []string `xml:"LINE"`
}

// Analyze takes a link to an xml version of a play and returns a list of characters sorted by lines spoken
func Analyze(link string) (Characters, error) {
  response, err := http.Get(link)
  if err != nil {
    return nil, err
  }
  // TODO: if an error is returned, is the response.Body always nil?
  defer response.Body.Close()

  decoder := xml.NewDecoder(response.Body)
  play := &play{}
  decoder.Decode(play)

  characters := NewCharacters(play)
  return characters, nil
}
