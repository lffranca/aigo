package aws

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestClient_PreSign(t *testing.T) {
	bucket := os.Getenv("BUCKET")

	pkg, err := New(&bucket)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := context.Background()
	key := "datalake/company/testfile1.csv"
	contentType := "text/csv"

	path, err := pkg.PreSign(ctx, &key, &contentType)
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(*path)
}

func TestClient_Upload(t *testing.T) {
	bucket := os.Getenv("BUCKET")

	pkg, err := New(&bucket)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := context.Background()
	key := "datalake/company/testfile.csv"
	contentType := "text/csv"
	buff := bytes.NewBuffer([]byte(`test,,,,`))

	if err := pkg.Upload(ctx, &key, &contentType, buff); err != nil {
		t.Error(err)
		return
	}
}

func TestClient_ListObjects(t *testing.T) {
	bucket := os.Getenv("BUCKET")

	pkg, err := New(&bucket)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := context.Background()

	files, err := pkg.ListObjects(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	jsonResult, _ := json.Marshal(files)
	log.Println(string(jsonResult))
}
