package klog

/**
 * Global configuration for log
 */
const (
    LOG_LEVEL          = DEBUG     // Log level, value in [VERBOSE, DEBUG, INFO, WARN, ERROR], just like in Android
    FILE_LOG_PATH      = "testlog" // Log file name
    FILE_LOG_ENABLE    = false     // Enable file log
    FILE_LOG_MAX_LINES = 50000     // Max lines each log file
    FILE_LOG_MAX_SIZE  = 5000000   // Max size
)
