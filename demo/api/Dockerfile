FROM       gliderlabs/alpine:3.1
MAINTAINER Nick Schuch

RUN apk --update add ca-certificates

ADD api /api
RUN chmod a+x /api

CMD ["/api"]
