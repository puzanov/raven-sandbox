package main

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/getsentry/raven-go"
)

func TestSendEventWithEmptyDSN(t *testing.T) {
	raven.CaptureErrorAndWait(errors.New("test"), map[string]string{"browser": "Firefox"})
	t.Fail()
}
