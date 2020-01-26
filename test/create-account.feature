Feature: Create Account
  In order to create an Account
  As the API user
  I need to set OrganisationID, Country, Bic and BankID

  Rules:
    - Country is only one of a set
    - Country conforms to ISO 3166-1 code
    - Bic is either 8 or 11 character format
    - BankID is the local country bank identifier and the format depends on the Country
    - BankID maximum length is 11

  Scenario: Create an Account in GB
    Given the country is "GB"
    And the Bic is "NWBKGB22"
    And the BankID is "400300"
    And the OrganisationID is "369ace13-f19c-413f-b512-2bbacfacf5a3"
    Then the Account should be created in memory
    And the Account should be created over the API
    And the CreatedOn date should be today
    And the ModifiedOn date should be today
    And Version should be zero
    And the OrganisationID should be "369ace13-f19c-413f-b512-2bbacfacf5a3"

  Scenario: Create an Account should fail without a country
    Given the country is ""
    Then the Account should be created in memory
    And the Account should not be created over the API

  # Scenario: Create an Account in GB without a Bic
  #   Given the country is "GB"
  #   And the Bic is ""
  #   Then the Account should be created in memory
  #   And the Account should not be created over the API

  # Scenario: Create an Account in GB without a BankID
  #   Given the country is "GB"
  #   And the BankID is ""
  #   Then the Account should be created in memory
  #   And the Account should not be created over the API

