# Weather kata
We cannot control the weather but we can predict it.

This kata has a code that request the weather prediction from Metaweather.

## Goal
- Test coupled code.
- Remove the external dependency when testing in order to make the tests repeatable and fast

## Steps 
1. Write a test that validates the weather forecast for today in Barcelona.
2. Make the test repeatable, fast and isolating it external calls.
   2a. Extract and override the call to get the geolocation of Barcelona.
   2b. Extract and override the call to get the prediction for that geolocation.
3. Write more tests to increase the code coverage.


## Authors
Luis Rovirosa [@luisrovirosa](https://www.twitter.com/luisrovirosa)

Jordi Anguela [@jordianguela](https://www.twitter.com/jordianguela)

[Original repository](https://github.com/CodiumTeam/weather-kata)
