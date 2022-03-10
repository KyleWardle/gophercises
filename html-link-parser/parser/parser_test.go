package parser

import (
	"testing"
)

func TestExampleOne(t *testing.T) {
	links := ParseLinks("examples/ex1.html")
	assertEquals(t, links[0].Href, "/other-page")
	assertEquals(t, links[0].Text, "A link to another page")
}

func TestExampleTwo(t *testing.T) {
	links := ParseLinks("examples/ex2.html")
	assertEquals(t, links[0].Href, "https://www.twitter.com/joncalhoun")
	assertEquals(t, links[0].Text, "Check me out on twitter")

	assertEquals(t, links[1].Href, "https://github.com/gophercises")
	assertEquals(t, links[1].Text, "Gophercises is on Github!")
}

func TestExampleThree(t *testing.T) {
	links := ParseLinks("examples/ex3.html")
	assertEquals(t, links[0].Href, "#")
	assertEquals(t, links[0].Text, "Login")

	assertEquals(t, links[1].Href, "/lost")
	assertEquals(t, links[1].Text, "Lost? Need help?")

	assertEquals(t, links[2].Href, "https://twitter.com/marcusolsson")
	assertEquals(t, links[2].Text, "@marcusolsson")
}

func TestExampleFour(t *testing.T) {
	links := ParseLinks("examples/ex4.html")
	assertEquals(t, links[0].Href, "/dog-cat")
	assertEquals(t, links[0].Text, "dog cat")
}

func assertEquals(t *testing.T, first string, second string) {
	if first != second {
		t.Errorf("Error asserting '%s' is equal to '%s'", first, second)
	}
}
