// Go NETCONF Client - Example
//
// Copyright (c) 2013-2018, Juniper Networks, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"log"

	"github.com/Juniper/go-netconf/netconf"
	"golang.org/x/crypto/ssh"
)

func main() {
	sshConfig := &ssh.ClientConfig{
		User:            "sample",
		Auth:            []ssh.AuthMethod{ssh.Password("sample")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	fmt.Println("Dial ssh connection")
	s, err := netconf.DialSSH("127.0.0.1:22", sshConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ssh connection successful")
	defer s.Close()
	fmt.Println(s.ServerCapabilities)
	fmt.Println(s.SessionID)
	// Sends raw XML
	// Get chassis inventory; return data is XML
	//reply, err := s.Exec(netconf.RawMethod("<get-chassis-inventory/>"))
	// show bgp summary; return data is JSON
	//reply, err := s.Exec(netconf.RawMethod("<get-config format=\"json\"/>"))
	//reply, err := s.Exec(netconf.RawMethod("<get-config><source><running/></source></get-config>"))
	//reply, err := s.Exec(netconf.RawMethod("<get-config> <source><running/></source> <filter type=\"subtree\"> <configuration> <system/> </configuration> </filter> </get-config>"))
	reply, err := s.Exec(netconf.RawMethod("<get-config> <source><running/></source> <filter type=\"subtree\"> <configuration> <protocols> <mld/> </protocols></configuration> </filter> </get-config>"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Reply: %+v", reply)
}
