package description_extractor

import (
	"io/ioutil"
)

func ExtractDeckDescription(descriptionFilePath string) (string, error) {
	// Read the entire contents of the file into a byte slice
	bytes, err := ioutil.ReadFile(descriptionFilePath)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string and return it
	return string(bytes), nil
}
