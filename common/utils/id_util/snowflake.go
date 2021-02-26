package id_util

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func NextSnowflakeId() snowflake.ID {
	return node.Generate()
}
func NextSnowflake() string {
	return node.Generate().String()
}
