FROM golang:latest

LABEL maintainer="UCM ACM Chapter"
LABEL maintainer.email="acm@ucmerced.edu"

WORKDIR /go/src/app

COPY . .

# Install Package
RUN go get github.com/blushft/go-diagrams

CMD ["go", "run", "architecture.go"]