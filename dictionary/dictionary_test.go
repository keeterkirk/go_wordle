package dictionary

import "testing"

func TestRandomWord(t *testing.T) {
	word := RandomWord()
	if word == "" {
		t.Error("Expected a non-empty word, got an empty string")
	}
	if len(word) != 5 {
		t.Errorf("Expected a word with a character count of 5, got %d", len(word))
	}
}
