package klog

/**
 * Global configuration for log
 */
const (
	LogLevel        = Debug     // Log level, value in [VERBOSE, DEBUG, INFO, WARN, ERROR], just like in Android
	FileLogPath     = "testlog" // Log file name
	FileLogEnable   = true      // Enable file log
	FileLogMaxLines = 50000     // Max lines each log file
	FileLogMaxSize  = 5000000   // Max size
)
