package mylib

import (
	"errors"
)

/*
* @doc 从第i个位置分割字符串成两部分
* @retrun 失败error为错误信息，成功error :=nil
*/
func Split(s string,i int) (string,string,error){
	if i<0 {
		return "","",errors.New("index is negative number")
	}
	if i>=len(s){
		return s,"",nil
	}
	return s[:i],s[i:],nil
}

func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}