package mylib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ====================================================================
// Api functions
// ====================================================================

func HttpTest() {
	StartHttpServer()
}

// ====================================================================
// Internal functions
// ====================================================================

type UserInfo struct {
	Id   int64  `json:id`
	Name string `json:name`
}

func StartHttpServer() {
	http.HandleFunc("/", HandleFuncRoot)
	addr := ":5000"
	DEBUG("开始启动http:[%v]", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		ERROR("http启动失败:%v", err)
		panic(err)
	}
}

func HandleFuncRoot(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	DEBUG("Form=%v", req.Form)
	fmt.Fprintf(res, req.Method+req.URL.Path[1:])
	// for k, _ := range req.Form {
	// 	DEBUG("key=%v\n", k)

	// }
	user := new(UserInfo)
	err := ParseReqJson(user, res, req)
	if err == nil {
		DEBUG("user=%v", user)
	}
}

func ParseReqJson(user *UserInfo, res http.ResponseWriter, req *http.Request) error {
	defer func() {
		req.Body.Close()
	}()

	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	err := d.Decode(user)
	if err != nil {
		ERROR("err=%v",err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
