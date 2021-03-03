package appdig

import (
	"go.uber.org/dig"
)

type AppContainer struct {
	*dig.Container
}

func NewAppDig() *AppContainer {
	return &AppContainer{
		Container: dig.New(),
	}
}

func (c *AppContainer) MustProvide(constructor interface{}, opts ...dig.ProvideOption) *AppContainer {
	CheckErr(c.Container.Provide(constructor, opts...))
	return c
}

func (c *AppContainer) MustProvides(constructors []interface{}, opts ...dig.ProvideOption) *AppContainer {
	for _, constructor := range constructors {
		CheckErr(c.Container.Provide(constructor, opts...))
	}
	return c
}

func (c *AppContainer) MustInvoke(function interface{}, opts ...dig.InvokeOption) *AppContainer {
	CheckErr(c.Container.Invoke(function, opts...))
	return c
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
