package analyzer

import (
  "encoding/xml"
  "net/http"
)

type Play struct {
  SpeechElements []SpeechElement `xml:"ACT>SCENE>SPEECH"`
}

type SpeechElement struct {
  Speaker string   `xml:"SPEAKER"`
  Lines   []string `xml:"LINE"`
}

func Analyze(link string) (*Play, error) {
  response, err := http.Get(link)
  if err != nil {
    return nil, err
  }
  // TODO: if an error is returned, is the response.Body always nil?
  defer response.Body.Close()

  decoder := xml.NewDecoder(response.Body)
  play := &Play{}
  decoder.Decode(play)

  return play, nil
}
