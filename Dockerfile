# Start from golang v1.12 base image
FROM golang:1.12

# Set current working directory
WORKDIR $GOPATH/src/TrelloReportTools

# Copy everything in current working directory
COPY . .

# Download all dependencies
RUN go get github.com/gin-gonic/gin
RUN go get github.com/adlio/trello
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/sqlite

# Set port
EXPOSE 3000

#RUN go build -o main .
CMD ["trello-report-tools"]