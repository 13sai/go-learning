package designPattern_test

import (
	"testing"

	"github.com/13sai/go-learing/designPattern"
	"github.com/stretchr/testify/assert"
)

func TestNewOrganization(t *testing.T) {
	num := designPattern.NewOrganization().Count()
	t.Log("count", num)
	assert.Equal(t, 20, num)
}