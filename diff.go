package diffr

import (
	"fmt"
	"os"
	"time"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
)

type FileDiff struct {
	OriginName string
	OriginTime time.Time
	NewName string
	NewTime time.Time
	Header []string
	Diffs []diffmatchpatch.Diff
}

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

	originInfo, _ := os.Stat(origin)
	newInfo, _    := os.Stat(new)
	fileDiff := &FileDiff{
		OriginName: originInfo.Name(),
		OriginTime: originInfo.ModTime(),
		NewName:    newInfo.Name(),
		NewTime:    newInfo.ModTime(),
	}
	fmt.Println(fileDiff)

	dmp := diffmatchpatch.New()
	chars1, chars2, lineArray := dmp.DiffLinesToChars(string(o), string(n))
	diffs := dmp.DiffMain(chars1, chars2, false)
	result := dmp.DiffCharsToLines(diffs, lineArray)
	fmt.Println(result)
	return ExitCodeOK
}
