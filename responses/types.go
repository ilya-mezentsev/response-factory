package responses

type Response interface {
	HttpStatus() int
	ApplicationStatus() string
	HasData() bool
	Data() interface{}
}
