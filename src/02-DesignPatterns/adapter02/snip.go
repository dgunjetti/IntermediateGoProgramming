// +build ignore
//START1 OMIT
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

//END1 OMIT

//START2 OMIT
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

//END2 OMIT