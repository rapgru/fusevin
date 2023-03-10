/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	f "tuwien.ac.at/fusevin/fuse"
	rpc "tuwien.ac.at/fusevin/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s MOUNTPOINT\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
		os.Exit(2)
	}
	mountpoint := flag.Arg(0)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	fuseserver := &f.FuseServer{Mountpoint: mountpoint, WaitGroup: wg}  
	grpcserver := &rpc.GRPCServer{Port: *port, WaitGroup: wg, Fuseserver: fuseserver}
	fuseserver.NotifyCallback = grpcserver.Notify

	go fuseserver.Start()
	go grpcserver.Start()

/*	c := make(chan os.Signal, 1)
 	signal.Notify(c, os.Interrupt)
	go func(){
	    for sig := range c {
	        // sig is a ^C, handle it
	    }
	}() */

	wg.Wait()
}
