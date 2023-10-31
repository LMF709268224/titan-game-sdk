package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/zscboy/titan-game-sdk/storage/memfile"
)

func TestCalculateCarCID(t *testing.T) {
	f, err := os.Open("./example/main.go")
	if err != nil {
		t.Fatal(err)
	}

	cid, err := CalculateCid(f)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("cid ", cid.String())
}

func TestCreateCarWithFile(t *testing.T) {
	// }
	input := "./example/example.exe"
	output := "./example/example.car"

	root, err := createCar(input, output)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("root ", root.String())

}

func TestCreateCarStream(t *testing.T) {
	f, err := os.Open("./example/example.exe")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	// carFile, err := os.Create("./example/example.car")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer carFile.Close()

	mFile := memfile.New([]byte{})
	root, err := createCarStream(f, mFile)
	if err != nil {
		t.Fatal(err)
	}

	mFile.Seek(0, 0)

	// stat, err := carFile.Stat()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// buf := mFile.Bytes()
	t.Log("car file size:", len(mFile.Bytes()))

	testFile, err := os.Create("./example/test")
	if err != nil {
		t.Fatal(err)
	}
	defer testFile.Close()

	io.Copy(testFile, mFile)

	stat, err := testFile.Stat()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("test file size:", stat.Size())
	t.Log("root ", root.String())
}

func TestUpload(t *testing.T) {
	const apiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJ1c2VyIl0sIklEIjoiMTA1MjQ0MTYwN0BxcS5jb20iLCJOb2RlSUQiOiIiLCJFeHRlbmQiOiIifQ.Yjoxg9JA7SuikMFL0hHMtOANH1CD2v3JKbpkhSC88XQ"
	const locatorURL = "https://120.79.221.36:5000/rpc/v0"
	storage, close, err := NewStorage(locatorURL, apiKey)
	if err != nil {
		t.Fatal("NewStorage error ", err)
	}
	defer close()

	progress := func(doneSize int64, totalSize int64) {
		t.Logf("upload %d of %d", doneSize, totalSize)
	}

	filePath := "./"
	visitFile := func(fp string, fi os.DirEntry, err error) error {
		// Check for and handle errors
		if err != nil {
			fmt.Println(err) // Can be used to handle errors (e.g., permission denied)
			return nil
		}
		if fi.IsDir() {
			return nil
		} else {
			// This is a file, you can perform file-specific operations here
			if strings.HasSuffix(fp, ".go") {
				path, err := filepath.Abs(fp)
				if err != nil {
					t.Fatal(err)
				}
				_, err = storage.UploadFile(context.Background(), path, progress)
				if err != nil {
					t.Log("upload file failed ", err.Error())
					return nil
				}

				t.Logf("totalSize %s success", fp)
			}

		}
		return nil
	}

	err = filepath.WalkDir(filePath, visitFile)
	if err != nil {
		t.Fatal("WalkDir ", err)
	}
}

func TestUploadStream(t *testing.T) {
	const apiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJ1c2VyIl0sIklEIjoiMTA1MjQ0MTYwN0BxcS5jb20iLCJOb2RlSUQiOiIiLCJFeHRlbmQiOiIifQ.Yjoxg9JA7SuikMFL0hHMtOANH1CD2v3JKbpkhSC88XQ"
	const locatorURL = "https://120.79.221.36:5000/rpc/v0"
	storage, close, err := NewStorage(locatorURL, apiKey)
	if err != nil {
		t.Fatal("NewStorage error ", err)
	}
	defer close()

	progress := func(doneSize int64, totalSize int64) {
		t.Logf("upload %d of %d", doneSize, totalSize)
	}

	filePath := "./storage_test.go"
	visitFile := func(fp string, fi os.DirEntry, err error) error {
		// Check for and handle errors
		if err != nil {
			fmt.Println(err) // Can be used to handle errors (e.g., permission denied)
			return nil
		}
		if fi.IsDir() {
			return nil
		} else {
			// This is a file, you can perform file-specific operations here
			if strings.HasSuffix(fp, ".go") {
				path, err := filepath.Abs(fp)
				if err != nil {
					t.Fatal(err)
				}

				f, err := os.Open(path)
				if err != nil {
					t.Fatal(err)
				}

				_, err = storage.UploadStream(context.Background(), f, progress)
				if err != nil {
					t.Log("upload file failed ", err.Error())
					return nil
				}

				t.Logf("totalSize %s success", fp)
			}

		}
		return nil
	}

	err = filepath.WalkDir(filePath, visitFile)
	if err != nil {
		t.Fatal("WalkDir ", err)
	}
}

func TestReadWrite(t *testing.T) {
	data := []byte("111111111111111111111111111111111111")
	carFile, err := os.Create("./example/example.car")
	if err != nil {
		t.Fatal(err)
	}
	defer carFile.Close()

	// mFile := memfile.New([]byte{})
	// root, err := createCarStream(f, carFile)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	carFile.Write(data)

	carFile.Seek(0, 0)

	stat, err := carFile.Stat()
	if err != nil {
		t.Fatal(err)
	}
	// buf := mFile.Bytes()
	t.Logf("buf len %d, car file size:%d", len(data), stat.Size())

	testFile, err := os.Create("./example/test")
	if err != nil {
		t.Fatal(err)
	}
	defer testFile.Close()

	io.Copy(testFile, carFile)

	stat, err = testFile.Stat()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("test file size:", stat.Size())
}
