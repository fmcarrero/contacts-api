package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewSnowflakeID(t *testing.T) {

	assert.Greater(t, NewSnowflakeID(), int64(0))
	assert.NotEmpty(t, NewSnowflakeID())
	assert.GreaterOrEqual(t, time.Unix(snowflake.ParseInt64(NewSnowflakeID()).Time(), 0), time.Now())
}
