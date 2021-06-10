"""e2e Test cases for copp.no"""
import unittest
from selenium import webdriver
from chromedriver_py import binary_path
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from selenium.common.exceptions import TimeoutException


class CoopTests(unittest.TestCase):
    """Main class of unittest"""

    def setUp(self):
        """Initiates a Chrome Webdriver with Selenium"""
        self.driver = webdriver.Chrome(binary_path)

    def tearDown(self):
        """Quits the webdriver"""
        self.driver.quit()

    def wait_until(self, xpath, timeout=10):
        """Method for waiting until an element loads on screen

        Args:
            xpath (string): The xpath to the element to wait for.
            timeout (int, optional): Number of seconds before timeout. Defaults to 10.

        Returns:
            [bool]: True if page is loaded, false if not
        """
        try:
            WebDriverWait(self.driver, timeout).until(
                EC.presence_of_element_located((By.XPATH, xpath))
            )
        except TimeoutException:
            return False
        return True

    def test_get_recipe(self):
        """e2e test for using search bar and finding recipe on coop.no"""
        driver = self.driver
        driver.get("https://coop.no/")
        self.assertIn("Coop litt ditt", driver.title)
        print("Test passed, found 'Coop litt ditt' in title of page")

        search_bar = '//*[@id="spa-header"]/header/div[2]/div/form/input'
        self.failIf(self.wait_until(search_bar))
        self.driver.find_element_by_xpath(search_bar).send_keys("Rullekake")
        search_button = '//*[@id="spa-header"]/header/div[2]/div/form/button'
        self.driver.find_element_by_xpath(search_button).click()
        print("Successfully entered 'rullekake' and clicked on search button")

        recipe = '//*[@id="spa-main"]/main/div/div/div[3]/div[2]/div/a[1]/div[1]'
        recipe_title = (
            '//*[@id="spa-main"]/main/div/div/div[3]/div[2]/div/a[1]/div[2]/h4'
        )
        recipe_link = '//*[@id="spa-main"]/main/div/div/div[3]/div[2]/div/a[1]'
        expected = "Enkel rullekake med jordbær og sjokolade"
        self.failIf(self.wait_until(recipe))
        print("Successfully found the recipe I was looking for")
        title_text = self.driver.find_element_by_xpath(recipe_title).text
        self.assertEqual(title_text, expected)
        print("Test passed, the title of the recipe is: ", expected)
        self.driver.find_element_by_xpath(recipe_link).click()

        cake_img = "/html/body/div[1]/main/section/header/div[1]/img"
        self.failIf(self.wait_until(cake_img))
        print("recipe is loaded")
        ingredience_elm = self.driver.find_element_by_xpath(
            "/html/body/div[1]/main/section/section/p[12]/strong"
        )
        self.driver.execute_script(
            "arguments[0].scrollIntoView(true);", ingredience_elm
        )
        print("scrolled down to ingrendience list")

        expected_url = "https://coop.no/mega/oppskrifter/kaker/enkel-rullekake-med-jordbar-og-sjokolade/"
        actual_url = self.driver.current_url
        self.assertEqual(expected_url, actual_url)
        print("Test passed The url of the recipe is correct")

        ingredience_list = self.driver.find_elements_by_xpath(
            "/html/body/div[1]/main/section/section/ul[3]/li"
        )
        actual_amount = len(ingredience_list)
        expected_amount = 7
        self.assertEqual(expected_amount, actual_amount)
        print("Test passed, there are 7 ingrediences in this recipe")

        actual = ingredience_list[0].text
        expected = "1 ts smør"
        self.assertEqual(actual, expected)
        print(f"Test passed, the first ingredience is '{actual}'")


if __name__ == "__main__":
    unittest.main()
