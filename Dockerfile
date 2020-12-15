FROM golang:1.14-alpine

ADD q /app/
ADD env/ /app/
WORKDIR /app

RUN adduser -D svcuser
USER svcuser

CMD [ "sh", "-c", "source /app/*.env; /app/q" ]