# Github Metrics

This project is a PoC to explore the Github API and retrieve certain metrics from a given repository.

It will perform certain API calls using HTTP requests and backoff strategies to mitigate rate limits in the Github API.

## Labels

We'd like to get, for each label, how many issues are labelled for the label.

```shell
ghm labels -O mdelapenya -R github-metrics
```

It will first get all the labels in the repository, and for each it will get the total number of issues, printing the result in console:

```
2022/10/17 12:14:18 Number of Issues for area/bitbucket-pipelines: 4
2022/10/17 12:14:18 Number of Issues for area/docker-compose: 85
2022/10/17 12:14:19 Number of Issues for area/java10: 0
2022/10/17 12:14:19 Number of Issues for area/java9: 4
2022/10/17 12:14:19 Number of Issues for area/logging: 3
2022/10/17 12:14:19 Number of Issues for area/okhttp: 5
2022/10/17 12:14:20 Number of Issues for area/reusable-containers: 12
2022/10/17 12:14:20 Number of Issues for area/shading: 26
2022/10/17 12:14:21 Number of Issues for area/test frameworks: 16
>>> [1][34.555458ms] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"blocker"". Retrying: GET request failed with 403
>>> [2][623.698666ms] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"blocker"". Retrying: GET request failed with 403
>>> [3][2.102354333s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"blocker"". Retrying: GET request failed with 403
>>> [4][4.467763291s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"blocker"". Retrying: GET request failed with 403
>>> [5][8.253982875s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"blocker"". Retrying: GET request failed with 403
2022/10/17 12:14:34 Number of Issues for blocker: 15
2022/10/17 12:14:35 Number of Issues for client/docker-for-mac: 19
2022/10/17 12:14:35 Number of Issues for client/docker-for-windows: 30
2022/10/17 12:14:36 Number of Issues for client/docker-machine: 16
2022/10/17 12:14:36 Number of Issues for client/in-container: 11
2022/10/17 12:14:36 Number of Issues for client/podman: 1
2022/10/17 12:14:37 Number of Issues for dependencies: 2485
2022/10/17 12:14:37 Number of Issues for don't merge: 1
2022/10/17 12:14:37 Number of Issues for github_actions: 70
2022/10/17 12:14:38 Number of Issues for good first issue: 54
>>> [1][34.670667ms] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [2][664.969917ms] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [3][1.266380333s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [4][2.616506s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [5][5.042094042s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [6][9.08359175s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [7][14.21232325s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [8][20.822407708s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [9][24.447099s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [10][28.889268292s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [11][33.018220125s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [12][37.902986667s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [13][41.856133625s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [14][45.882069125s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [15][51.820445417s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
>>> [16][55.450425292s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"gradle-wrapper"". Retrying: GET request failed with 403
2022/10/17 12:15:37 Number of Issues for gradle-wrapper: 22
2022/10/17 12:15:38 Number of Issues for hacktoberfest: 18
2022/10/17 12:15:38 Number of Issues for help wanted: 77
2022/10/17 12:15:39 Number of Issues for java: 1290
2022/10/17 12:15:39 Number of Issues for modules/azure: 16
2022/10/17 12:15:40 Number of Issues for modules/cassandra: 12
2022/10/17 12:15:40 Number of Issues for modules/clickhouse: 3
2022/10/17 12:15:40 Number of Issues for modules/cockroachdb: 10
2022/10/17 12:15:41 Number of Issues for modules/consul: 7
2022/10/17 12:15:42 Number of Issues for modules/couchbase: 42
>>> [1][35.041875ms] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [2][503.25725ms] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [3][1.612612083s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [4][4.374175041s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [5][7.582782666s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [6][11.606067083s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [7][17.906634458s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [8][21.476221791s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [9][28.34628175s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [10][34.366296583s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [11][39.523539291s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [12][42.206797666s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [13][45.536260541s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
>>> [14][51.1106055s] error while processing the Github search for "https://api.github.com/search/issues?q=repo:testcontainers/testcontainers-java+label:"modules%2Fdb2"". Retrying: GET request failed with 403
2022/10/17 12:16:42 Number of Issues for modules/db2: 2
```
