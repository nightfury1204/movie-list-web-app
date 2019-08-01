package logger

import (
	"github.com/go-logr/logr"
	"k8s.io/klog/klogr"
)

var Logr logr.Logger

func Initialize() {
	// change the name, if app name changes
	Logr = klogr.New().WithName("movie-listing-app")
}

func GetLogger() logr.Logger {
	if Logr == nil {
		Logr = klogr.New().WithName("movie-listing-app")
	}
	return Logr
}
