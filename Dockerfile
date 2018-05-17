FROM golang:onbuild

RUN mkdir -p /go/src/app && cd "$_"
COPY . .

RUN go get -d -v
RUN go install -v

#RUN go get github.com/kardianos/govendor
#RUN govendor init
#RUN govendor fetch github.com/gin-gonic/gin@v1.2
#RUN go build -o app

EXPOSE 9000

CMD ["app"]
