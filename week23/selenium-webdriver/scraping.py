"""Entrypoint for Scraping statens vegvesen for time slots"""
import time
import sys
from selenium import webdriver
from chromedriver_py import binary_path
from selenium.webdriver.support.ui import Select
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from selenium.common.exceptions import (
    TimeoutException,
)


class TimeSlotScraper:
    def __init__(self, station):
        self.driver = None
        self.station = station
        self.available_times = None

    def create_driver(self):
        """Initiates a new driver"""
        chrome_options = webdriver.ChromeOptions()
        chrome_options.add_argument("--headless")
        self.driver = webdriver.Chrome(binary_path, options=chrome_options)

    def wait_until(self, xpath, timeout=10):
        """Method for waiting until an element loads on screen

        Args:
            xpath (string): The xpath to the element to wait for.
            timeout (int, optional): Number of seconds before timeout. Defaults to 10.

        Raises:
            [SystemExit]: If element cannot be found
        """
        try:
            WebDriverWait(self.driver, timeout).until(
                EC.presence_of_element_located((By.XPATH, xpath))
            )
        except TimeoutException:
            print(f"failed to find element with xpath: {xpath}")
            sys.exit()

    def pick_service(self):
        """Picks Teori as the service"""
        button_elm = '//*[@id="Main"]/div[2]/div[1]/div/form/div[2]/input'
        self.driver.get(
            "https://ventus.enalog.se/Booking/Booking/Index/VegvesenRislokka"
        )
        self.wait_until(button_elm)
        selector_elm = Select(self.driver.find_element_by_id("ServiceGroupId"))
        selector_elm.select_by_visible_text("Teori")
        self.driver.find_element_by_xpath(button_elm).click()

    def accept_terms(self):
        """Accepting terms and conditions"""
        check_box = '//*[@id="AcceptInformationStorage"]'
        next_btn = '//*[@id="Main"]/form/div[2]/input'
        self.wait_until(check_box)
        self.driver.find_element_by_xpath(check_box).click()
        self.driver.find_element_by_xpath(next_btn).click()

    def pick_station(self):
        """Picks Trafikkstasjon based on user input"""
        button_elm = '//*[@id="Main"]/form[1]/div/div[5]/div/input[2]'
        self.wait_until(button_elm)
        selector_elm = Select(self.driver.find_element_by_id("SectionId"))
        selector_elm.select_by_visible_text(self.station)
        time.sleep(1)
        self.driver.find_element_by_xpath(button_elm).click()

    def find_available_times(self):
        """Updates self.available_times to the current available time slots"""
        self.wait_until('//*[@id="Main"]/form[2]/div[2]/table/thead')
        time_slots = '//*[@id="Main"]/form[2]/div[2]/table/tbody/tr[*]/td[*]/div/div'
        slots = self.driver.find_elements_by_xpath(time_slots)
        available_times = []
        for slot in slots:
            if slot.text != "Booked":
                available_times.append(slot.get_attribute("aria-label"))
        self.available_times = available_times
        self.driver.quit()

    def print_availeble_times(self):
        """Prints out the available times to terminal"""
        print("These are the available times")
        for time in self.available_times:
            print(time)


def main():
    """Main function initiating the web scraper"""
    station = sys.argv[1]
    client = TimeSlotScraper(station)
    client.create_driver()
    client.pick_service()
    client.accept_terms()
    client.pick_station()
    client.find_available_times()
    client.print_availeble_times()


if __name__ == "__main__":
    main()
