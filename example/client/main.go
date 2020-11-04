package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	gate "github.com/Code-Hex/grpc-gate"
	"github.com/go-sql-driver/mysql"
)

func main() {
	dialer, err := gate.NewDialer("127.0.0.1", 4000)
	if err != nil {
		log.Fatal(err)
	}
	mysql.RegisterDial("grpc", func(addr string) (net.Conn, error) {
		log.Printf("connecting to %s", addr)
		return dialer.Dial("tcp", addr)
	})

	db, err := sql.Open("mysql", "root@grpc(127.0.0.1:3306)/mysql")
	if err != nil {
		log.Fatalf("open failed: %v", err)
	}
	rows, err := db.Query("SELECT * FROM time_zone")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("result: %v\n", rows)
}
