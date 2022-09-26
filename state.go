package gosd

// valid states
// see: https://www.freedesktop.org/software/systemd/man/sd_notify.html#

const (
	// ReadyState tells the service manager that service startup is finished, or the service finished loading its configuration
	ReadyState = "READY=1"

	ReloadingState = "RELOADING=1" // send ReadyState once reloading is done

	StoppingState = "STOPPING=1"

	WatchdogState = "WATCHDOG=1" // Tells the service manager to update the watchdog timestamp

	WatchdogTrigger = "WATCHDOG=trigger"
)
