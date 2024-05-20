# Fixing docker image of Go
FROM golang:1.22

# Establishing work directory to simplify next commands
WORKDIR /app

# Copy go.mod into app
COPY go.mod .

# Download go modules and verify that dependencies of the main modules
#   cache have not been modified
RUN go mod download && go mod verify

# Download all local files into app
COPY . .

# Compile and run the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping .

# Entry command of the application (executed command to run the application)
CMD ["/docker-gs-ping"]