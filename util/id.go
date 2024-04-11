package util

import (
	"github.com/bwmarrin/snowflake"
)

var snowNode, _ = snowflake.NewNode(0)

func NextSnowId() int64 {
	return snowNode.Generate().Int64()
}
