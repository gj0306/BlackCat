package main

import (
	"fmt"
	"strconv"
	"time"
	"./github.com/golang/protobuf/proto"
	"./proto/gs"
)


func Write(key,value string,chs chan []byte) {
	obj := &gs.StrMap{
		Key:key,
		Value:value,
	}
	//编码数据
	data, err := proto.Marshal(obj);
	if err != nil{
		fmt.Println("Write Marshal err :%+v",err)
		return
	}
	//把数据写入chan
	chs<-data
}

func Read(chs  chan []byte){
	data := <-chs
	obj := &gs.StrMap{}
	if err := proto.Unmarshal(data,obj);err != nil{
		fmt.Println("Read Unmarshal err :%+v",err)
		return
	}
	fmt.Println("结果:%+v",obj)
}

func main()  {
	chs := make(chan []byte,5)
	for i:=int64(0);i<10;i++{
		key := "key:" + strconv.FormatInt(i,10)
		value := "value:" + strconv.FormatInt(i,10)
		go Write(key,value,chs)
	}
	time.Sleep(time.Second*1)
	for len(chs)>0{
		go Read(chs)
	}
	time.Sleep(time.Second*1)
	fmt.Println("hello word")

}
