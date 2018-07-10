// Package filemap walks a filesystem and returns
// a map[string][]string. Eg. m["ABC"] -> ["abc.go","a/b/Abc.txt","c/aBc"]
// where c/aBc is a file and not a directory.
package filemap

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Map walks path recursively and returns a map[string][]string
// and a slice of errors.
func Map(dir string) (map[string][]string, []error) {
	fmap := make(map[string][]string)
	errs := []error{}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			errs = append(errs, fmt.Errorf("walk encountered error for %v: %v", path, err))
		}
		if info.IsDir() {
			return nil
		}
		ubn := strings.ToUpper(filepath.Base(path)) // upper-cased basename, abc.go -> ABC.GO
		ubn = strings.Split(ubn, ".")[0]            // ABC.GO -> ABC
		fmap[ubn] = append(fmap[ubn], path)
		return nil
	})
	return fmap, errs
}
