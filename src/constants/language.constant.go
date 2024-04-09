package constants

type LanguageMap map[string]map[string]string

const (
	INTERNAL_SERVER_ERROR string = "internal_server_error"
	URL_NOT_FOUND         string = "url_not_found"
	REQUEST_SUCCESS       string = "request_success"
	DATA_NOT_FOUND        string = "data_not_found"
	WRONG_MAIL_PASS       string = "wrong_mail_pass"
	UNVERIFIED_MAIL       string = "unverified_mail"
	LOGIN_SUCCESS         string = "login_success"
)

var Languages = LanguageMap{
	"en": {
		INTERNAL_SERVER_ERROR: "Internal server error.",
		URL_NOT_FOUND:         "URL not found.",
		REQUEST_SUCCESS:       "Request successful.",
		DATA_NOT_FOUND:        "Data not found.",
		WRONG_MAIL_PASS:       "Wrong email or password.",
		UNVERIFIED_MAIL:       "The email has not been verified yet.",
		LOGIN_SUCCESS:         "Login success.",
	},
	"id": {
		INTERNAL_SERVER_ERROR: "Terjadi kesalahan sistem.",
		URL_NOT_FOUND:         "URL tidak ditemukan.",
		REQUEST_SUCCESS:       "Permintaan berhasil.",
		DATA_NOT_FOUND:        "Data tidak ditemukan.",
		WRONG_MAIL_PASS:       "Email atau kata sandi salah.",
		UNVERIFIED_MAIL:       "Email belum terverifikasi.",
		LOGIN_SUCCESS:         "Berhasil masuk.",
	},
}
