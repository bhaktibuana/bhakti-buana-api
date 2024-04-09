FROM golang:1.21-alpine

# define work directory
WORKDIR /app

# install packages
COPY go.mod go.sum ./
RUN go mod download

# collect all required files
COPY . .

# compile app to binary file
RUN chmod +x build.sh
RUN sh build.sh

# serve the app
RUN chmod +x serve.sh

CMD ["sh", "serve.sh"]
