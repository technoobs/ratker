package repository

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// ConnProperty defines the structure of database connection property
type ConnProperty struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

// ServiceRepository is the database connection for service
type ServiceRepository struct {
	Name string
	Conn *sql.DB
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

// Repository function
func Repository(serviceName string, dnsString string) *ServiceRepository {
	db, err := sql.Open("mysql", dnsString)
	if err != nil {
		e := fmt.Errorf("Encountered error when intiating repository: %s", err)
		error.Error(e)
	} else {
		fmt.Println("Database initiating success...")
	}

	err = db.Ping()
	if err != nil {
		e := fmt.Errorf("Database ping is not success")
		error.Error(e)
	} else {
		fmt.Println("Database ping success....")
	}

	serviceRepository := ServiceRepository{Name: serviceName, Conn: db}
	return &serviceRepository
}
