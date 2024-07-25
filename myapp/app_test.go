package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Json 복습
/*
type Users struct {
	firstname string
	lastname  string
	email     string
	age       int
	Createdat time.Time
}

func (u *Users) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user := new(Users)

	//Create Decoder and make the r.Body Json to readable things for Go object.
	json.NewDecoder(r.Body).Decode(user)
	//after create a object that retrived from Json.

	//add field that are empty or not created from input value.
	user.Createdat = time.Now()

	//change the object to Json value for sending the data via http. with marshal feature.
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user) // ignore err, >> if there's non issue it will retrive the empty data with NIL.
	fmt.Fprint(w, string(data))
	fmt.Println("ok")

}
*/
func Indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")

}

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)
	//test용 Recorder and Rquest
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	Indexhandler(res, req)
	/*
		Indexhandler(res, req)
		if res.Code != http.StatusOK {
			t.Fatal("failed to get a page", res.Code)
		}
	*/
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))

	muxtest1 := http.NewServeMux()
	muxtest1.HandleFunc("/", Indexhandler)

	http.ListenAndServe("", muxtest1)
	//test2
	//
}

func TestBarPathHandler(t1 *testing.T) {
	assert := assert.New(t1)
	//test용 Recorder and Rquest
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	//Barhandler(res, req)
	/*
		Indexhandler(res, req)
		if res.Code != http.StatusOK {
			t.Fatal("failed to get a page", res.Code)
		}
	*/
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))

	muxtest1 := http.NewServeMux()
	muxtest1.HandleFunc("/", Indexhandler)

	http.ListenAndServe("", muxtest1)
	//test2
	//
}

func TestFooPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

}
func TestFooPathWithJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name":"hyunho", "last_name":"hong", "email":"hyunho.hong@naver.com", "age":32}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	user := new(User)
	json.NewDecoder(res.Body).Decode(user)

	assert.Equal("hyunho", user.FistName)
	assert.Equal("hong", user.Lastname)

	/*
		fmt.Println("test")
		user.CreatedAt = time.Now()

		Data, _ := json.Marshal(user)

		fmt.Fprint(res, Data)
	*/

}

//Data의 주체 - response(출력) 와 request(입력) 인스턴스화 되어있음
//MUX MULTFLEXER를 통해 루트를 등록 또는 Request를 ㅂ내고 Response를 받을 수 있음.
//assert.Equal를 통하여 기대된 값과, 실제 값을 비교하여 Status 를 확인 할 수 있음. with goconvey
//Json타입의 입력된 값은 Decoder 를 통해 인스턴스화 하여, Decode를 통하여 Go object로 변환해야 하며, > Unmarshal
//Go object를 통하여 입력된 값을 처리해주고 다시 Json 타입으로 Marshal 하여 해당 데이터를 보낼 수 있음.

//위 모든 과정은 NewHTTPHandler() http.handler 를 반환하는 함수로 MUX를 구성할 수 있음.
// ServeMux()는 ServeHTTP를 구현하고 있음. Handler인터페이스를 통하여 ServeHTTP 사용 가능.
// mux := http.NewServeMux() 새로운 서브 먹스를 구현하여 인덱스 및 각 주소루트를 작성 가능함.
// 메인에서 Listen and Serve로 포트 값과 서브먹스핸드러 리스트함수를 등록가능함, NewHTTPhandler()는 Mux를 반환 함.
/* func main() {
	http.ListenAndServe("", myapp.NewHttpHandler())
}
*/
//먹스를 통해 등록하고 먹스에 등록되어있는 경로 값을 토대로 함수 핸드러를 반환함.

// 메인 핸드러 받이 대기
// 핸드러 리스트
// 핸드러 구현 부
