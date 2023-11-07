FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./
RUN mkdir /etc/server/
COPY ./etc/server/config.prod.toml /etc/server/config.toml
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/main.go


EXPOSE 8000

# Run
CMD ["/server", "-c",  "/etc/server/config.yaml"]