(function () {
	"use strict";

	function getLinks() {
		return document.getElementsByClassName("logEntryLink");
	}

	function getLineNumber(el) {
		return window.parseInt(el.getAttribute("data-lineNumber"), 10);
	}

	function onLineNumberClick() {
		var lineNumber = getLineNumber(this);

		window.LogFileService.getLogEntry(window.logFile, lineNumber)
			.then(function (xhr, response) {
				render(response);
			})
			.catch(function (err) {
				alert("Error rendering log entry");
				console.log(err);
			});
	}

	function render(entry) {
		var html = logEntryDetailsTemplate({
			entry: entry
		});

		document.getElementById("detailPanel").innerHTML = html;
	}

	/*
	 * Constructor
	 */
	var logEntryDetailsTemplate;

	window.TemplateService.load("LogEntryDetails")
		.then(function (template) {
			logEntryDetailsTemplate = template;
		})
		.catch(function (err) {
			alert("Error loading log entry details template!");
			console.log(err);
		});

	window.onload = function () {
		var links = getLinks();

		[].forEach.call(links, function (el) {
			el.onclick = onLineNumberClick;
		});
	};

}());