package test

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestYear(t *testing.T){
	str := "0001-01-01"
	date, err := time.Parse("2006-01-02", str)
	assert.NoError(t, err)
	assert.Equal(t, 1, date.Year())
}