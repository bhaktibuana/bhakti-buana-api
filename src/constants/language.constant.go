package constants

type LanguageMap map[string]map[string]string

const (
	INTERNAL_SERVER_ERROR string = "internal_server_error"
	URL_NOT_FOUND         string = "url_not_found"
	REQUEST_SUCCESS       string = "request_success"
	DATA_NOT_FOUND        string = "data_not_found"
)

var Languages = LanguageMap{
	"en": {
		INTERNAL_SERVER_ERROR: "Internal server error.",
		URL_NOT_FOUND:         "URL not found.",
		REQUEST_SUCCESS:       "Request successful.",
		DATA_NOT_FOUND:        "Data not found.",
	},
	"id": {
		INTERNAL_SERVER_ERROR: "Terjadi kesalahan sistem.",
		URL_NOT_FOUND:         "URL tidak ditemukan.",
		REQUEST_SUCCESS:       "Permintaan berhasil.",
		DATA_NOT_FOUND:        "Data tidak ditemukan.",
	},
}
