package httpserv

type From int

const (
	NODATA From = iota
	URL
	JSONBODY
	FORMDATA
)

type httpdata struct {
	DataFrom From
}
