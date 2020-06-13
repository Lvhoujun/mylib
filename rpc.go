package mylib

/*
* rpc
*/

import (
    "net/http"    
    "net"
    "net/rpc"
    "time"    
)

// ====================================================================
// Struct functions
// ====================================================================

type Args struct{
    M,N int
}

func (t *Args) Multiply(args *Args, retVal *int) error{
    *retVal = args.M * args.N
    return nil
}

func (t *Args) Add(args *Args, retVal *int) error{
    *retVal = args.M + args.N
    return nil
}

// ====================================================================
// Api functions
// ====================================================================

func RPCEntry(){
	RpcSerEntry()
}

func RpcSerEntry(){
    calc := new(Args)
    rpc.Register(calc)
    rpc.HandleHTTP()
    listener,err := net.Listen("tcp",":5000")
    if err != nil {
        ERROR("start rpc server fail,err:=%v",err)
        return 
    }
    go http.Serve(listener,nil)
    for i := 0; i < 10; i++ {
    	go Client(i,i*i)	
    }
    time.Sleep(30*time.Second)
}

func Client(m,n int){
    client,err := rpc.DialHTTP("tcp", "localhost:5000")
    if err != nil {
    	ERROR("connect server failed, err:%v",err)
    	return 
    }
    args := &Args{m,n}
    var reply int
    err = client.Call("Args.Multiply", args, &reply)
    if err != nil {
    	ERROR("rpc call failed, err=%v",err)
    	return 
    }
    DEBUG("rpc succ:%d * %d = %d",args.M,args.N,reply)
}

// ====================================================================
// Internal functions
// ====================================================================




