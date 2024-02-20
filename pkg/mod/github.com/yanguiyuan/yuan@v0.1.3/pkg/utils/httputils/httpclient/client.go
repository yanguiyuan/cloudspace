package httpclient

type Response struct {
}
type Request struct {
}
type url struct {
	protocol int8
	addr     string
	port     int16
	path     string
	params   map[string]string
}

func Get(url string) *Response {
	return nil
}
func BaseURL(addr string) *Request {
	return nil
}
func IP() {

}
func SetDefaultBaseUrl(url string) {

}
