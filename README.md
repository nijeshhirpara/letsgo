# Let's Go - A sample GO app


## Overview

Let's Go is an HTTP server. It has various apis to play with. It is a small app that can group users of a company to multiple teams and provide access permissions to their work accordingly.
1. The users belong to a single company at any time.
2. A company can have multiple teams
3. One user can belong to multiple teams within the same company.
4. Any work done by a user under a certain team belongs to the team. It means that all users within the team will be able to see each other's work. Users belonging to other teams cannot see the work created under another team.

## Development Environment

* To set path variables on Linux run  ````. ./go.sh````
* Go src/letsgo and Copy .env.sample to .env. 

### build

```go build```

### run

```./letsgo```

### run unit tests

```go test ./...```
