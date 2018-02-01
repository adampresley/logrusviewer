(function () {
	"use strict";

	window.TemplateService = {
		load: function (name) {
			return new Promise(function (resolve, reject) {
				qwest.get("/www/logrusviewer/templates/" + name + ".hbs")
					.then(function (xhr, response) {
						return resolve(Handlebars.compile(response));
					})
					.catch(function (err) {
						return reject(err);
					});
			});
		}
	}
}());