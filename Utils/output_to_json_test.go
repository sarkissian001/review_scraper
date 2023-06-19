package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestOutputToJSON(t *testing.T) {
	// Create a temporary file.
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		t.Fatalf("Cannot create temporary file: %s", err)
	}
	// Remember to clean up the file after we're done.
	defer os.Remove(tmpFile.Name())

	// Prepare some data.
	data := map[string]string{"foo": "bar"}

	// Use the function under test.
	err = OutputToJSON(data, tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to write to JSON file: %s", err)
	}

	// Now let's check the file's contents.
	content, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read file: %s", err)
	}

	// Unmarshal the file's contents into a map.
	var gotData map[string]string
	err = json.Unmarshal(content, &gotData)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON data: %s", err)
	}

	// Finally, let's assert that the data in the file matches what we expect.
	if gotData["foo"] != "bar" {
		t.Errorf("Expected 'bar', got '%s'", gotData["foo"])
	}
}
