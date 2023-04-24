Feature: Handler Test
  Scenario: Good POST request
    Given a user request POST with payload:
      | yoonjeong |
      | golang    |
      | cucumber  |
    Then the response code should be 200
    And the response body should be:
      | Hello, yoonjeong! |
      | Hello, golang!    |
      | Hello, cucumber!  |

  Scenario: Bad POST request
    Given a user request POST with empty payload
    Then the response code should be 400

  Scenario: GET request
    Given a user request GET
    Then the response code should be 405