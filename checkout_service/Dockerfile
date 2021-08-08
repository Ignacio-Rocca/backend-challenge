FROM golang

## We create an /app directory within our image that will hold our application source files
RUN mkdir /app

## We copy everything in the root directory into our /app directory
ADD . /app

## we run go mod to download the dependencies & build to compile
WORKDIR /app
RUN go mod download
WORKDIR /app/cmd/api
RUN go build -o main .

## Our start command which kicks off our newly created binary executable
CMD ["/app/cmd/api/main"]