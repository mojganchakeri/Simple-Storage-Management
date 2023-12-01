package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_uploadRetrieve(t *testing.T) {
	ts := setupTestServer()
	defer ts.Server.Close()

	t.Run("upload and retrieve by name", func(t *testing.T) {
		prepareDir()

		statusCode, err := ts.sendUploadRequest("test.txt", "1,2", "text")
		if err != nil {
			t.Fatalf(err.Error())
		}
		assert.Equal(t, statusCode, 200)

		err = ts.sendRetrieveRequestWithName("test.txt")
		if err != nil {
			t.Fatalf(err.Error())
		}

		err = extractZipFile()
		if err != nil {
			t.Fatalf(err.Error())
		}
	})

	t.Run("upload and retrieve by tag", func(t *testing.T) {
		prepareDir()

		statusCode, err := ts.sendUploadRequest("test.txt", "1,2", "text")
		if err != nil {
			t.Fatalf(err.Error())
		}
		assert.Equal(t, statusCode, 200)

		err = ts.sendRetrieveRequestWithTag([]string{"1"})
		if err != nil {
			t.Fatalf(err.Error())
		}

		err = extractZipFile()
		if err != nil {
			t.Fatalf(err.Error())
		}
	})

}
