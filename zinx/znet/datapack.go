package znet

import (
	"DAY03/zinx/utils"
	"DAY03/zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包头长度的方法
func (d *DataPack) GetHeadLen() uint32 {
	//Datalen uint32 + ID uint32
	return 8
}

//封包方法
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuf := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuf.Bytes(), nil
}

//拆包方法
func (d *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	dataBuf := bytes.NewReader(binaryData)

	msg := &Message{}

	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("packag too big")
	}

	return msg, nil
}
