package test

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"store_service/api"
	"store_service/configs"
	"store_service/internal/models"
	"store_service/repository"
	"strings"
)

const (
	downloadDir  = "downloads/"
	downloadPath = downloadDir + "export.zip"

	dbHost = "localhost"
	dbPort = "3306"
	dbUser = "root"
	dbPass = "password"
	dbName = "store_test"
)

type testServer struct {
	Server *httptest.Server
}

func setupTestServer() testServer {

	// connect to database
	configs.Env.DBUser = dbUser
	configs.Env.DBPass = dbPass
	configs.Env.DBHost = dbHost
	configs.Env.DBPort = dbPort
	configs.Env.DBName = dbName

	repository.SetClient()
	repository.DBClient.DB.Exec("delete from storage where 1=1")
	repository.DBClient.DB.Exec("delete from tag where 1=1")

	// set storage env
	configs.Env.StoragePath = "store/"
	configs.Env.MaxStorageSize = 300
	configs.Env.MaxFileSize = 10

	// create test server
	return testServer{Server: httptest.NewServer(api.SetupServer())}
}

func (ts *testServer) sendUploadRequest(reqName, reqTag, reqType string) (int, error) {
	url := ts.Server.URL + "/api/v1/upload"
	method := "POST"
	filePath := "test.txt"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	part1, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return 0, err
	}
	_, err = io.Copy(part1, file)
	if err != nil {
		return 0, err
	}
	err = writer.WriteField("name", reqName)
	if err != nil {
		return 0, err
	}
	err = writer.WriteField("tag", reqTag)
	if err != nil {
		return 0, err
	}
	err = writer.WriteField("type", reqType)
	if err != nil {
		return 0, err
	}
	err = writer.Close()
	if err != nil {
		return 0, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	return 0, err
	// }

	return res.StatusCode, nil
}

func (ts *testServer) sendRetrieveRequestWithName(searchName string) error {
	url := ts.Server.URL + "/api/v1/retrieve"
	method := "POST"

	var request models.RequestRetreive
	request.Name = searchName
	reqBody, _ := json.Marshal(request)
	payload := strings.NewReader(string(reqBody))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = downloadFile(res)
	if err != nil {
		return err
	}
	return nil
}

func (ts *testServer) sendRetrieveRequestWithTag(searchTag []string) error {
	url := ts.Server.URL + "/api/v1/retrieve"
	method := "POST"

	var request models.RequestRetreive
	request.Tag = searchTag
	reqBody, _ := json.Marshal(request)
	payload := strings.NewReader(string(reqBody))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = downloadFile(res)
	if err != nil {
		return err
	}
	return nil
}

func downloadFile(res *http.Response) (err error) {
	// Create the file
	out, err := os.Create(downloadPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Writer the body to file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func extractZipFile() error {
	// Open the ZIP archive
	reader, err := zip.OpenReader(downloadPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	if len(reader.Reader.File) == 0 {
		return fmt.Errorf("empty zip file")
	}

	// Extract each file from the ZIP archive
	for _, f := range reader.Reader.File {
		// Create the destination file or directory

		var dstFile *os.File

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(f.Name, 0755)
			if err != nil {
				return err
			}
		} else {
			dstFile, err = os.Create(f.Name)
			if err != nil {
				return err
			}
			defer dstFile.Close()
		}

		// Write the file contents
		r, err := f.Open()
		if err != nil {
			continue
		}
		defer r.Close()

		if _, err := io.Copy(dstFile, r); err != nil {
			return err
		}
	}

	return nil
}

func prepareDir() error {
	// remove storage
	os.RemoveAll(configs.Env.StoragePath)
	os.RemoveAll(downloadDir)

	os.Mkdir(downloadDir, 0777)

	return nil
}
