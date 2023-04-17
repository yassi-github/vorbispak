package cut

import ()

func FindOggsBofList(pak_file []byte) (bof_oggs_symbol_idx_list []uint32, err error) {
	first_vorbis_symbol := append([]byte{01}, []byte("vorbis")...)
	first_vorbis_symbol_idx_list, err := findSymbolIdxAll(pak_file, first_vorbis_symbol)
	if err != nil {
		return nil, err
	}

	oggs_symbol := []byte("OggS")
	bof_ogg_index_list, err := findLastSymbol(pak_file, first_vorbis_symbol_idx_list, oggs_symbol)
	if err != nil {
		return nil, err
	}

	return bof_ogg_index_list, nil
}
