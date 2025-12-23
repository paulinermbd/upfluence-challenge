# Upfluence Challenge 🌟
This is a technical challenge to join Upfluence company.

## Description 

(WIP) 
- net/http to build the route GET /analysis
- testing for unit tests and test coverage 
- golint et go fmt for code style and code quality 
- archi hexa 
  - facile à moduler
  - si on veut ajouter plus de choses
  - 

## What I need ? 
- GET stats from a stream 
  - Choose the duration of analysis that user wants => query params 
  - Choose the stats that user wants => query params

## API contract
### Route
- GET localhost:8080/analysis => keep it simple at first but we could also think about versioning
- Query params
  - duration (s, m, h)
  - dimension (likes,comments,favorites,retweets) any of => cumulable ?
    - Je comprends que c'est pas cumulable il ne parle que d'une dimension

### Response
- JSON
````
{
  "total_posts": int64,
  "minimum_timestamp": timestamp,
  "maximum_timestamp": timestamp,
  "dimension_p50": int,
  "dimension_p90": int,
  "dimension_p99": int,
}
````
- percentiles can be challenged following the efficiency of calculus
- 
  


## Reasoning

First, I read the instruction and then identify what I know and what I don't know.
For things that I didn't know I looked quickly on the internet to understand concepts.

On a second time, I will identify the features and find technical tasks associated to them.
That will allow me to think about my architecture, and code organization. 

Once my ideas are clear and tasks identified, I start to implement MVP. 

The MVP objective is to connect to the SSE stream and read data for 5s. 

Then I will try to optimize and do better to join the 10m, 24h or anytime that is longer than 5s. 

## Notes 
They said they like interface and abstract object during the interview => try to stick to it 
Why 404  We could have send 405 (Method Not Allowed) => on another hand we could presume that is for safety because a 405 give a more precise info than a 404.

## Questions 
* 10min, 24h, we need to not timeout, how to deal with?
  * Do they want partial content that refresh or once the time is passed all info computed? 


