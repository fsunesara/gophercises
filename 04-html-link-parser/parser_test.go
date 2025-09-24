package parser

import (
	"testing"
)

func TestParseHTMLEx1(t *testing.T) {
	testParseHTMLGeneric(t, "ex4.html", 1, []string{"/other-page"}, []string{"A link to another page"})
}

func TestParseHTMLEx2(t *testing.T) {
	expectedLen := 2
	expectedHrefs := []string{"https://www.twitter.com/joncalhoun", "https://github.com/gophercises"}
	expectedTexts := []string{"Check me out on twitter", "Gophercises is on Github !"}
	testParseHTMLGeneric(t, "ex2.html", expectedLen, expectedHrefs, expectedTexts)
}

func TestParseHTMLEx3(t *testing.T) {
	expectedLen := 3
	expectedHrefs := []string{"#", "/lost", "https://twitter.com/marcusolsson"}
	expectedTexts := []string{"Login", "Lost? Need help?", "@marcusolsson"}
	testParseHTMLGeneric(t, "ex4.html", expectedLen, expectedHrefs, expectedTexts)
}

func TestParseHTMLEx4(t *testing.T) {
	testParseHTMLGeneric(t, "ex4.html", 1, []string{"/dog-cat"}, []string{"dog cat"})
}

func testParseHTMLGeneric(t *testing.T, fileName string, expectedLen int, expectedHrefs []string, expectedTexts []string) {
	got, err := ParseHTML(fileName)
	if err != nil {
		t.Errorf("error occurred: %s", err.Error())
	}

	if len(got) != expectedLen {
		t.Errorf("len(abs) = %d, expected %d", len(got), expectedLen)
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
