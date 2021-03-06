package log

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultLog(t *testing.T) {
	l := Default()
	assert.Equal(t, l.level, LevelError)
}

func TestWithLevel(t *testing.T) {
	l := Default()
	ll := l.WithLevel(LevelDebug)
	assert.Equal(t, ll.level, LevelDebug)
}

func TestWithWriter(t *testing.T) {
	l := Default()
	var w bytes.Buffer

	ll := l.WithWriter(&w)
	message := "Hello, Log!"
	n, err := fmt.Fprint(ll.writer, message)

	assert.Equal(t, len(message), n)
	require.NoError(t, err)
}

func TestFatal(t *testing.T) {
	if os.Getenv("SHOULD_CRASH") == "1" {
		var w bytes.Buffer
		l := New(LevelDebug, &w, "", true)
		l.Fatalf("that's it... we're done here")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
	cmd.Env = append(os.Environ(), "SHOULD_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); !ok && !e.Success() {
		t.Fatalf("process ran with err %v, want exit status 1", err)
	}
}

func TestError(t *testing.T) {
	w := new(bytes.Buffer)
	l := New(LevelDebug, w, "", true)
	message := "I love this package %d!"
	formatted := fmt.Sprintf(message, 100)
	withHeader := l.header(time.Now(), LevelError) + " " + formatted
	l.Errorf(message, 100)

	assert.Equal(t, len(withHeader)+1, w.Len())

	b, err := ioutil.ReadAll(w)
	require.NoError(t, err)
	printed := string(b)

	t.Log(printed)

	assert.True(t, strings.Contains(printed, "ERROR"))
	assert.True(t, strings.HasSuffix(printed, formatted+"\n"))
}

func TestPrefix(t *testing.T) {
	w := new(bytes.Buffer)
	prefix := "LOG"
	l := New(LevelDebug, w, prefix, true)
	message := "prefixed message here"
	withHeader := l.header(time.Now(), LevelError) + " " + message
	l.Info(message)

	assert.Equal(t, len(withHeader)+1, w.Len())

	b, err := ioutil.ReadAll(w)
	require.NoError(t, err)
	printed := string(b)

	t.Log(printed)

	assert.True(t, strings.Contains(printed, "INFO"))
	assert.True(t, strings.HasSuffix(printed, message+"\n"))
}

func TestPanic(t *testing.T) {
	if os.Getenv("SHOULD_CRASH") == "1" {
		var w bytes.Buffer
		l := New(LevelDebug, &w, "", true)
		l.Panic("that's it... we're done here")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestPanic")
	cmd.Env = append(os.Environ(), "SHOULD_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); !ok && !e.Success() {
		t.Fatalf("process ran with err %v, want exit status 1", err)
	}
}
