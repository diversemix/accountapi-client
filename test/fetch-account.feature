Feature: Fetch Account
  In order to fetch an Account
  As the API user
  I need the correct Account ID

  Rules:
    - Account ID is a UUID

  Scenario: Fetch an exiting Account
    Given the Account exists for known ID
    Then the Account should be successfully fetched

  Scenario: Fetch an non-existing Account
    Given a random ID
    Then the Account does not exist with the known ID

