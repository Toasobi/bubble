FROM golang:1.18-alpine



ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"


WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o /app/myapp .

RUN mkdir src .

RUN cp -r /src/app/static ./src/

EXPOSE 8080

CMD [ "/app/myapp" ]