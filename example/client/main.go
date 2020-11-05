package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	gate "github.com/Code-Hex/grpc-gate"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	dialer, err := gate.NewDialer("127.0.0.1:4000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),
	)
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

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	var (
		timeZoneID     string
		useLeapSeconds string
	)
	row := db.QueryRow("SELECT * FROM time_zone WHERE Time_zone_id = 1")
	if err := row.Scan(&timeZoneID, &useLeapSeconds); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result: %q, %q\n", timeZoneID, useLeapSeconds)

	db.Close()
}
