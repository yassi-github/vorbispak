package cut

import ()

func FindXingBofList(pak_file []byte) (bof_xing_symbol_idx_list []uint32, err error) {
	first_xing_symbol := []byte("Xing")
	first_xing_symbol_idx_list, err := findSymbolIdxAll(pak_file, first_xing_symbol)
	if err != nil {
		return nil, err
	}

	return first_xing_symbol_idx_list, nil
}
