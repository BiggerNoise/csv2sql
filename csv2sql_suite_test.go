package csv2sql_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	CurrentDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	FilesDir      = filepath.Join(CurrentDir, "../test_files")
)

func TestCsv2sql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Csv2sql Suite")
}
