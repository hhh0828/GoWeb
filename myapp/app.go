package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Foohandler struct{}

func (f *Foohandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//빈 유저 객체
	user := new(User)
	//Json value 를 go object 객체로 변환// err가 비어있으면~
	//	err := json.NewDecoder(r.Body).Decode(user)
	//err1 := json.Unmarshal([]byte(r.Body), user)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//		fmt.Fprint(w, err)
	//		return
	//	}
	json.NewDecoder(r.Body).Decode(user)
	user.CreatedAt = time.Now()
	//user를 다시 json으로 변환
	data, _ := json.Marshal(user)
	//err123 := json.Unmarshal([]byte(data), &user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	//타입변환 해줘야함, Json타입이기때문에 읽을수있는 ...
	fmt.Fprint(w, string(data))
	fmt.Println()
}

type User struct {
	FistName  string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt time.Time
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	//handle함수를 직접 등록할때
	//http 패키지에서, Handfunc 함수 >> "문자열" 값을 인덱스로 지정하고 handler 함수를 파라미터로 가짐,
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Json-
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
	http.ListenAndServe("", mux) ///
	return mux
}
