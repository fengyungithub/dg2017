FROM previousnext/php:7.0-dev

RUN apt-get update && \
    apt-get -y install zlib1g-dev && \
    echo 'extension=grpc.so' >> /usr/local/etc/php/php.ini

RUN pecl install grpc
