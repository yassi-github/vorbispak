package cut

import (
	"fmt"
	"path/filepath"

	"golang.org/x/exp/constraints"
)

func CutFile2Map(bof_oggs_symbol_idx_list []uint32, pname, out string, pak_file []byte) (path_body_map map[string]interface{}, err error) {
	path_body_map = make(map[string]interface{}, 0)

	// remove non-ogg header
	header_len := bof_oggs_symbol_idx_list[0]
	pak_file = pak_file[header_len:]

	// bof_ogg_index_list_header_aligned := make([]uint32, len(bof_oggs_symbol_idx_list))
	// for idx, bof_ogg_index := range bof_oggs_symbol_idx_list {
	// 	bof_ogg_index_list_header_aligned[idx] = bof_ogg_index - header_len
	// }
	bof_ogg_index_list_header_aligned := fnmap_number_list[uint32](bof_oggs_symbol_idx_list, func(e uint32) uint32 { return e - header_len })

	// remove first index because it is fixed to zero
	bof_ogg_index_list_header_aligned = bof_ogg_index_list_header_aligned[1:]

	old_ogg_index := uint32(0)
	for file_idx, ogg_index := range bof_ogg_index_list_header_aligned {
		// ogg_index is actual index but `pak_file` is relative index data
		single_ogg_data := pak_file[:ogg_index - old_ogg_index]
		// os.WriteFile(fmt.Sprintf("./fixtures/out_%v.ogg", file_idx), single_ogg_data, os.ModePerm)
		path_body_map[filepath.Join(out, fmt.Sprintf("%v%v.ogg", pname, file_idx))] = single_ogg_data
		// cut off already found data
		pak_file = pak_file[ogg_index - old_ogg_index:]
		old_ogg_index = ogg_index
	}
	// till EOF
	path_body_map[filepath.Join(out, fmt.Sprintf("%v%v.ogg", pname, len(bof_ogg_index_list_header_aligned)))] = pak_file[:]

	return path_body_map, nil
}

type Number interface {
	constraints.Integer | constraints.Float
}

func fnmap_number_list[T Number](list []T, fn func(T) T) []T {
	result := make([]T, len(list))
	for idx, elem := range list {
		result[idx] = fn(elem)
	}
	return result
}
