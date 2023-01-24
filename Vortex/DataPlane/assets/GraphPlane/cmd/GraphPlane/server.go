/*
 * Written in 2021 by Nicholas S. Damuth
 * V.1.0
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
 */
 package main

 import (
	//  "os"
	 "fmt"
	 "net"
	 "strings"
	 "net/http"
	//  "crypto/tls"
	 "google.golang.org/grpc"
	 "github.com/spf13/viper"
	 "github.com/cockroachdb/cmux"
	 "github.com/zenazn/goji/bind"
	 "github.com/zenazn/goji/graceful"
	 "damuth.nick/GraphPlane/internal/Logger"
	 "github.com/grpc-ecosystem/grpc-gateway/runtime"
	 // "github.com/golang/protobuf/proto"
 )
 var loaded_port interface{}
 var gracefullyStopped bool = false
 
 type ServiceHandlers struct {
 }
 
 func ServeAndWait(port int, configs *viper.Viper, listen *net.Listener) {
	logger.LogBCast("INFO", "Attempting to Server & Wait gRPC & HTTP Services")
	 if (configs.Get("server.port") != nil) {
		 loaded_port = int((configs.Get("server.port")).(float64))
		 if (loaded_port != nil) {
			 port = int((configs.Get("server.port")).(float64))
		 }
	 }
	 startServer(port, listen)
	 if (gracefullyStopped == false) {
		 ServeAndWait(port, configs, listen)
	 }
 }
 func StopServer(listen *net.Listener) {
	(*listen).Close()
 }
 func startServer(port int, listen *net.Listener) {
	var err error = nil
	*listen, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.LogBCast("ERROR", fmt.Sprintf("failed to listen: %v", err))
	}

	gSvrOpts := setupGrpcServerOptions()
	grpcServer := grpc.NewServer(gSvrOpts...)

	svrMuxOpts := setupServeMuxOptions()
	grpcGatewayMux := runtime.NewServeMux(svrMuxOpts...)

	multiplexer := cmux.New(*listen)
	grpcL := multiplexer.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := multiplexer.Match(cmux.Any())
	httpMux := http.ServeMux{}
	httpMux.Handle("/", grpcGatewayMux)

	eps := make(chan error, 2)
	go func() { eps <- grpcServer.Serve(grpcL) }()
	go func() { eps <- graceful.Serve(httpL, grpcHandlerFunc(grpcServer, httpMux)) }()

	// Here we handle some signals so that we can stop the server when running.
	graceful.HandleSignals()
	// Handle gracefully stopping the server on signals.
	bind.Ready()
	graceful.PreHook(func() {
		gracefullyStopped = true
		logger.Log("Server received signal, gracefully stopping.")
	})
	graceful.PostHook(func() {
		logger.Log("Server stopped")
	})
	if (multiplexer != nil && listen != nil) {
		logger.LogBCast("INFO", fmt.Sprintf("listening and serving (multiplexed) on: %d", port))
		err = multiplexer.Serve()
		if err != nil {
			logger.LogBCast("ERROR", fmt.Sprintf("ListenAndServe: %s", err))
		}
	} else {
		logger.LogBCast("ERROR", fmt.Sprintf("Nil multiplexer: %s", err))
	}
 }
 func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (strings.Contains(r.Header.Get("Content-Type"), "application/grpc")) {
			logger.LogByType("ACCESS", fmt.Sprintf("Access via gRCP : %s", r.Header))
			grpcServer.ServeHTTP(w, r)
		} else {
			logger.LogByType("ACCESS", fmt.Sprintf("Access via HTTP1: %s", r.Header))
			otherHandler.ServeHTTP(w, r)
		}
	})
}