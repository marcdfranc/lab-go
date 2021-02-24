const puppeteer = require("puppeteer");

const Robot = function (params) {
	this.baseUrl = params.url;
	this.url;
	this.pages;
	this.browser;
	const linkSeletor = ".xLon";
	
	loadBrowser = async function (params) {
		this.pages = [];
		for (let i = 0; i < 10; i++) {
			this.pages.push(i);		
		}
		this.browser = await puppeteer.launch({headless: false});
	}

	this.doLogin = function() {

	}

	this.findLink = function () {
		var index = -1;
		
	}


	this.findLink = function () {
		
	}





	// create an Array of pages and test if the login works in each page.

	// search Link and open the schedule page.

	// search the days/time of work

	// book my schedule

	// check success Message.

	this.executar = function () {
		var a = 6;
		var b = 3;
		return a + b;
	}

	loadBrowser();
	console.log(this.pages);
}

module.exports = Robot;