(function () {
	"use strict";

	window.LogFileService = {
		getLogEntry: function (logFile, lineNumber) {
			return qwest.get("/entry?logfile=" + logFile + "&lineNumber=" + (lineNumber - 1));
		}
	};
}());