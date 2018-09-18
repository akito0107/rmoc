package rmoc_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/akito0107/rmoc"
)

func TestCreateFileWithAbort(t *testing.T) {
	buf := bytes.NewBufferString("test")
	err := rmoc.CreateFileWithAbort(buf, "testdata", "dummy")

	if ok, _ := rmoc.IsFileAlreadyExists(err); !ok {
		t.Errorf("error must be FileAlreadyExists byt %+v", err)
	}
}

func TestOverrideFile(t *testing.T) {
	filename := "testdata/override.txt"
	ts := fmt.Sprint(time.Now())
	func() {
		if _, err := os.Stat(filename); !os.IsNotExist(err) {
			os.Remove(filename)
		}
		pre := bytes.NewBufferString("pre")
		out, err := os.Create(filename)
		if err != nil {
			t.Fatal(err)
		}
		defer out.Close()
		io.Copy(out, pre)
	}()

	after := bytes.NewBufferString(ts)
	if err := rmoc.OverrideFile(after, "testdata", "override.txt"); err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var dist bytes.Buffer
	io.Copy(&dist, f)

	if dist.String() != ts {
		t.Errorf("File contents must be after but %s", dist.String())
	}

}
