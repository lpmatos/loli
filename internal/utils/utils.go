package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/lpmatos/loli/internal/constants"
	"github.com/sirupsen/logrus"
)

type markdownRenderOpts []glamour.TermRendererOption

// RenderMarkdown function that return a pretty markdown in output.
func RenderMarkdown(text string) (string, error) {
	opts := markdownRenderOpts{
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(75),
		glamour.WithEmoji(),
	}
	return renderMarkdown(text, opts)
}

func renderMarkdown(text string, opts markdownRenderOpts) (string, error) {
	text = strings.ReplaceAll(text, "\r\n", "\n")

	tr, err := glamour.NewTermRenderer(opts...)
	if err != nil {
		return "", err
	}

	out, err := tr.Render(text)
	if err != nil {
		return "", err
	}

	return out, nil
}

// IsEmpty function - check if a string is empty.
func IsEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// IsDirExists function - check fi a directory exist in te system.
func IsDirExists(path string) bool {
	result, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return result.IsDir()
}

// IsFileExists function - check fi a file exist in te system.
func IsFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// MakeDirIfNotExist create a new directory if they not exist.
func MakeDirIfNotExist(dir string) {
	fullDir, _ := filepath.Abs(filepath.Dir(dir))
	mode := os.FileMode(0775)
	if !IsDirExists(fullDir) {
		err := os.MkdirAll(fullDir, mode)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

// CreateLogFile function
func CreateLogFile(logdir, logfile string) string {
	pid := strconv.Itoa(os.Getpid())
	return filepath.Join(logdir,
		logfile+
			".pid"+
			pid+
			"."+
			time.Now().Format(constants.DefaultTimestampFormat)+
			".log",
	)
}
