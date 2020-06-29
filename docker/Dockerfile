FROM golang:1.14.4 

ADD . $GOPATH/src/
WORKDIR $GOPATH/src/
COPY . .

RUN go get -u "github.com/gin-gonic/contrib/cors"
RUN go get -u "github.com/codegangsta/gin"
RUN go get -u "github.com/lib/pq"
RUN go get -u "github.com/gorilla/mux"
RUN go mod init puppkitt.com/v1

CMD ["gin", "run", "main.go"]
