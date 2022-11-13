package cut

import (
	"fmt"
)

func FindOggsBofList(pak_file []byte) (bof_oggs_symbol_idx_list []uint32, err error) {
	oggs_symbol := []byte("OggS")
	vorbis_symbol := []byte("vorbis")

	bof_oggs_symbol_idx_list = make([]uint32, 0)
	// read loop till eof
		var reading_idx uint32

		first_vorbis_symbol := byte(01) + vorbis_symbol
		first_vorbis_symbol_idx, err := findSymbolIdx(pak_file, reading_idx, first_vorbis_symbol)
		if err != nil {
			// first vorbis not found
			// exit success if already found
			if bof_oggs_symbol_idx_list > 0 {
				return 0, nil
			}
			return 1, fmt.Errorf("ogg vorbis not found")
		}

		bof_oggs_symbol_idx, err := findSymbolIdxBackforward(pak_file, first_vorbis_symbol_idx, oggs_symbol)

}

func findSymbolIdx(pak_file []byte, begin_idx uint32, symbol []byte) (index uint32, err error) {
	// 
}
func findSymbolIdxBackforward(pak_file []byte, begin_idx uint32, symbol []byte) (index uint32, err error) {
	// 
}
