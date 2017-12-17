package analyzer

// The following 2 structs are unexported because they are only used intermediately in this package to go from XML to
// the Characters struct. The elements within each struct need to be exported for `decoder.Decode` to work.

type play struct {
	SpeechElements []speechElement `xml:"ACT>SCENE>SPEECH"`
}

type speechElement struct {
	Speaker string   `xml:"SPEAKER"`
	Lines   []string `xml:"LINE"`
}