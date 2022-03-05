package utils

// ReturnError Return various errors to JSON
func ReturnError(code int, err error) {
	
	recovered := map[string]interface{}{
		"code": code,
		"err":  err,
	}

	panic(recovered)
}
