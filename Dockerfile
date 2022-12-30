FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /moneyconverter

EXPOSE 3000

CMD [ "/moneyconverter" ]