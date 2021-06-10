# BDD and API's

> Example how to use BDD scenarios to test APIs

## Why it's important 
Behavioral Driven Development (BDD) provides us organized conversation between developers and domain experts.
- Acts as E2E testing for infrastructure with all dependencies.
- Each feature (Story) provides a piece of domain logic in the same cases can give clues about business processes.

## Additional resources to read:
- [Five why](https://en.wikipedia.org/wiki/Five_whys)
- [The Practical Test Pyramid](https://martinfowler.com/articles/practical-test-pyramid.html)
- [Pickled State](https://blog.cleancoder.com/uncle-bob/2018/06/06/PickledState.html)

## Usage of example
It's based on [godog](https://github.com/cucumber/godog)

In root of this demo, just execute `godog`
Example of output:
```text
┌──(linmad㉿m)-[/devil/go/coopnorge/demos/week23/bdd-api]
└─$ godog
Feature: Consuming a joke from Chucknorris.io API

  Scenario: Get a Joke JSON schema                                                          # features/joke_api.feature:3
    Given I set header "Content-Type" with value "application/json"                         # api_actor_test.go:76 -> github%2ecomcoopnorgedemosweek23bdd-api.iSetHeaderWithValue
    When I send "GET" request to "https://api.chucknorris.io/jokes/random?category=science" # api_actor_test.go:84 -> github%2ecomcoopnorgedemosweek23bdd-api.iSendRequestTo
    Then the response code should be 200                                                    # api_actor_test.go:117 -> github%2ecomcoopnorgedemosweek23bdd-api.theResponseCodeShouldBe
    And the response should be a valid json                                                 # api_actor_test.go:125 -> github%2ecomcoopnorgedemosweek23bdd-api.theResponseShouldBeValidJSON
    Then the response should match json schema "testdata/schemas/joke.json"                 # api_actor_test.go:134 -> github%2ecomcoopnorgedemosweek23bdd-api.theResponseShouldMatchJsonSchema

1 scenarios (1 passed)
5 steps (5 passed)
264.222239ms
```