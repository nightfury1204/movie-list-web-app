package main

import (
	"github.com/nightfury1204/movie-listing-app/cmds"
	"github.com/nightfury1204/movie-listing-app/pkg/logger"
	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	logger.Initialize()
	defer klog.Flush()

	if err := cmds.NewRootCmd().Execute(); err != nil {
		klog.Fatal(err)
	}
}
