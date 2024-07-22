package main

import (
	"fmt"
	"net/http"
)

type Foohandler struct{}

func (f *Foohandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "fooooo")
}

func main() {
	mux := http.NewServeMux()

	//handle함수를 직접 등록할때
	//http 패키지에서, Handfunc 함수 >> "문자열" 값을 인덱스로 지정하고 handler 함수를 파라미터로 가짐,
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {

			name = "world"
		}
		fmt.Fprintf(w, "hello  %s", name)
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello bar")
	})
	//instance로 해서 ServeHTTP()를 직접 구현하여. 사용하는 형태.
	mux.Handle("/foo", &Foohandler{})
	//TCP 연결 유지.아이피주소에 포트 또는 주소,연결

	//여기에 Mux인스턴스를 넘겨서, 라우터 를 지정 mux가 들어올때!
	http.ListenAndServe("", mux)
}
