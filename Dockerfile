# The base go-image
FROM golang:1.19

# Set working directory
WORKDIR /app

# Copy all files from the current directory to the app directory
COPY . /app
COPY meta.json /

RUN git rev-parse HEAD > /git_hash

# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o /server

EXPOSE 80

# Run the server executable
CMD [ "/server" ]
