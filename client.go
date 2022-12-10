package main

import (
	"fmt"
	"net"
)

// UDP 客户端
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("UDP connect error，err: ", err)
		return
	}
	defer socket.Close()
	sendData := make([]byte, 1024)
	for {
		_, err = socket.Write(sendData)
		if err != nil {
			fmt.Println("send message error，err: ", err)
			return
		} else {
			fmt.Println("ok")
		}
	}

}
