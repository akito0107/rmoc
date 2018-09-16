package rmoc_test

import (
	"bytes"
	"github.com/akito0107/rmoc"
	"io"
	"os"
	"testing"
)

func TestCreateFileWithAbort(t *testing.T) {
	buf := bytes.NewBufferString("test")
	err:= rmoc.CreateFileWithAbort(buf, "testdata", "dummy")

	if ok, _ := rmoc.IsFileAlreadyExists(err); !ok {
		t.Errorf("error must be FileAlreadyExists byt %+v", err)
	}
}

func TestOverrideFile(t *testing.T) {
	func() {
		pre := bytes.NewBufferString("pre")
		out, err := os.Create("testdata/override.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer out.Close()
		io.Copy(out, pre)
	}()

	after := bytes.NewBufferString("after")
	if err := rmoc.OverrideFile(after, "testdata", "override.txt"); err != nil {
		t.Fatal(err)
	}

	f, err := os.Open("testdata/override.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var dist bytes.Buffer
	io.Copy(&dist, f)

	if dist.String() != "after" {
		t.Errorf("File contents must be after but %s", dist.String())
	}

}