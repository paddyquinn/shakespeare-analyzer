package analyzer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCharacters(t *testing.T) {
	Convey("Given a play", t, func() {
		p := &play{
			SpeechElements: []speechElement{
				{
					Speaker: "Character One",
					Lines:   []string{"I only have one line."},
				},
				{
					Speaker: "character two",
					Lines:   []string{"I have two lines."},
				},
				{
					Speaker: "Character Three",
					Lines:   []string{"I have three lines."},
				},
				{
					Speaker: "CHARACTER TWO",
					Lines:   []string{"This is my second line."},
				},
				{
					Speaker: "cHARACTER tHREE",
					Lines:   []string{"This is my second line.", "This is my third line."},
				},
			},
		}

		Convey("a call to NewCharacters should return a Characters type sorted in order of number of lines with names correctly capitalized", func() {
			characters := NewCharacters(p)

			So(len(characters), ShouldEqual, 3)
			So(characters[0].Name, ShouldEqual, "Character Three")
			So(characters[0].NumLines, ShouldEqual, 3)
			So(characters[1].Name, ShouldEqual, "Character Two")
			So(characters[1].NumLines, ShouldEqual, 2)
			So(characters[2].Name, ShouldEqual, "Character One")
			So(characters[2].NumLines, ShouldEqual, 1)
		})
	})
}

func TestCharacters_Len(t *testing.T) {
	Convey("Given a Characters type of length 3", t, func() {
		characters := Characters{&Character{}, &Character{}, &Character{}}

		Convey("a call to Len should return 3", func () {
			So(characters.Len(), ShouldEqual, 3)
		})
	})
}

func TestCharacters_Less(t *testing.T) {
	Convey("Given a Characters type of length 3", t, func() {
		characters := Characters{
			&Character{
				Name:     "Character With One Line",
				NumLines: 1,
			},
			&Character{
				Name:     "Another Character With One Line",
				NumLines: 1,
			},
			&Character{
				Name:     "Character With Two Lines",
				NumLines: 2,
			},
		}

		Convey("a call to Less given the indices of two characters with the same number of lines should return false", func() {
			So(characters.Less(0, 1), ShouldBeFalse)
		})

		Convey("a call to Less given the indices of two characters with a different number of lines", func() {
			Convey("where the character at the first index has less lines should return true", func() {
				So(characters.Less(0, 2), ShouldBeTrue)
			})

			Convey("where the character at the first index has more lines should return false", func() {
				So(characters.Less(2, 0), ShouldBeFalse)
			})
		})
	})
}

func TestCharacters_Swap(t *testing.T) {
	Convey("Given a Characters type of length 2", t, func() {
		characters := Characters{
			&Character{
				Name:     "Character One",
				NumLines: 1,
			},
			&Character{
				Name:     "Character Two",
				NumLines: 2,
			},
		}

		Convey("a call to swap swaps the two characters at the indices passed", func() {
			characters.Swap(0, 1)

			So(characters[0].Name, ShouldEqual, "Character Two")
			So(characters[0].NumLines, ShouldEqual, 2)
			So(characters[1].Name, ShouldEqual, "Character One")
			So(characters[1].NumLines, ShouldEqual, 1)
		})
	})
}