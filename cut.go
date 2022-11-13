package vorbispak

import (
	"fmt"
	"os"

	"github.com/yassi-github/vorbispak/cut"
)

func Cut(out, pname, ifile, pak_file_path string) (exitcode int, err error) {
	pak_file, err := os.ReadFile(pak_file_path)
	if err != nil {
		return 1, err
	}

	bof_oggs_symbol_idx_list, err := cut.FindOggsBofList(pak_file)
	if err != nil {
		return 1, err
	}

	path_body_map, err := cut.CutFile2Map(bof_oggs_symbol_idx_list, pname, out, pak_file)
	if err != nil {
		return 1, err
	}

	for path, body_if := range path_body_map {
		body, ok := body_if.([]byte)
		if !ok {
			return 1, fmt.Errorf("unexpected error: data `%v` not byte but %T", body, body)
		}
		err := os.WriteFile(path, body, os.ModePerm)
		if err != nil {
			return 1, err
		}
	}

	return 0, nil
}
