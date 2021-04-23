package utils

import (
	"DAY03/zinx/ziface"
	"encoding/json"
	"io/ioutil"
)

/*
	存储一切有关zinx框架的全局参数，供其他模块使用

*/

type GlobalObj struct {
	/*
		server
	*/
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	/*
		zinx
	*/
	Version           string
	MaxConn           int
	MaxPackageSize    uint32
	WorkerPoolSize    uint32
	MaxWorkerTaskSize uint32
}

var GlobalObject *GlobalObj

/*
	提供一个init方法， 初始化当前的GlobalObject
*/

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}

}

func init() {
	GlobalObject = &GlobalObj{
		Name:              "ZinxServerApp",
		Version:           "V0.4",
		TcpPort:           8999,
		Host:              "0.0.0.0",
		MaxConn:           1000,
		MaxPackageSize:    4096,
		WorkerPoolSize:    10,
		MaxWorkerTaskSize: 1024,
	}

	//尝试从conf/zinx.json去加载一些用户自定义的参数
	GlobalObject.Reload()
}
