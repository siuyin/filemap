package filemap

import "fmt"

func Example() {
	fmap, errs := Map("..")
	fmt.Println(fmap["FILEMAP"])
	fmt.Println(fmap["FILEMAP_TEST"])
	fmt.Println(errs)
	// Output:
	// [../scm-filemap/filemap.go]
	// [../scm-filemap/filemap_test.go]
	// []
}
