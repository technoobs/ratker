package repository

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type ConnProperty struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

const tcpProtocol string = "tcp"
const udpProtocol string = "udp"

// BuildConnString function builds the database DNS string
func BuildConnString(conn *ConnProperty) (*string, error) {
	if !strings.EqualFold(conn.Protocol, tcpProtocol) && !strings.EqualFold(conn.Protocol, udpProtocol) {
		err := fmt.Errorf("Unsupported connection protocol: %s", conn.Protocol)
		return nil, err
	}
	var buf bytes.Buffer
	buf.WriteString(conn.UserName + ":" + conn.Password)
	buf.WriteString("@" + strings.ToLower(conn.Protocol) + "(" + conn.Host + ":" + strconv.Itoa(conn.Port) + ")")
	buf.WriteString("/" + conn.Database)
	result := buf.String()
	return &result, nil
}
