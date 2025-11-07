package middleware

import "net/http"

// CustomResponseWriter overloads http.ResponseWriter in order to expose status and length
type CustomResponseWriter struct {
	http.ResponseWriter
	status int
	length int
}

// WriteHeader overloads http.ResponseWriter.WriteHeader() in order to expose the response status
func (w *CustomResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// Write overloads http.ResponseWriter.Write() in order to expose the response length
func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}

	n, err := w.ResponseWriter.Write(b)
	w.length += n

	return n, err
}
