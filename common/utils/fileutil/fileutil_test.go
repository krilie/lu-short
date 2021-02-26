package fileutil

import "testing"

func TestGetFileExtension(t *testing.T) {
	t.Log(GetFileExtension(""))
	t.Log(GetFileExtension("."))
	t.Log(GetFileExtension(".23"))
	t.Log(GetFileExtension(".txt"))
	t.Log(GetFileExtension("1.txt"))
	t.Log(GetFileExtension("  .  "))
	t.Log(GetFileExtension("121.png"))
	t.Log(GetFileExtension("33. jpg"))
}
