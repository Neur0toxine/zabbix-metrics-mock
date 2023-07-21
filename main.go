package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	packetHeader     = []byte("ZBXD\x01")
	packetHeaderSize = len(packetHeader)
)

func main() {
	listen := ":10051"
	if envListen := os.Getenv("LISTEN"); envListen != "" {
		listen = envListen
	}
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("listen: %s", err)
	}
	defer listener.Close()
	log.Printf("listening on %s", listen)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept: %s", err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Printf("read: %s", err)
		return
	}

	if bytes.Compare(buf[0:packetHeaderSize], packetHeader) != 0 {
		conn.Write([]byte(""))
		return
	}

	var packet Packet
	length := binary.LittleEndian.Uint32(buf[packetHeaderSize : packetHeaderSize+8])
	err = json.Unmarshal(buf[packetHeaderSize+8:int(length)+packetHeaderSize+8], &packet)
	if err != nil {
		log.Printf("unmarshal: %s", err)
		return
	}

	if packet.Request != SenderDataReq {
		log.Printf("skipping unsupported request type: %s", packet.Request)
		return
	}

	respData, err := Response{
		Response: "success",
		Info: fmt.Sprintf("processed: %d; failed: 0; total: %d; seconds spent: 0.01235813",
			len(packet.Data), len(packet.Data)),
	}.Packet()
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	// `log` package is redundant here; we already have a timestamp from packet data.
	fmt.Println(packet.String())
	conn.Write(respData)
	conn.Close()
}
