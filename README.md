# Atlas Corporation DNS Service

### Overview

The service runs an API server with a single endpoint `/api/v1/dns` that 
takes X, Y, Z coordinates and velocity of the drone as input and responds with 
location.

### Sample I/O

##### Input
```
{
	"x": "123.5",
	"y": "234.5",
	"z": "456.5",
	"vel": "1000.5"
}
```

##### Output

```
{
    "loc": 5073
}
```

### Run Instructions

Local

While in directory `PORT=9090 go run ./cmd/dns/main.go`

Docker

`docker build -t dns .`

`docker run -p 9090:9090 -e PORT=9090 dns`

### Question/Answers

**What instrumentation this service would need to ensure its observability and operational
transparency?**

- Load/stress tests
- Log events of API requests with metrics of time taken for API response, 
and the response codes to observe how much of the requests are successful

**Why throttling is useful (if it is)? How would you implement it here?**

- Throttling is useful because in the presence of a lot of drones our servers 
might be overloaded and we would like to serve all the drones consistently
- Rate limiting can be easily implemented here by a middleware that uses the token
bucket algorithm 

**What we have to change to make DNS be able to service several sectors at the same
time?**

- We would need to take the Sector ID as input, either in the url 
or the request API JSON

**Our CEO wants to establish B2B integration with Mom's Friendly Robot Company by
allowing cargo ships of MomCorp to use DNS. The only issue is - MomCorp software
expects loc value in location field, but math stays the same. How would you
approach this? What’s would be your implementation strategy?**

- Ideally we should build an intermediate service that provides API to serve Mom's 
Friendly Robot Company and the service will call our original DNS service 
and modify the response
- Or simple way to tackle this could be to introduce a new API endpoint, really
depends on the use case

**Atlas Corp mathematicians made another breakthrough and now our navigation math is
even better and more accurate, so we started producing a new drone model, based on
new math. How would you enable scenario where DNS can serve both types of clients?**

- API versioning i.e `/api/v2/dns`

**In general, how would you separate technical decision to deploy something from
business decision to release something?**

- Make canary deployments for A/B testing, keep observing and slowly roll 
out to more and more customers
- If successful, announce release to customers 