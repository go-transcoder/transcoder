package smil

import "testing"

func TestSmil_createSmilFile(t *testing.T) {
	var f CreateSmilFile

	outputFile := "/home/namir/Projects/my-transcoding/projects/uploads/transcoder/sample_test_cousa"

	err := f.CreateSmilFile(outputFile)

	if err != nil {
		t.Fatalf("Error while creating smil file err: %v", err)
	}
}
