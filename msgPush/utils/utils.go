package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"net"
)

func Writen(conn net.Conn, bytes []byte) error {
	leftToWrite := len(bytes)
	for leftToWrite > 0 {
		n, err := conn.Write(bytes)
		if err != nil {
			return err
		}
		leftToWrite -= n
	}
	return nil
}

// create message json bytes
func CreatMessage(m map[string]interface{}) ( []byte, error ) {
	messageBytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	var bytesLength int32 = int32(len(messageBytes)+4)
	bs := []byte{}
	buf := bytes.NewBuffer(bs)

	binary.Write(buf, binary.LittleEndian, &bytesLength)
	binary.Write(buf, binary.LittleEndian, messageBytes)
	messageBytes = buf.Bytes()
	return messageBytes, nil
}

func DeleteConnFromMap(m map[string]net.Conn, conn net.Conn) {
	for k, v := range m {
		if conn == v {
			delete(m, k)
		}
	}
}
