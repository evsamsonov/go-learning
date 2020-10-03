Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Eat 5 out of 12
    Given There are 12 godogs
    When I eat 5
    Then There should be 7 remaining
