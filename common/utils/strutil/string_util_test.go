package strutil

import (
	"lu-short/common/utils/jsonutil"
	"testing"
)

func TestToJsonPretty(t *testing.T) {
	pretty := jsonutil.ToJsonPretty(123231)
	t.Log(pretty)
}
