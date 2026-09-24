package main

import "testing"

func TestProcessText(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "cap single",
			input: "this is a test (cap) end",
			want:  "this is a Test end",
		},
		{
			name:  "cap numbered",
			input: "this is a test (cap, 3) end",
			want:  "this Is A Test end",
		},
		{
			name:  "up single",
			input: "hello world (up) today",
			want:  "hello WORLD today",
		},
		{
			name:  "low single",
			input: "LOUD noise (low) here",
			want:  "LOUD noise here",
		},
		{
			name:  "hex",
			input: "1E (hex) files were added",
			want:  "30 files were added",
		},
		{
			name:  "bin",
			input: "It has been 10 (bin) years",
			want:  "It has been 2 years",
		},
		{
			name:  "a to an vowel",
			input: "There it was. A amazing rock!",
			want:  "There it was. An amazing rock!",
		},
		{
			name:  "punctuation spacing",
			input: "I was sitting over there ,and then BAMM !!",
			want:  "I was sitting over there, and then BAMM!!",
		},
		{
			name:  "ellipsis stays glued",
			input: "I was thinking ... You were right",
			want:  "I was thinking... You were right",
		},
		{
			name:  "quote pairing single word",
			input: "I am exactly how they describe me: ' awesome '",
			want:  "I am exactly how they describe me: 'awesome'",
		},
		{
			name:  "quote pairing multi word",
			input: "As Elton John said: ' I am the most well-known thing in the world '",
			want:  "As Elton John said: 'I am the most well-known thing in the world'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processText(tt.input)
			if got != tt.want {
				t.Errorf("processText(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
