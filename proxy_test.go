package proxy

import (
	"testing"
	"time"
)

func TestTm(t *testing.T) {
	now := time.Now() // time.Time

	t.Log(now.Unix(), now.Nanosecond())
	time.Unix(100000, 0) // time.Time

	t.Log(now.Format("2006"))
	t.Log(now.Format("0102"))
	t.Log(now.Format("20060102"))
	t.Log(now.Format("20060102 15:04:05"))
	t.Log(now.Format("15:04:05"))
	t.Log(now.Format("0102150405"))
}
