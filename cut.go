package vorbispak

import (
	"fmt"

	"github.com/yassi-github/vorbispak/cut"
)

func Cut(out, pname, ifile, pak_file_path string) (exitcode int, err error) {
	pak_file, err := readByteFileAll(pak_file_path)
	if err != nil {
		return 1, err
	}

	bof_oggs_symbol_idx_list, err := cut.FindOggsBofList(pak_file)
	if err != nil {
		return 1, err
	}

	path_body_map, err := cut.CutFile2Map(bof_oggs_symbol_idx_list, out, pname, pak_file)
	if err != nil {
		return 1, err
	}

	for path, body := range path_body_map {
		err := writeByteFileAll(path, body)
		if err != nil {
			return 1, err
		}
	}
}
