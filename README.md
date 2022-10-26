# Github Metrics

This project is a PoC to explore the Github API and retrieve certain metrics from a given repository.

It will perform certain API calls using HTTP requests and backoff strategies to mitigate rate limits in the Github API.

## Labels

We'd like to get, for each label, how many issues are labelled for the label.

```shell
ghm issues labels -O testcontainers -R testcontainers-go
```

It will first get all the labels in the repository, and for each it will get the total number of issues, printing the result in console:

```
{"level":"info","ts":1666776981.5668511,"caller":"log/logger.go:38","msg":"Processing labels","count":18}
{"level":"warn","ts":1666776981.601449,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":1,"elapsedTime":0.034401667,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"compose\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666776982.191045,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":2,"elapsedTime":0.62400025,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"compose\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666776983.668726,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":3,"elapsedTime":2.101689958,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"compose\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666776986.034081,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":4,"elapsedTime":4.467057167,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"compose\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666776989.820698,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":5,"elapsedTime":8.253687333,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"compose\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666776994.496547,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":6,"elapsedTime":12.929558833,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"compose\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777000.46876,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":7,"elapsedTime":18.901799292,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"compose\"","error":"GET request failed with 403"}
{"level":"info","ts":1666777003.696355,"caller":"log/logger.go:38","msg":"compose","count":12}
{"level":"info","ts":1666777004.304099,"caller":"log/logger.go:38","msg":"dependencies","count":99}
{"level":"info","ts":1666777004.647705,"caller":"log/logger.go:38","msg":"go","count":10}
{"level":"info","ts":1666777004.954628,"caller":"log/logger.go:38","msg":"good-first-issue","count":10}
{"level":"info","ts":1666777005.331155,"caller":"log/logger.go:38","msg":"hacktoberfest","count":15}
{"level":"info","ts":1666777005.499369,"caller":"log/logger.go:38","msg":"os/mac","count":0}
{"level":"info","ts":1666777005.8398788,"caller":"log/logger.go:38","msg":"os/windows","count":2}
{"level":"info","ts":1666777006.142145,"caller":"log/logger.go:38","msg":"podman","count":7}
{"level":"info","ts":1666777006.376652,"caller":"log/logger.go:38","msg":"python","count":1}
{"level":"info","ts":1666777006.730922,"caller":"log/logger.go:38","msg":"type/bc-break","count":10}
{"level":"warn","ts":1666777006.765012,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":1,"elapsedTime":0.034032667,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777007.129405,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":2,"elapsedTime":0.398426125,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777007.7629418,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":3,"elapsedTime":1.031964792,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777009.401166,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":4,"elapsedTime":2.670194458,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777013.4986129,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":5,"elapsedTime":6.767663167,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777020.106155,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":6,"elapsedTime":13.3752205,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777023.715019,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":7,"elapsedTime":16.984120417,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777028.1540868,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":8,"elapsedTime":21.423204667,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777032.2879229,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":9,"elapsedTime":25.557060958,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777037.169069,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":10,"elapsedTime":30.4382275,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777041.1225789,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":11,"elapsedTime":34.391753792,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777045.123342,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":12,"elapsedTime":38.392541708,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777051.056513,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":13,"elapsedTime":44.32573925,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777054.6858501,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":14,"elapsedTime":47.955094583,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777058.238646,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":15,"elapsedTime":51.507907583,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777062.5850592,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":16,"elapsedTime":55.85433475,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fbug\"","error":"GET request failed with 403"}
{"level":"warn","ts":1666777062.585168,"caller":"log/logger.go:42","msg":"failed to get issues count. Will be retried later","label":"type/bug","error":"GET request failed with 403"}
{"level":"warn","ts":1666777062.619869,"caller":"log/logger.go:42","msg":"error while processing the Github search. Retrying.","retryCount":1,"elapsedTime":0.034648709,"url":"https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-go+label:\"type%2Fdocs\"","error":"GET request failed with 403"}
{"level":"info","ts":1666777063.912952,"caller":"log/logger.go:38","msg":"type/docs","count":36}
{"level":"info","ts":1666777064.6252182,"caller":"log/logger.go:38","msg":"type/feature","count":62}
{"level":"info","ts":1666777065.251224,"caller":"log/logger.go:38","msg":"type/housekeeping","count":46}
{"level":"info","ts":1666777065.526626,"caller":"log/logger.go:38","msg":"type/question","count":6}
{"level":"info","ts":1666777065.8371048,"caller":"log/logger.go:38","msg":"type/security","count":5}
{"level":"info","ts":1666777066.095794,"caller":"log/logger.go:38","msg":"type/test-flakiness","count":3}
{"level":"info","ts":1666777066.4683468,"caller":"log/logger.go:38","msg":"type/test-improvement","count":12}
{"level":"info","ts":1666777066.4684732,"caller":"log/logger.go:38","msg":"Retrying failed labels","labels":["type/bug"]}
{"level":"info","ts":1666777067.092025,"caller":"log/logger.go:38","msg":"type/bug","count":39}
```

## Code Search

We'd like to get how many occurrences of a string are present across all Github repositories. For that we are going to perform an authenticated call to `https://cs.github.com/api/count`, which is still in _Technology Preview_ and there is no API available yet.

For that reason, we need to export the `GHM_GITHUB_COOKIE` environment variable, so that the code is able to pass it as an HTTP header to the HTTP request.

> To get your HTTP cookie with Github, simply use a sniffer (i.e Chrome's network tab), go to github.com and get the cookie from Headers section once you are logged in.

To use this command, please define a search query using this syntax: https://cs.github.com/about/syntax 

```shell
ghm cs --query "testcontainers-go v0.15.0 path:go.mod"
```

It will print the count of the files using the application formatter (default is console log):

```
{"level":"info","ts":1666777433.197243,"caller":"log/logger.go:38","msg":"Total files","count":66}
```
