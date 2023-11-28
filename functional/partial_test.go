package functional

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	localhost := "127.0.0.1"
	var maxConnCnt uint32 = 1024
	localAddr := WithAddr(localhost)
	maxConn := WithMaxConn(maxConnCnt)

	c1 := NewConfig(localAddr)
	require.Equal(t, Config{Addr: localhost}, c1)

	c2 := NewConfig(maxConn)
	require.Equal(t, Config{MaxConn: maxConnCnt}, c2)

	c3 := NewConfig([]Option{localAddr, maxConn}...)
	require.Equal(t, Config{
		Addr:    localhost,
		MaxConn: maxConnCnt,
	}, c3)

}
