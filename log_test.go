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
	l := DefaultLog()
	assert.Equal(t, l.level, LvlError)
}

func TestWithLevel(t *testing.T) {
	l := DefaultLog()
	ll := l.WithLevel(LvlDebug)
	assert.Equal(t, ll.level, LvlDebug)
}

func TestWithWriter(t *testing.T) {
	l := DefaultLog()
	var w bytes.Buffer

	ll := l.WithWriter(&w)
	message := "Hello, Log!"
	n, err := fmt.Fprintf(ll.writer, message)

	assert.Equal(t, len(message), n)
	require.NoError(t, err)
}

func TestFatal(t *testing.T) {
	if os.Getenv("SHOULD_CRASH") == "1" {
		var w bytes.Buffer
		l := NewLog(LvlDebug, &w)
		l.Fatal("that's it... we're done here")
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
	l := NewLog(LvlDebug, w)
	message := "I love this package %d!"
	formatted := fmt.Sprintf(message, 100)
	withHeader := header(time.Now(), LvlError) + " " + formatted
	l.Error(message, 100)

	assert.Equal(t, len(withHeader)+1, w.Len())

	b, err := ioutil.ReadAll(w)
	require.NoError(t, err)
	printed := string(b)

	t.Logf("printed: %s", printed)

	assert.True(t, strings.Contains(printed, "ERROR"))
	assert.True(t, strings.HasSuffix(printed, formatted+"\n"))
}
