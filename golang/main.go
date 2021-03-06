package main

import (
	"fmt"
)

//data type
const (
	DataTypeTxt   = 1
	DataTypeJsonp = 2
	DataTypeXml   = 3
)

func main() {
	//方法一载入内存
	findIpLocation()

	//方法二文件查找
	findIpLocation2()
}


func findIpLocation(){
	location := new(IpLocation)
	location.InitFromFile("../ip.dat")
	result := location.Find("255.255.255.255", DataTypeJsonp, "find")
	fmt.Println(result)
}

func findIpLocation2(){
	location2 := new(IpLocation2)
	location2.InitFromFile("../ip.dat")

	fmt.Println(location2.Query("0.0.0.0", DataTypeJsonp, "find"));
	fmt.Println(location2.Query("255.255.255.255", DataTypeJsonp, "find"));
	fmt.Println(location2.Query("152.32.193.0", DataTypeJsonp, "find"));
	fmt.Println(location2.Query("1.0.2.1", DataTypeJsonp, "find"));
	fmt.Println(location2.Query("110.81.112.0", DataTypeJsonp, "find"));
}