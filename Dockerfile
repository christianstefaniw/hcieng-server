FROM golang

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

EXPOSE 8080

CMD ["./bin/nonacoin"]