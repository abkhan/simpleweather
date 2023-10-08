# simpleweather
A simple weather web server.

The code is done only to make a simple web server and call to openWeather api work.

It does not fully follow all good rest api and programming best practice guidelines.


## How to Run

Export env var: OPENWEATHER_API_KEY
and also, OPENWEATHER_URL 
with correct values,

then cd to cmd/service from root of the repo.
$ go run main.go
(make sure to do go mod tidy)

## Basic Test

curl -XPOST 'localhost:8080/weather' --data '{
    "lat": "30.3",
    "lon": "-80.6"
}'


### Response:

{
    "condition": "Clouds",
    "temperature": "moderate"
}

## ToDo:

The API needs bunch of improvements for it to be a good api.

Some suggestions;

 - Authentication
 - Access Control
 - Circuit breaker
 - Cache to save API costs
 - Monitoring
 - Logging (structured logging) + better logging