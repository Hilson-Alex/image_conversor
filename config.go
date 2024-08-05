package main

import (
	"context"

	"github.com/Hilson-Alex/image_converter/conversor"
)

var app = conversor.NewApp()

// startup is called at application startup
func startup(ctx context.Context) {
	conversor.SetAppContext(app, ctx)
}

// domReady is called after front-end resources have been loaded
func domReady(ctx context.Context) {}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func shutdown(ctx context.Context) {}
