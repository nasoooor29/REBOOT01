# Multi stage building using Dockerfiles

# FROM golang:1.22.4-bookworm as build
# WORKDIR /src

# # Copy go.mod files first to leverage Docker's cache
# COPY go.mod /src
# RUN go mod download

# # Copy remaining files
# COPY . /src

# # Build the Go application
# RUN go build -o /src/ascii-art-web .

# LABEL org.reboot.title="ASCII Art Web Docker" \
#       org.reboot.description="A web application to generate ASCII art built with Go." \
#       org.reboot.version="1.0" \
#       org.reboot.authors="nhussain/yabuzuha/etarada" \
#       org.reboot.source="" \
#       org.reboot.created="" \
#       org.reboot.licenses=""

# CMD ["/src/ascii-art-web"]

# Stage 1: Build the Go application
FROM golang:1.22.4 AS build

# set the working directory in the container
WORKDIR /src

# Copy go.mod file first to leverage Docker's cache
# best practice if any dependencies are required for the build
COPY go.mod ./
RUN go mod download

# Copy the rest of the source files
COPY . .

# Build the Go application
# https://stackoverflow.com/questions/61515186/when-using-cgo-enabled-is-must-and-what-happens
RUN CGO_ENABLED=0 GOOS=linux go build -o ascii-art-web .
# RUN go build -o ascii-art-web .

# Stage 2: Create a smaller image for the final binary
# https://hub.docker.com/_/alpine/
FROM alpine:3.20

# Set up environment variables
ENV APP_HOME /app

# Create the application directory
WORKDIR $APP_HOME

# Copy the built binary from the build stage
COPY --from=build /src/ascii-art-web .

# Copy the static, templates, and fonts directories from the build stage
COPY --from=build /src/templates ./templates
COPY --from=build /src/fonts ./fonts
COPY --from=build /src/static ./static


# https://www.howtogeek.com/devops/how-and-why-to-add-health-checks-to-your-docker-containers/
# this will not work in scratch since the container does not have the curl binary
# HEALTHCHECK CMD curl --fail http://localhost:8080/ || exit 1

# Set the binary as the entry point
ENTRYPOINT ["./ascii-art-web"]

# Label the image
LABEL org.reboot.title="ASCII Art Web Docker" \
      org.reboot.description="A web application to generate ASCII art built with Go." \
      org.reboot.version="1.0" \
      org.reboot.authors="nhussain/yabuzuha/etarada" \
      org.reboot.source="https://learn.reboot01.com/git/etarada/ascii-art-web-dockerize" \
      org.reboot.created="26/June/2024" \
      org.reboot.licenses="MIT"