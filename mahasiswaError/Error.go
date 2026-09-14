package mahasiswaError

import "fmt"

type ErrorMahasiswa struct {
	StatusCode int
	Message    string
}

func (e *ErrorMahasiswa) Error() string {
	return fmt.Sprintf("statusCode: %d | Pesan: %s\n", e.StatusCode, e.Message)
}
