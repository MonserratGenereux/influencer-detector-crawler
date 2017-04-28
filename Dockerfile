FROM golang:1.7.4
MAINTAINER toledo.diego09@gmail.com

RUN mkdir -p $GOPATH/src/github.com/wizeline/wizeline-bot-analytics-client-api
WORKDIR $GOPATH/src/github.com/wizeline/wizeline-bot-analytics-client-api
COPY . .
RUN go get ./...

EXPOSE 8000

RUN chmod +x start-services.sh
CMD ["./start-services.sh"]