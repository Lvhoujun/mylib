package mylib

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"bufio"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

/*
* chap19_v1,利用读写锁保护临界区,数据保存在文件
*
 */

// ====================================================================
// Struct functions
// ====================================================================

type URLStore struct {
	urls map[string]string //短url=>长url
	mu   sync.RWMutex      //读写锁
	file *os.File
}

type record struct {
	Key, URL string
}

func NewURLStore(filename string) *URLStore {
	f,err := os.OpenFile(filename,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		ERROR("open file %s error:%v",filename,err)
		os.Exit(1)
	}

	br := bufio.NewReader(f)
	s:= &URLStore{urls: make(map[string]string)}
	s.file=f
	for{
		line,_,err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			ERROR("read file %s error:%v",filename,err)
			os.Exit(1)
		}
		var rec record
		err = json.Unmarshal(line, &rec)
		if err != nil {
			ERROR("Unmarshal %s error:%v",string(line),err)
			os.Exit(1)
		}
		s.doSet(rec.Key,rec.URL)
	}
	return  s
}

// @doc 短url=>长url
func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

// @doc 设置url
func (s *URLStore) Set(url string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := BKDRHash(url)
	return s.doSet(key,url)
}

func (s *URLStore) doSet(key,url string) (string,bool){
	_, present := s.urls[key]
	if present {
		return "", false
	}
	s.urls[key] = url
	return key, true
}

// 写文件
func (s *URLStore) save(key,url string) bool{
	rec := record{key,url}
	b,err := json.Marshal(rec)
	if err != nil {
		ERROR("json Marshal error:%v,rec:%v",err,rec)
		return false
	}
	byteL := make([][]byte, 2)
	byteL[0] = b
	byteL[1] = []byte("\n")
	b = bytes.Join(byteL, []byte(""))
	_,err = s.file.Write(b)
	if err != nil {
		ERROR("write file error:%v,rec:%v",err,rec)
	}
	return true
}

// @doc BKDR hash算法
func BKDRHash(s string) string {
	var seed uint = 131
	var hash uint = 0

	for _, v := range s {
		hash = hash*seed + uint(v)
	}

	hash = hash & 0x7FFFFFFF
	return strconv.FormatUint(uint64(hash), 10)
}

func (s *URLStore) Json() (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, err := json.Marshal(s.urls)
	if err != nil {
		return "", err
	}
	return string(b), err
}

// ====================================================================
// Api functions
// ====================================================================

const  FileName = "./url.json"
var store = NewURLStore(FileName)

func Chap19V1Entry() {
	http.HandleFunc("/", View)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/view", View)
	addr := ":8080"
	DEBUG("启动http服务[%s]", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		ERROR("http服务启动失败:%v", err)
		panic(err)
	}

	defer func() {
		if store.file != nil {
			store.file.Close()
		}
	}()
}

// ====================================================================
// http handle method
// ====================================================================

func Add(w http.ResponseWriter, req *http.Request) {
	url := req.FormValue("url")
	if url == "" {
		fmt.Fprint(w, "please input url")
		return
	}
	key, succ := store.Set(url)
	if !succ {
		fmt.Fprint(w, url+" already exists")
		return
	}
	store.save(key,url)
	fmt.Fprint(w, key)
}

// @doc 查询返回json
func View(w http.ResponseWriter, req *http.Request) {
	s, err := store.Json()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, s)
}
