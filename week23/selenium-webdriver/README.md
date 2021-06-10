# Demo code of Selenium with chromedriver

## Installation
In order to run these scripts you will need to have python3 installed and pip install the packages in requirements.txt

Run this command to install:
```bash
pip install -r requirements.txt
```


## Web Scraper
These days it is really difficult to find an available time for taking the theoretical exam at Statens Vegvesen.
Almost all time slots are filled up two weeks in advance, and you never know when they will drop new time slots.

This is a great opportunity to show how Selenium in Python can be used as a webscraper and prove itself usefull.

### How do I check for available times??

Run this command:
```bash
python3 scraping.py [Navn-på-trafikkstasjon]
```

> **Note:** Change [navn-på-trafikkskole] to the traffic station you want to look at. For example Risløkka for Oslo.

The python script will visit Statens Vegvesen's website for you and see if there is anything available in no time.

I's as easy as that!


## Testing Web Applications
Selenium is a great way to to automate testing for web applications.
Instead of just testing small parts of the web application you can automate and immitate the behavior of a real user.
Selenium will load the web application in a real web browser allowing us to test the application from the user perspective e2e.


### What are we testing?
For this demo I decided to create a sample test of coop.no to showcase.
The test will try to search for a recipe and check that the results are as expected.
So the main thing we are testing is really the search functionality at coop.no.

### How to run the tests?
run this command:
```bash
python3 -m unittest coop_test.py
```