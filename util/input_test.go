package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInput(t *testing.T) {
	f, err := os.Create("testIO")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove("testIO")
	}()
	data := []byte("test\ninput")
	_, err = f.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	empty, err := IsFileEmpty("testIO")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, false, empty)

	_, err = f.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	str, err := ReadStrings(f)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, []string{"test", "input"}, str)

	_, err = f.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	s, err := ReadString(f)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "testinput", s)
}
