package cut

import (
	"bytes"
	"fmt"
)

func FindBofList(pak_file []byte) (bof_symbol_idx_list []uint32, extension string, err error) {
	if first_vorbis_symbol_idx_list, err := FindOggsBofList(pak_file); err == nil {
		return first_vorbis_symbol_idx_list, "ogg", nil
	}
	if first_xing_symbol_idx_list, err := FindXingBofList(pak_file); err == nil {
		return first_xing_symbol_idx_list, "mp3",  nil
	}
	
	fmt.Println("Failed to Find Any Music")
	return nil, "", fmt.Errorf("nosupported fmt")
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

	if len(indexes) == 0 {
		return nil, fmt.Errorf("not found")
	}

	return indexes, nil
}

func findLastSymbol(search_data []byte, devide_index_list []uint32, symbol []byte) ([]uint32, error) {
	indexes := make([]uint32, 0)

	for _, devide_index := range devide_index_list {
		search_data_devided_from_bof := search_data[:devide_index]
		found_idx := bytes.LastIndex(search_data_devided_from_bof, symbol)
		if found_idx > -1 {
			indexes = append(indexes, uint32(found_idx))
		} else {
			return nil, fmt.Errorf("not found: index list is wrong")
		}
	}

	if len(indexes) == 0 {
		return nil, fmt.Errorf("not found")
	}

	return indexes, nil
}
