package cut

import (
	"bytes"
	"fmt"
)

func FindOggsBofList(pak_file []byte) (bof_oggs_symbol_idx_list []uint32, err error) {
	bof_oggs_symbol_idx_list = make([]uint32, 0)

	first_vorbis_symbol := append([]byte{01}, []byte("vorbis")...)
	first_vorbis_symbol_idx_list, err := findSymbolIdxAll(pak_file, first_vorbis_symbol)
	if err != nil {
		return nil, err
	}

	bof_ogg_index_list, err := findLastOggS(pak_file, first_vorbis_symbol_idx_list)
	if err != nil {
		return nil, err
	}

	return bof_ogg_index_list, nil
}

func findSymbolIdxAll(pak_file []byte, symbol []byte) (indexes []uint32, err error) {
	indexes = make([]uint32, 0)

	found_idx := bytes.Index(pak_file, symbol)
	for found_idx_from_BOF := 0; found_idx > -1; {
		indexes = append(indexes, uint32(found_idx_from_BOF+found_idx))
		// cut till found idx
		pak_file = pak_file[found_idx+1:]
		// save actual index of new pak_file
		found_idx_from_BOF += (found_idx + 1)
		// re search into sword-offed pak_file
		found_idx = bytes.Index(pak_file, symbol)
	}

	if indexes == nil {
		return nil, fmt.Errorf("not found")
	}

	return indexes, nil
}

func findLastOggS(search_data []byte, devide_index_list []uint32) ([]uint32, error) {
	indexes := make([]uint32, 0)
	search_bytes := []byte("OggS")

	for _, devide_index := range devide_index_list {
		search_data_devided_from_bof := search_data[:devide_index]
		found_idx := bytes.LastIndex(search_data_devided_from_bof, search_bytes)
		if found_idx > -1 {
			indexes = append(indexes, uint32(found_idx))
		} else {
			return nil, fmt.Errorf("not found: index list is wrong")
		}
	}

	return indexes, nil
}
