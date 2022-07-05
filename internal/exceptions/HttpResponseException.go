package exceptions

type HttpResponseException struct {
}

func (respEx HttpResponseException) Error() string {
	return "http Response error"
}
