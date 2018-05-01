FROM golang:onbuild
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get github.com/labstack/echo
RUN go build -o main . 
EXPOSE 8000
CMD ["sh", "-c", "/app/main $GOAPPPORT $GOAPPDBUSER $GOAPPDBPASS $GOAPPDBHOST $GOAPPDBPORT"]
