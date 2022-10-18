FROM golang:1.17-alpine
WORKDIR /server

COPY . .
RUN go mod download & go mod verify
RUN chmod +rx /server

RUN go build -o /server

EXPOSE 8080
CMD [ "/server" ]



