Feature: Delete Account
  In order to delete an Account
  As the API user
  I need the correct Account ID

  Rules:
    - Account ID is a UUID

  Scenario: Delete an exiting Account
    Given the Account exists for known ID
    Then the Delete call should be successfull
    And the Account does not exist with the known ID

  Scenario: Delete an non-existing Account
    Given the Account does not exist with the known ID
    Then the Delete call should be successfull

