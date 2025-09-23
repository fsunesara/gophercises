package parser

import (
	"testing"
)

func TestParseHTMLEx1(t *testing.T) {
	got := ParseHTML("ex1.html")
	if len(got) != 1 {
		t.Errorf("len(abs) = %d, expected %d", len(got), 1)
	}

	if got[0].Href != "/other-page" {
		t.Errorf("href = %s, expected %s", got[0].Href, "\"/other-page\"")
	}

	if got[0].Text != "A link to another page" {
		t.Errorf("text = %s, expected %s", got[0].Href, "\"A link to another page\"")
	}
}

func TestParseHTMLEx2(t *testing.T) {
	expectedHrefs := []string{"https://www.twitter.com/joncalhoun", "https://github.com/gophercises"}
	expectedTexts := []string{"Check me out on twitter", "Gophercises is on Github !"}
	got := ParseHTML("ex2.html")
	if len(got) != 2 {
		t.Errorf("len(abs) = %d, expected %d", len(got), 2)
	}

	for i := range expectedHrefs {
		if got[i].Href != expectedHrefs[i] {
			t.Errorf("href at index %d = %s, expected \"%s\"", i, got[i].Href, expectedHrefs[i])
		}
	}

	for i := range expectedTexts {
		if got[i].Text != expectedTexts[i] {
			t.Errorf("text at index %d = %s, expected \"%s\"", i, got[i].Text, expectedTexts[i])
		}
	}
}
