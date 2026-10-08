package utils

import "latihan_rest_api/pkg/constant"

func ValidateFile(extension string) bool {
	return constant.ALLOWED_EXTENSIONS[extension]
}
