package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// exists returns true if the domain was registered.
// If you need available domain, oposite flag.
func exists(domain string) (bool, error) {
	const whoisServer string = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		return false, err
	}
	defer conn.Close()

	conn.Write([]byte(domain + "\r\n"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "no match") {
			return false, nil
		}
	}

	return true, nil
}

var marks = map[bool]string{true: "Available", false: "Unavailable"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")
		exist, err := exists(domain)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(marks[!exist])
		time.Sleep(1 * time.Second)
	}
}
