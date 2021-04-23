package znet

import (
	"DAY03/zinx/ziface"
	"errors"
	"fmt"
	"sync"
)

type ConnManager struct {
	connection map[uint32]ziface.IConnection

	connLock sync.RWMutex
}

func NewConnManager() ziface.IConnManager {
	return &ConnManager{
		connection: make(map[uint32]ziface.IConnection),
	}

}

func (cm *ConnManager) Add(conn ziface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connection[conn.GetConnID()] = conn

	fmt.Println("connID = ", conn.GetConnID(), " add to ConnManager successflly : conn num = ", cm.Len())
}

//删除链接
func (cm *ConnManager) Remove(conn ziface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connection, conn.GetConnID())
	fmt.Println("connID = ", conn.GetConnID(), " remove from ConnManager successflly : conn num = ", cm.Len())

}

func (cm *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn, ok := cm.connection[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not FOUND")
	}
}

func (cm *ConnManager) Len() int {
	return len(cm.connection)
}

func (cm *ConnManager) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for connID, conn := range cm.connection {
		conn.Stop()

		delete(cm.connection, connID)
	}

	fmt.Println("Clear All connection succ! conn num = ", cm.Len())
}
