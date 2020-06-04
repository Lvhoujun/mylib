package mylib
//package main

import "fmt"

// ====================================================================
// struct defination
// ====================================================================

//接口定义
type Shaper interface{
	Area() float32
}

type isquare struct{
	side float32
}


type irectange struct {
	width float32
	length float32
}

// ====================================================================
// API functions
// ====================================================================

type List []int

func (l List) Len() int {
    return len(l)
}

func (l *List) Append(val int) {
    *l = append(*l, val)
}

type Appender interface {
    Append(int)
}

func CountInto(a Appender, start, end int) {
    for i := start; i <= end; i++ {
        a.Append(i)
    }
}

type Lener interface {
    Len() int
}


func (sql *isquare) Area() float32 {
	return sql.side*sql.side
}

func (rec *irectange) Area() float32 {
	return rec.width*rec.length
}

func LongEnough(l Lener) bool {
    return l.Len()*10 > 42
}

func Test(){
	// sql := new(isquare)
	// sql.side=10
	// var shapterI Shaper
	// shapterI=sql
	
	// Classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)

	// Classifier(shapterI)

	// fmt.Printf("shapterI=%v,Area=%f\n",shapterI,shapterI.Area())

	// rec := new(irectange)
	// rec.width = 10
	// rec.length = 20
	// shapterI=rec

	// Classifier(shapterI)

	// fmt.Printf("shapterI=%v,Area=%f\n",shapterI,shapterI.Area())

	// A bare value
    //var lst List
    // compiler error:
    // cannot use lst (type List) as type Appender in argument to CountInto:
    //       List does not implement Appender (Append method has pointer receiver)
    //CountInto(lst, 1, 10)
    // if LongEnough(lst) { // VALID:Identical receiver type
    //     fmt.Printf("- lst is long enough\n")
    // }

    // A pointer value
    plst := new(List)
    CountInto(plst, 1, 10) //VALID:Identical receiver type
    fmt.Printf("plst=%v\n",plst)
    if LongEnough(plst) {
        // VALID: a *List can be dereferenced for the receiver
        fmt.Printf("- plst is long enough\n")
    }


}

func Classifier(items ...interface{}){
	for i,x := range items{
		switch x.(type) {
			case bool:
				fmt.Printf("Param #%d is a bool\n",i)
			case float64:
				fmt.Printf("Param #%d is a float64\n",i)
			case int,int64:
				fmt.Printf("Param #%d is a int\n",i)
			case nil:
				fmt.Printf("Param #%d is a nil\n",i)
			case string:
				fmt.Printf("Param #%d is a string\n",i)
			default:
				fmt.Printf("Param #%d is unkown x=%v\n",i,x)
		}
	}
}