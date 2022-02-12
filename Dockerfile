FROM golang:alpine
ENV GO111MODULE=on
ENV PKG_NAME=gitlab.com/s2.1-backend/shm-product-svc/
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
#ENV GITLAB_TOKEN=glpat-7j4ouAzWy_9NBEYj1VWE
ENV GOPRIVATE=gitlab.com/*
RUN apk update && apk upgrade
RUN apk add --no-cache git
#RUN git config --global url."https://it.shoesmart:47Pax8bptr7jN7Zeiny5@gitlab.com".insteadOf "https://gitlab.com"
#RUN git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com/"
RUN git config --global url."https://oauth2:glpat-7j4ouAzWy_9NBEYj1VWE@gitlab.com/".insteadOf "https://git@gitlab.com/"
WORKDIR $PKG_PATH
COPY . $PKG_PATH
COPY .netrc /root/.netrc
RUN echo $PWD
#RUN git config credential.helper store
RUN go mod vendor
WORKDIR $PKG_PATH/server/http
RUN echo $PWD
RUN go build main.go
EXPOSE 9004
CMD ["sh", "-c", "./main"]