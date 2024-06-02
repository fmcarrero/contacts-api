package snowflake

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func NewSnowflakeID() int64 {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(fmt.Sprintf("Error creating snowflake node: %s", err))
	}

	// Generate a snowflake ID.
	return node.Generate().Int64()
}
