package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(handlers ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(handlers) - 1; i >= 0; i-- {
			handler := handlers[i]
			next = handler(next)
		}
		return next
	}
}

type ErrorMsg struct {
	Msg any
}

func ErrorMessagePage(w http.ResponseWriter, status int, message string) {
	fmt.Printf("status: %v\n", status)
	w.WriteHeader(status)
	html := `
		<!DOCTYPE html>
		<html>

		<head>
			<link rel="icon" href="/favicon.ico" type="image/x-icon">
			<title>Error</title>
		</head>

		<body>
			<h1>Error</h1>
			<p>%v</p>
		</body>

		</html>
	`
	if status == http.StatusNotFound {
		fmt.Println("not found should fire")
		tmpl, err := template.ParseFiles(IMPORTANT_FILES[http.StatusNotFound])
		if err != nil {
			w.Write([]byte(fmt.Sprintf(html, message)))
			return
		}
		err = tmpl.Execute(w, struct {
			Msg string
		}{Msg: message})
		if err != nil {
			w.Write([]byte(fmt.Sprintf(html, message)))
			return
		}

	} else if status == http.StatusBadRequest {
		// CHECK THIS
		tmpl, err := template.ParseFiles(IMPORTANT_FILES[http.StatusBadRequest])
		if err != nil {
			w.Write([]byte(fmt.Sprintf(html, message)))
			return
		}
		err = tmpl.Execute(w, struct {
			Msg string
		}{Msg: message})
		if err != nil {
			w.Write([]byte(fmt.Sprintf(html, message)))
			return
		}
	} else if status == http.StatusInternalServerError {
		tmpl, err := template.ParseFiles(IMPORTANT_FILES[http.StatusInternalServerError])
		if err != nil {
			w.Write([]byte(fmt.Sprintf(html, message)))
			return
		}
		err = tmpl.Execute(w, struct {
			Msg string
		}{Msg: message})
		if err != nil {
			w.Write([]byte(fmt.Sprintf(html, message)))
			return
		}
	} else if status == http.StatusNotAcceptable {
		tmpl, err := template.ParseFiles(IMPORTANT_FILES[http.StatusNotAcceptable])
		if err != nil {
			w.Write([]byte(fmt.Sprintf(html, message)))
			return
		}
		err = tmpl.Execute(w, struct {
			Msg string
		}{Msg: message})
		if err != nil {
			w.Write([]byte(fmt.Sprintf(html, message)))
			return
		}
	}

}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				errMsg, ok := err.(string)
				if !ok {
					fmt.Println("1")
					ErrorMessagePage(w, http.StatusInternalServerError, "Internal server error")
					return
				}
				fmt.Println("2")
				ErrorMessagePage(w, http.StatusInternalServerError, errMsg)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}

func CheckTemplates(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exists := ensureFiles(IMPORTANT_FILES)
		if !exists {
			panic("missing important template files")
		}
		next.ServeHTTP(w, r)
	})
}
