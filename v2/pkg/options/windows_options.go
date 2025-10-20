//go:build windows
// +build windows

package options

import (
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// WindowsOptions is the actual Windows options type on Windows
type WindowsOptions = windows.Options
