FROM nginx:1.17.1-alpine

RUN apk upgrade --update \
    && apk add bash git

ADD entrypoint.sh /entrypoint.sh

CMD ["/entrypoint.sh"]
