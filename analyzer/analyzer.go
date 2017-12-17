package analyzer

import (
  "encoding/xml"
  "net/http"
)

// Play is a struct that contains all of the speech elements from a play
type Play struct {
  SpeechElements []SpeechElement `xml:"ACT>SCENE>SPEECH"`
}

// SpeechElement is a struct that represents a line or group of lines a character speaks during the play
type SpeechElement struct {
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
  play := &Play{}
  decoder.Decode(play)

  characters := NewCharacters(play)
  return characters, nil
}
