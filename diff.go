package diffr

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
)

func diffMain(origin, new string) int {
	o, err := ioutil.ReadFile(origin)
	if err != nil {
		fmt.Printf("diffr: %v: No such file or directory\n", origin)
		return ExitCodeError
	}
	n, err := ioutil.ReadFile(new)
	if err != nil {
		fmt.Printf("diffr: %v: No such file or directory\n", new)
		return ExitCodeError
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(o), string(n), true)
	fmt.Println(dmp.DiffPrettyText(diffs))
	return ExitCodeOK
}
