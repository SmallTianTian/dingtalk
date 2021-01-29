package taobao

type HttpResponse struct {
	Header map[string]string
	Body   []byte
	Code   int
}
