package comb_test

import (
	"errors"
	"testing"

	"github.com/kode4food/comb"
	"github.com/stretchr/testify/assert"
)

func TestCompose(t *testing.T) {
	as := assert.New(t)

	c := comb.Compose(
		func(in bool) (bool, error) {
			if in {
				return false, nil
			}
			return false, errors.New("test error")
		},
		func(in bool) (bool, error) {
			return !in, nil
		},
	)

	r, err := c(true)
	as.True(r)
	as.Nil(err)

	r, err = c(false)
	as.False(r)
	as.EqualError(err, "test error")
}

func TestMust(t *testing.T) {
	as := assert.New(t)

	c := comb.Comb[bool, bool](func(b bool) (bool, error) {
		if b {
			return true, nil
		}
		return false, errors.New("test error")
	}).Must()

	r := c(true)
	as.True(r)

	defer func() {
		rec := recover()
		as.EqualError(rec.(error), "test error")
	}()

	r = c(false)
	as.False(r)
}

func TestBind(t *testing.T) {
	as := assert.New(t)

	c := comb.Comb[bool, string](func(in bool) (string, error) {
		if in {
			return "true", nil
		}
		return "false", nil
	}).Bind(func(s string) comb.Comb[string, string] {
		if s == "true" {
			return func(s string) (string, error) {
				return "Hello", nil
			}
		}
		return func(_ string) (string, error) {
			return "", errors.New("was false")
		}
	})

	s, err := c(true)
	as.Nil(err)
	as.Equal("Hello", s)

	s, err = c(false)
	as.Equal("", s)
	as.EqualError(err, "was false")
}
