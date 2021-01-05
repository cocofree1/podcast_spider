package common

var IndexNameMap map[int]string
const (
	ORIGIN_ALBUM_DATA = iota
	ORIGIN_PROGRAM_DATA
	ALBUM_DATA
	PROGRAM_DATA
)

func init(){
	IndexNameMap = map[int]string{
		ORIGIN_ALBUM_DATA:     "origin_album_data",
		ORIGIN_PROGRAM_DATA:   "origin_program_data",
		ALBUM_DATA:            "album_data",
		PROGRAM_DATA:          "program_data",
	}
}