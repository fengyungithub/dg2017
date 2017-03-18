FROM       gliderlabs/alpine:3.1
MAINTAINER Nick Schuch

RUN apk --update add ca-certificates

ADD api /usr/local/bin/api
RUN chmod a+x /usr/local/bin/api

ADD cli /usr/local/bin/cli
RUN chmod a+x /usr/local/bin/cli

CMD ["api"]
