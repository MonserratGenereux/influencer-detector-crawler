# Influencer Detector
Influencer Detector is a system designed with the purpose of minning Facebook pages info and analyzing their relations in order to calculate influence levels within certain category over a predefined graph.

![alt text](https://github.com/dtoledo23/influencer-detector-front/blob/master/src/assets/img/Arquitectura.png?raw=true Influencer Detector Architecture)

- [influencer-detector-front](https://github.com/dtoledo23/influencer-detector-front)
- [influencer-detector-back](https://github.com/dtoledo23/influencer-detector-back)
- [influencer-detector-crawler](https://github.com/dtoledo23/influencer-detector-crawler)
- [influencer-detector-analytics](https://github.com/dtoledo23/influencer-detector-analytics)

### About us
We developed Influencer Detector as a school project in the Advanced Databases course. The team:

- Monserrat Genereux
- Victor Garcia
- Diego Toledo

# influencer-detector-crawler
Facebook Graph API mining. This module fetches data from the Graph API. It provides an API to enable requesting data over a POST request. It stores the results on Cassandra database. Go was chosen for this task to make the fetching process a lot faster by using goroutines and make concurrent calls to Facebook.

## Requirements
- Go 1.7
- Cassandra 3.0

# Setup
1. Run the cql scripts under `cassandra_init.cql` on your Cassandra instance

## How to run locally
1. Clone repo under `$GOPATH/src/github.com/dtoledo23`
2. You need a Facebook Page Access Token. Get one from https://developers.facebook.com/docs/pages/access-tokens
3. Setup environment variables. Create a `.env` file based on the example under `.env.example`
4. `go run server.go`

## How to deploy
1. The app is already dockerized. Make sure you have `git` and `docker` installed on your host server.
2. Create `.env` file with the configuration needed. Take `.env.example` format.
3. Build: `docker build -t influencer-detector-crawler .`
4. Run: `docker run -d -p 8000:8000 influencer-detector-crawler`
