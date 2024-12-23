package app

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func findGoModFiles(root string, callback func(string)) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "go.mod" {
			callback(path)
		}
		return nil
	})
}


func upgradePackageVersion(content, pkg, newVersion string) string {
    re := regexp.MustCompile(`\b` + regexp.QuoteMeta(pkg) + ` v(\d+\.\d+\.\d+)\b`)
    content = re.ReplaceAllStringFunc(content, func(match string) string {
        version := re.FindStringSubmatch(match)[1]
        if version < newVersion {
            return pkg + " v" + newVersion
        }
        return match
    })
    return content
}

func upgradeGoMod(path string) {
	// path = "./go.mod"
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if !strings.Contains(string(data), "golang.org/x/net") {
		return
	}

	content := string(data)
	pkgs := map[string]string{
		"golang.org/x/net": "0.33.0",
	}

	println(content)
	for pkg, version := range pkgs {
		content = upgradePackageVersion(content, pkg, version)
	}
	println(content)
	os.Exit(0)
	// err = os.WriteFile(path, []byte(content), 0644)
	// if err != nil {
	// 	panic(err)
	// }
}

func TestUpgrade(t *testing.T) {
	dir := "."
	err := findGoModFiles(dir, upgradeGoMod)
	if err != nil {
		t.Fatalf("Error finding go.mod files: %v", err)
	}
}
