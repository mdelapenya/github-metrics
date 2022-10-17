# Github Metrics

This project is a PoC to explore the Github API and retrieve certain metrics from a given repository.

It will perform certain API calls using HTTP requests and backoff strategies to mitigate rate limits in the Github API.

## Labels

We'd like to get, for each label, how many issues are labelled for the label.

```shell
ghm labels -O mdelapenya -R github-metrics
```

It will first get all the labels in the repository, and for each it will get the total number of issues, printing the result in console.
