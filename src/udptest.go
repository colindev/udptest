package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {

	var (
		subcmd string
		ip     string
		port   int
	)

	if len(os.Args) < 2 {
		fmt.Println("miss subcommand server/client")
		os.Exit(1)
	}

	subcmd = os.Args[1]
	flag.CommandLine.StringVar(&ip, "ip", "", "Server IP")
	flag.CommandLine.IntVar(&port, "port", 1194, "UDP port")
	flag.CommandLine.Parse(os.Args[2:])

	switch subcmd {
	case "server":
		server(port)
	case "client":
		client(net.ParseIP(ip), port)
	default:
		fmt.Printf("Unknown subcommand [%s]\n", subcmd)
		os.Exit(1)
	}

}

func server(port int) {

	listenAddr := &net.UDPAddr{Port: port}
	conn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("listen UDP:", listenAddr)

	for {
		buf := make([]byte, 512)
		_, addr, err := conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Println(addr, err)
			continue
		}

		go func(b []byte, addr *net.UDPAddr) {
			fmt.Printf("receive [%s] from [%s]\n", string(b), addr)
			_, err := conn.WriteTo(b, addr)
			if err != nil {
				fmt.Println(addr, err)
			}
		}(buf, addr)
	}
}

func client(ip net.IP, port int) {

	dstUDPAddr := &net.UDPAddr{IP: ip, Port: port}

	conn, err := net.DialUDP("udp", &net.UDPAddr{}, dstUDPAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("conn from [%s] to [%s]\n", conn.LocalAddr(), conn.RemoteAddr())

	buf := bufio.NewReader(os.Stdin)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			fmt.Println(err)
			return
		}

		conn.Write(line)

		receive := make([]byte, 512)
		_, addr, err := conn.ReadFromUDP(receive)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("receive [%s] from [%s]\n", receive, addr)

	}

}
