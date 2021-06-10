Feature: Consuming a joke from Chucknorris.io API

  Scenario: Get a Joke JSON schema
    Given I set header "Content-Type" with value "application/json"
    When I send "GET" request to "https://api.chucknorris.io/jokes/random?category=science"
    Then the response code should be 200
    And the response should be a valid json
    Then the response should match json schema "testdata/schemas/joke.json"
