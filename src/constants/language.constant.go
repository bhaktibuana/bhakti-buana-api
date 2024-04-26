package constants

import (
	"bhakti-buana-api/src/models"
	"fmt"
)

type LanguageMap map[string]map[string]string

const (
	INTERNAL_SERVER_ERROR string = "internal_server_error"
	URL_NOT_FOUND         string = "url_not_found"
	REQUEST_SUCCESS       string = "request_success"
	DATA_NOT_FOUND        string = "data_not_found"
	WRONG_MAIL_PASS       string = "wrong_mail_pass"
	UNVERIFIED_MAIL       string = "unverified_mail"
	LOGIN_SUCCESS         string = "login_success"
	INVALID_USER          string = "invalid_user"
	FILE_LIMIT_5MB        string = "file_limit_5mb"
	FILE_ACCEPT_PDF       string = "file_accept_pdf"
	INVALID_SORT          string = "invalid_sort"
	INVALID_RESUME_STATUS string = "invalid_resume_status"
	ID_NOT_FOUND          string = "id_not_found"
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
		INVALID_USER:          "Invalid user.",
		FILE_LIMIT_5MB:        "File too large (Max 5 MB).",
		FILE_ACCEPT_PDF:       "File type must be .pdf.",
		INVALID_SORT:          "Sort must be \"acs\" or \"desc\".",
		INVALID_RESUME_STATUS: fmt.Sprintf("Resume status must be \"%s\" or \"%s\".", models.RESUME_STATUS_ACTIVE, models.RESUME_STATUS_INACTIVE),
		ID_NOT_FOUND:          "ID not found.",
	},
	"id": {
		INTERNAL_SERVER_ERROR: "Terjadi kesalahan sistem.",
		URL_NOT_FOUND:         "URL tidak ditemukan.",
		REQUEST_SUCCESS:       "Permintaan berhasil.",
		DATA_NOT_FOUND:        "Data tidak ditemukan.",
		WRONG_MAIL_PASS:       "Email atau kata sandi salah.",
		UNVERIFIED_MAIL:       "Email belum terverifikasi.",
		LOGIN_SUCCESS:         "Berhasil masuk.",
		INVALID_USER:          "Pengguna tidak valid.",
		FILE_LIMIT_5MB:        "File terlalu besar (Max 5 MB).",
		FILE_ACCEPT_PDF:       "Tipe file harus .pdf.",
		INVALID_SORT:          "Sort harus \"acs\" atau \"desc\".",
		INVALID_RESUME_STATUS: fmt.Sprintf("Status Resume harus \"%s\" atau \"%s\".", models.RESUME_STATUS_ACTIVE, models.RESUME_STATUS_INACTIVE),
		ID_NOT_FOUND:          "ID tidak ditemukan.",
	},
}
