package solution

import "testing"

func TestCamelCase(t *testing.T) {
	got := CamelCase("saveChangesInTheEditor")
	if got != 5 {
		t.Errorf("expected 5, got %d", got)
	}
}

func TestCaesarCipher(t *testing.T) {
	got := CaesarCipher("Always-Look-on-the-Bright-Side-of-Life", 5)
	if got != "Fqbfdx-Qttp-ts-ymj-Gwnlmy-Xnij-tk-Qnkj" {
		t.Errorf("expected \"Fqbfdx-Qttp-ts-ymj-Gwnlmy-Xnij-tk-Qnkj\", got %q", got)
	}
}
