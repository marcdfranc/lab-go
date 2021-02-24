// const puppeteer = require('puppeteer')
const Credentials = require("./src/environment/credentials");
const Robot = require("./src/Robot");

async function robot() {	
	var robot = new Robot({
		url: Credentials.uri
	});	

	console.log(robot);

	
	
	
	
	
	/*const browser = await puppeteer.launch({headless: false });
	const page = await browser.newPage();
	await page.goto('https://example.com');
	await page.screenshot({ path: 'example.png' });	
	
	
	// await browser.close();	*/
}

robot();