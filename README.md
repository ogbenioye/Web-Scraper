## Web-Scraper Application
### Introduction
After learning Go, this was the first application I built. A web-scraper application built with Colly framework that later became useful in projects where I was interning. Understanding Colly and getting the app to work was the difficult part, after I figured that out, I only needed to modify and make it better each time.
What you're looking at is a web scaper application engineered to scrape targeted data(product names and prices) from the clothing section of an e-commerce web app.
  
### How It Works
Here is how the app works in 5 stages:
1. The app creates and opens a file
2. Using Colly, it accesses the allowed domain(s) and sends requests to the domain(s).
3. Using HTML tags, class and ids to identify the targeted data, the app scrapes the data on that page and writes it into the opened file.
4. It visits and sends requests to other pages specified and repeats the 3rd stage.
5. Closes the file and exits program.
