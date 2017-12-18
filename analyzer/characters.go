package analyzer

import (
  "fmt"
  "sort"
  "strings"
)

const (
  all             = "All"
  characterFormat = "%d\t%s\n"
)

// Character is a struct that contains a character's name and the number of lines they have
type Character struct {
  Name     string
  NumLines int
}

// Print prints a character in the correct format
func (c *Character) Print() {
  fmt.Printf(characterFormat, c.NumLines, c.Name)
}

// Characters is a slice of Character structs
type Characters []*Character

// NewCharacters creates a Characters type from a play, with the characters sorted by the number of lines they have
func NewCharacters(play *play) Characters {
  characterMap := map[string]int{}
  for _, element := range play.SpeechElements {
    speaker := strings.Title(strings.ToLower(element.Speaker))
    if speaker == all {
      continue
    }

    // If the character is not found in the map numLines will default to 0
    numLines := characterMap[speaker]
    characterMap[speaker] = numLines + len(element.Lines)
  }

  characters := Characters{}
  for name, numLines := range characterMap {
    characters = append(characters, &Character{Name: name, NumLines: numLines})
  }
  sort.Sort(sort.Reverse(characters))

  return characters
}

// Print prints each character
func (c Characters) Print() {
  for _, character := range c {
    character.Print()
  }
}

// The following three functions are used to implement the sort.Interface interface so the Characters type can be sorted

// Len returns the length of the characters slice
func (c Characters) Len() int {
  return len(c)
}

// Less returns whether or not the number of lines the character at position i spoke was less  than the number of lines
// the character at position j spoke
func (c Characters) Less(i, j int) bool {
  return c[i].NumLines < c[j].NumLines
}

// Swap swaps the characters at positions i and j
func (c Characters) Swap(i, j int) {
  c[i], c[j] = c[j], c[i]
}