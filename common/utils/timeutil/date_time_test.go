package timeutil

import (
	"fmt"
	"testing"
	"time"
)

func TestGetBeijingZeroTime(t *testing.T) {
	fmt.Println(time.Now())
	fmt.Println(GetBeijingZeroTime(time.Now()))
	fmt.Println(GetBeijingMonthStartTime(time.Now()))
	fmt.Println(GetBeijingLastDateOfMonth(time.Now()))
}
