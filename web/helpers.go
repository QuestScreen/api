package web

import "syscall/js"

var smartphoneMode = js.Global().Get("window").Call(
	"matchMedia", "screen and (max-width: 35.5em)")

// InSmartphoneMode returns whether the css selector used for enabling
// smartphone UI currently matches.
func InSmartphoneMode() bool {
	return smartphoneMode.Get("matches").Bool()
}

// LogLevel describes the level at which logs are made
type LogLevel int

const (
	// LogDebug is the lowest log level, used for debugging
	LogDebug LogLevel = iota
	// LongInfo is used for logging informative messages
	LogInfo
	// LogWarn is used for logging warnings
	LogWarn
	// LogError is used for logging errors
	LogError
)

func (level LogLevel) String() string {
	switch level {
	case LogDebug:
		return "debug"
	case LogInfo:
		return "info"
	case LogWarn:
		return "warn"
	default:
		return "error"
	}
}

// Log writes the given message with the given loglevel to the console.
func Log(level LogLevel, msg string) {
	js.Global().Get("console").Call(level.String(), msg)
}

// LogGrouped write the given message with the given loglevel to the console.
// It appends the content at the same level inside a group, each content item
// logged separately. group is initially collapsed when collapsed=true.
func LogGrouped(level LogLevel, msg, groupLabel string, collapsed bool, content ...string) {
	f := level.String()
	console := js.Global().Get("console")
	console.Call(f, msg)

	var gf string
	if collapsed {
		gf = "groupCollapsed"
	} else {
		gf = "group"
	}
	if groupLabel == "" {
		console.Call(gf)
	} else {
		console.Call(gf, groupLabel)
	}

	for _, item := range content {
		console.Call(f, item)
	}

	console.Call("groupEnd")
}
