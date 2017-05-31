package ligacecbank

import (
	"testing"
	"os"
)

const FILE_PATH string = "resources/ligacecbank.html"
type FaceFixture struct {}

func TestGetFixtures(t *testing.T) {
	expectedSize := 17

	file, err := os.Open(FILE_PATH)

	if nil != err {
		t.Fatalf("Cannot open the file %s", FILE_PATH)
	}

	defer file.Close()

	repository, err := WithDataStream(file)

	if nil != err {
		t.Fatalf("Could not create repository")
	}

	fixtures := repository.GetFixtures()

	if expectedSize != len(fixtures) {
		t.Fatalf("Excepted %d and got %d", expectedSize, len(fixtures))
	}

}


