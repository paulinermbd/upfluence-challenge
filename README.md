# Upfluence Challenge 🌟
This is a technical challenge to join Upfluence company.

## What?
It's an REST API with only one route that allows user to get statistics from social media publication. Those statistics are based on a SSE sends an infinite stream of data from differents social media. 
I tried to be pragmatic and Golang idiomatic. 

## How?
- net/http to build the route GET /analysis
- testing for unit tests and test coverage 
- golint et go fmt for code style and code quality

I comment the code as godoc. I add some logs to debug and knowing where errors happen. 

## How to test?
- Clone the repository 
- Check requirements : Go is installed version 1.23
- In a terminal run `go mod tidy` 
- Go into `cmd` directory
- In a terminal run `go run main.go` 
- In Postman or terminal (new one) you can run : `curl "http://localhost:8080/analysis?duration=30s&dimension=likes"` 

## Why, I did not achieve all the goals?
I'm not very familiar with coding from scratch and I do not manipulate amount of data and streams in my every day life, so it was challenging to me. 
I did not want to do complicate code because I want to be able to explain it. 
I did not have time to test the 24h : 
  - This is a really tricky case, we don't want the client to timeout so we will have to think to better solution like sending partials data or sending header telling explicitly to keep the connection open. 

### Stream data
Stream data was a discovery so I understood concepts and then start to implement buffer and scanner.
Once I save my data I wanted to parse them and then do my maths asynchronously it's not the better option but as I'm not comfortable with channels I did not want to use it. 

I would like more times to improve the data storage and not doing it all the 100 events for small duration like 5s. 


## API contract
### Route
- GET localhost:8080/analysis
  - Implement a server with net/http library
  - If I had more time I would added middlewares and a router to be cleaner than just a handler 
- Query params
  - duration (s, m, h)
  - dimension (likes,comments,favorites,retweets)

### Response
- JSON response, with an output object that I encode in JSON
````
{
  "total_posts": int64,
  "minimum_timestamp": timestamp,
  "maximum_timestamp": timestamp,
  "p50": int,
  "p90": int,
  "p99": int,
}
````