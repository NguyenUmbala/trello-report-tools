FROM golang
RUN mkdir /app 
ADD . /app/
WORKDIR /app
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/adlio/trello
RUN go get github.com/jinzhu/gorm/dialects/sqlite
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
#RUN go build -o main .
EXPOSE 3000
CMD ["./main"]