package rmoc

import (
	"bytes"
	"testing"
)

func TestCreateFileWithAbort(t *testing.T) {
	buf := bytes.NewBufferString("test")
	err:= CreateFileWithAbort(buf, "testdata", "dummy")

	if ok, _ := IsFileAlreadyExists(err); !ok {
		t.Errorf("error must be FileAlreadyExists byt %+v", err)
	}
}
