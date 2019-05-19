package myloger

import "testing"

// %T 变量类型
func TestConstLevel(t *testing.T) {
	t.Logf("%v %T\n", DebugLevel, DebugLevel)
	t.Logf("%v %T\n", ErrorLevel, ErrorLevel)
}
