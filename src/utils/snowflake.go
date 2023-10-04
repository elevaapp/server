package utils

import (
	"math/rand"

	"github.com/bwmarrin/snowflake"
)

func GenerateSnowflake() (string, error) {
	minNodeIdValue := 0
	maxNodeIdValue := 1023

	nodeId := rand.Intn(maxNodeIdValue-minNodeIdValue) + minNodeIdValue
	node, err := snowflake.NewNode(int64(nodeId))
	if err != nil {
		return "", err
	}

	return node.Generate().String(), nil
}
