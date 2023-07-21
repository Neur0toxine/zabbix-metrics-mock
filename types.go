package main

import (
    "bytes"
    "encoding/binary"
    "encoding/json"
    "fmt"
    "strings"
    "time"
)

const SenderDataReq = "sender data"

// Metric class.
type Metric struct {
    Host  string `json:"host"`
    Key   string `json:"key"`
    Value string `json:"value"`
    Clock int64  `json:"clock"`
}

// Packet class.
type Packet struct {
    Request string    `json:"request"`
    Data    []*Metric `json:"data"`
    Clock   int64     `json:"clock"`
}

type Response struct {
    Response string `json:"response"`
    Info     string `json:"info"`
}

func (p Packet) String() string {
    var sb strings.Builder
    sb.Grow(512)
    sb.WriteString(fmt.Sprintf("packet type `%s`, timestamp: %s", p.Request, time.Unix(p.Clock, 0)))
    if len(p.Data) == 0 {
        sb.WriteString(", no metrics data.")
        return sb.String()
    }

    sb.WriteString("\n")
    for _, item := range p.Data {
        if item == nil {
            continue
        }
        sb.WriteString(fmt.Sprintf(" - [%s] %s: %s\n", item.Host, item.Key, item.Value))
    }
    sb.WriteString("end of packet")
    return sb.String()
}

func (r Response) Packet() ([]byte, error) {
    data, err := json.Marshal(r)
    if err != nil {
        return data, err
    }

    size := uint64(len(data))

    var buf bytes.Buffer

    sizeBuf := make([]byte, 8)
    binary.LittleEndian.PutUint64(sizeBuf, size)

    buf.Write(packetHeader)
    buf.Write(sizeBuf)
    buf.Write(data)

    return buf.Bytes(), nil
}
