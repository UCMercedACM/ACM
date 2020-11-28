# Global Argument to be shared across multiple stages
ARG NAME_OF_DIAGRAM=name

# ################## #
# Go-Diagram Library #
# ################## #
FROM golang:latest AS builder

# Author the maintainers
LABEL maintainer="UCM ACM Chapter"
LABEL maintainer.email="acm@ucmerced.edu"

# Get the name of the diagram in which to run
ARG DIAGRAM_NAME

# Set working directory
WORKDIR /go/src/app

# Install Packages
COPY go.mod .
COPY go.sum .
RUN go mod download

# If this folder exists then the build will crash
RUN rm -rf /go/src/app/go-diagrams

# Copy the necessary file to generate a diagram with
COPY ${DIAGRAM_NAME}.go ${DIAGRAM_NAME}.go

# Run the code from the necessary file
RUN go run ${DIAGRAM_NAME}.go

# ##################### #
# Graphviz Dot Compiler #
# ##################### #
FROM alpine:3.10 AS compiler

# Author the maintainers
LABEL maintainer="UCM ACM Chapter"
LABEL maintainer.email="acm@ucmerced.edu"

# Get the name of the  diagram in which will be generated
ARG DIAGRAM_NAME

# Generate directory to mount 
RUN mkdir /home/diagrams

# Install graphviz and create a user/group to work with it
RUN apk add --update --no-cache graphviz ttf-dejavu
RUN addgroup -g 1000 dot
RUN adduser -u 1000 -G dot -s /bin/sh -D dot

# Give user permission to write to folder
RUN chown -R 1000 /home

# Change user to the new allowed user
USER 1000

# Set the main directory
WORKDIR /home

# Copy necessary files from the builder stage to the compiler stage
COPY --from=builder /go/src/app/go-diagrams/assets /home/dot/assets
COPY --from=builder /go/src/app/go-diagrams/${DIAGRAM_NAME}.dot /home/dot/${DIAGRAM_NAME}.dot

# fixes "Warning: No such file or directory while opening" issue
WORKDIR /home/dot

CMD FILE="$(basename *.dot .dot)" && dot -Tpng ${FILE}.dot > /home/diagrams/${FILE}.png
