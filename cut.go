package vorbispak

import (
	"fmt"

	"github.com/yassi-github/vorbispak/cut"
)

func Cut(out, pname, ifile, pak_file string) {
	fmt.Println("Cut:", out, pname, ifile, pak_file)

	cut.Something()
}
