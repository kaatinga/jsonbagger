package jsonbagger

import "testing"

func TestExtractJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"multiple jsons", "{\n\t\"title\": \"Convert Multi-line String\"\n}\n\nor\n\n{\n\t\"title\": \"Single-line Readable String\"\n}\n\nor\n\n{\n\t\"title\": \"Multi-line to Single-line\"\n}\n\nor\n\n{\n\t\"title\": \"Readable Single-line Text\"\n}\n\nor\n\n{\n\t\"title\": \"Maintain Readability String\"\n}\n\nor\n\n{\n\t\"title\": \"Single-line Format Text\"\n}\n\nor\n\n{\n\t\"title\": \"Multi-line to Single-line Readability\"\n}\n\nor\n\n{\n\t\"title\": \"Single-line Text Conversion\"\n}\n\nor\n\n{\n\t\"title\": \"Multi-line to Single-line Readability for Mistral AI\"\n}\n\nYou can choose any title that you prefer as long as it is short and descriptive.", "{\n\t\"title\": \"Convert Multi-line String\"\n}", false},
		{"no json", "You can choose any title that you prefer as long as it is short and descriptive.", "", true},
		{"json", "{\n\t\"title\": \"Income/Loss Dec 2023\"\n}", "{\n\t\"title\": \"Income/Loss Dec 2023\"\n}", false},
		{"json", "```json\n{\n\t\"title\": \"Income/Loss December 2023\"\n}\n```\n\nFor the provided OpenAPI specification, to find out the income or loss for December 2023, you would need to call the...", "{\n\t\"title\": \"Income/Loss December 2023\"\n}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractJSON() got:\n%s\nWant:\n%s\n", got, tt.want)
			}
		})
	}
}
