package appdig

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"
	"testing"
)

type Hello interface{ Hello() }
type HelloOne struct{}
type HelloTwo struct{}

func NewHelloOne() Hello   { return &HelloOne{} }
func NewHelloTwo() Hello   { return &HelloTwo{} }
func (h *HelloOne) Hello() { fmt.Println("form hello one") }
func (h *HelloTwo) Hello() { fmt.Println("form hello two") }

func TestAutoDig(t *testing.T) {
	container := dig.New()
	err := container.Provide(NewHelloOne)
	require.Nil(t, err)
	err = container.Invoke(func(hello Hello) {
		hello.Hello()
	})
	require.Nil(t, err)
	err = container.Provide(NewHelloTwo)
	require.Error(t, err)
	err = container.Invoke(func(hello Hello) {
		hello.Hello()
	})
	require.Nil(t, err)
}
