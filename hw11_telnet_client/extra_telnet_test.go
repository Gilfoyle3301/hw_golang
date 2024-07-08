package main

import (
	"bytes"
	"io"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAllStart(t *testing.T) {
	t.Run("Test New Client Connect", func(t *testing.T) {
		ln, err := net.Listen("tcp", "localhost:")
		if err != nil {
			t.Fatalf("could not create listener: %v", err)
		}
		defer ln.Close()

		client := NewTelnetClient(ln.Addr().String(), 1*time.Second, nil, nil)
		if err := client.Connect(); err != nil {
			t.Errorf("Connect() error = %v, wantErr %v", err, false)
		}
		client.Close()
	})
	t.Run("Test New Client Send and Receive", func(t *testing.T) {
		in := &bytes.Buffer{}
		out := &bytes.Buffer{}
		ln, err := net.Listen("tcp", "localhost:")
		require.NoError(t, err)
		defer ln.Close()
		in.WriteString("hello\n")

		go func() {
			conn, err := ln.Accept()
			require.NoError(t, err)
			defer conn.Close()
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			require.NoError(t, err)
			require.Equal(t, "hello\n", string(buf[:n]))
			_, err = conn.Write([]byte("done test\n"))
			require.NoError(t, err)
		}()
		client := NewTelnetClient(ln.Addr().String(), 3*time.Second, io.NopCloser(in), out)
		require.NoError(t, client.Connect())
		err = client.Send()
		require.NoError(t, err)
		err = client.Receive()
		require.NoError(t, err)
		require.Equal(t, "done test\n", out.String())
		client.Close()
	})
	t.Run("Test New Client Close", func(t *testing.T) {
		ln, err := net.Listen("tcp", "localhost:")
		require.NoError(t, err)
		defer ln.Close()
		client := NewTelnetClient(ln.Addr().String(), 1*time.Second, nil, nil)
		err = client.Connect()
		require.NoError(t, err)
		err = client.Close()
		require.NoError(t, err)
	})
}
