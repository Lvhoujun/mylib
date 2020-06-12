package mylib

import (
	//"bytes"
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"net/http"
	"strconv"
)

// ====================================================================
// Api functions
// ====================================================================

//http post上来的json对象严格定义
type UserInfo struct {
	Id   int64  `json:id`
	Name string `json:name`
	Id2  int64  `json:id`
}

// ====================================================================
// Api functions
// ====================================================================

func HttpTest() {
	StartHttpServer()
}

// ====================================================================
// Internal functions
// ====================================================================

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
	user := new(UserInfo)
	err := ParseRequest(user, res, req)
	if err == nil {
		DEBUG("user=%v", user)
		bytes, err := json.Marshal(user)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		DEBUG("user string=%s", bytes)
		//fmt.Fprintf(res, req.Method+req.URL.Path[1:])
		fmt.Fprintf(res, string(bytes))
	}
}

func ParseRequest(user *UserInfo, res http.ResponseWriter, req *http.Request) error {
	Method := req.Method
	switch Method {
	case "POST":
		ParsePost(user, res, req)
	case "GET":
		ParseGet(user, res, req)
	}
	return nil
}

// @doc post方法解析，严格按照预定义结构解析
func ParsePost(user *UserInfo, res http.ResponseWriter, req *http.Request) error {
	defer func() {
		req.Body.Close()
	}()
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	err := d.Decode(user)
	DEBUG("user=%v", user)
	if err != nil {
		ERROR("err=%v", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

// @doc get方法解析
func ParseGet(user *UserInfo, res http.ResponseWriter, req *http.Request) error {
	values := req.URL.Query()

	ID, _ := strconv.ParseInt(values.Get("id"), 10, 64)
	user.Id = ID
	user.Name = values.Get("name")
	ID, _ = strconv.ParseInt(values.Get("id2"), 10, 64)
	user.Id2 = ID	
	return nil
}
