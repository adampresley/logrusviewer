package main

import "flag"

var host = flag.String("host", "0.0.0.0:8080", "Address and port to run this application on. Defaults to 0.0.0.0:8080")
var logLevel = flag.String("loglevel", "info", "Level of logs to write. Valid values are 'debug', 'info', or 'error'. Default is 'info'")
