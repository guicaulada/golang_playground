// SRP - Single Responsibility Principle
package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

// the Journal should have the single responsibility
// of storing and removing entries, and maybe listing them as a string
type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// here we break separation of concerns because
// Journal shouldn't handle persistency
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// ...
}

// this is another way of doing this while respecting
// separation of concerns
var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	content := strings.Join(j.entries, LineSeparator)
	_ = ioutil.WriteFile(filename, []byte(content), 0644)
}

// and another way
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	content := strings.Join(j.entries, p.lineSeparator)
	_ = ioutil.WriteFile(filename, []byte(content), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I had fun today")
	j.AddEntry("I ate a cake")
	fmt.Println(j.String())
	SaveToFile(&j, "journal.txt")
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
