package main

import (
	"context"
	"net"
	"strings"
	"testing"
	"time"
)

func TestRunShutsDown(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() {
		done <- run(ctx, "127.0.0.1:0")
	}()
	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("run: %v", err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("server did not stop after cancellation")
	}
}

func TestRunReportsListenFailure(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	err = run(context.Background(), listener.Addr().String())
	if err == nil || !strings.Contains(err.Error(), "listen:") {
		t.Fatalf("run error = %v, want listen failure", err)
	}
}
