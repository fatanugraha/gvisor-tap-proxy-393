package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:50406")
			if err != nil {
				panic(err)
			}

			d := net.Dialer{
				Timeout:   time.Millisecond * time.Duration(10000),
				KeepAlive: -1,
				LocalAddr: addr,
			}

			conn, err := d.DialContext(ctx, network, "8.8.8.8:53")
			if err != nil {
				panic(err)
			}

			fmt.Println("LocalAddr: ", conn.LocalAddr())

			return conn, err
		},
	}

	lookup := func() {
		fmt.Printf("%s starting LookupIP\n", time.Now())
		_, err := r.LookupIP(context.Background(), "ip4", "www.google.com")
		if err != nil {
			fmt.Println("err", err)
		} else {
			fmt.Println("ok")
		}
	}

	lookup()                     // ok
	time.Sleep(95 * time.Second) // wait for the UDPConnTimeout
	lookup()                     // this will fail after 2 retries
}
