package thumbnails

import "testing"

func TestExtractThumbs_CreateThumbs(t *testing.T) {
	var extractThumbs ExtractThumbs

	err := extractThumbs.CreateThumbs("/home/namir/Projects/my-transcoding/projects/uploads/transcoder/sample_test.mp4", "/home/namir/Projects/my-transcoding/projects/uploads/transcoder/sample_test_cousa")

	if err != nil {
		t.Fatalf("Error while extracting images err: %v", err)
	}
}
