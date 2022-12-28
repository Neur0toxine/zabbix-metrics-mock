package main

import (
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
