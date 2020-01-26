Feature: List Accounts
  In order to List Accounts
  As the API user

  Scenario: List all Accounts
    Given there are at least 10 Accounts
    Then the Accounts should be successfully listed

  Scenario: List all Accounts with page size of 3
    Given there are at least 10 Accounts
    Then only 3 Accounts should be listed with page size of 3

  Scenario: List all Accounts with page size of 7
    Given there are at least 10 Accounts
    Then only 7 Accounts should be listed with page size of 7

  Scenario: List all Accounts with page size of 4
    Given there are an odd number of Accounts
    Then each page of 4 contains the expected number
