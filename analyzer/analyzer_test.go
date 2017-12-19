package analyzer

import (
  "errors"
  "io/ioutil"
  "strings"
  "testing"

  "github.com/paddyquinn/shakespeare-analyzer/analyzer/dao"
  . "github.com/smartystreets/goconvey/convey"
)

func TestAnalyzer_Analyze(t *testing.T) {
  Convey("Given an analyzer", t, func() {
    mock := &dao.MockDao{}
    a := &Analyzer{Dao: mock}

    Convey("that is passed an invalid link", func() {
      mock.On("GetXMLStream", "invalidlink.com").Return(nil, errors.New("http.Get error"))
      characters, err := a.Analyze("invalidlink.com")

      Convey("there should be an error", func() {
        So(err, ShouldNotBeNil)
        So(err.Error(), ShouldEqual, "http.Get error")
      })

      Convey("no characters should be returned", func() {
        So(characters, ShouldBeNil)
      })

      mock.AssertExpectations(t)
    })

    Convey("that is passed a valid link", func() {
      Convey("that returns invalid xml", func() {
        xmlString := `<PLAY>
                        <ACT>
                          <SCENE>
                            <SPEECH>
                              <SPEAKER>Character One
                              <LINE>I Only have one line. </LINE>
                            </SPEECH>
                            <SPEECH>
                              <SPEAKER>Character Two</SPEAKER>
                              <LINE>I have two lines!</LINE>
                              <LINE>This is my second line.</LINE>
                            </SPEECH>
                          </SCENE>
                        </ACT>
                      </PLAY>`
        xmlStream := ioutil.NopCloser(strings.NewReader(xmlString))
        mock.On("GetXMLStream", "validlink.com").Return(xmlStream, nil)
        characters, err := a.Analyze("validlink.com")

        Convey("there should be an error", func() {
          So(err, ShouldNotBeNil)
          So(err.Error(), ShouldEqual, "XML syntax error on line 7: element <SPEAKER> closed by </SPEECH>")
        })

        Convey("no characters should be returned", func() {
          So(characters, ShouldBeNil)
        })

        mock.AssertExpectations(t)
      })

      Convey("that returns valid xml", func() {
        xmlString := `<PLAY>
                        <ACT>
                          <SCENE>
                            <SPEECH>
                              <SPEAKER>Character One</SPEAKER>
                              <LINE>I Only have one line. </LINE>
                            </SPEECH>
                            <SPEECH>
                              <SPEAKER>Character Two</SPEAKER>
                              <LINE>I have two lines!</LINE>
                              <LINE>This is my second line.</LINE>
                            </SPEECH>
                          </SCENE>
                        </ACT>
                      </PLAY>`
        xmlStream := ioutil.NopCloser(strings.NewReader(xmlString))
        mock.On("GetXMLStream", "validlink.com").Return(xmlStream, nil)
        characters, err := a.Analyze("validlink.com")

        Convey("there should not be an error", func() {
          So(err, ShouldBeNil)
        })

        Convey("the characters should be returned in order of lines spoken", func() {
          So(characters[0].Name, ShouldEqual, "Character Two")
          So(characters[0].NumLines, ShouldEqual, 2)
          So(characters[1].Name, ShouldEqual, "Character One")
          So(characters[1].NumLines, ShouldEqual, 1)
        })

        mock.AssertExpectations(t)
      })

      Convey("that returns non-xml data", func() {
        jsonString := `{
                         "PLAY": {
                           "ACT": {
                             "SCENE": {
                               "SPEECH": [
                                 {
                                   "SPEAKER": "Character One",
                                   "LINE": [
                                     "I Only have one line. "
                                   ]
                                 },
                                 {
                                   "SPEAKER": "Character Two",
                                   "LINE": [
                                     "I have two lines!",
                                     "This is my second line."
                                   ]
                                 }
                               ]
                             }
                           }
                         }
                       }`
        jsonStream := ioutil.NopCloser(strings.NewReader(jsonString))
        mock.On("GetXMLStream", "validlink.com").Return(jsonStream, nil)
        characters, err := a.Analyze("validlink.com")

        Convey("there should be an error", func() {
          So(err, ShouldNotBeNil)
          So(err.Error(), ShouldEqual, "EOF")
        })

        Convey("no characters should be returned", func() {
          So(characters, ShouldBeNil)
        })

        mock.AssertExpectations(t)
      })
    })
  })
}