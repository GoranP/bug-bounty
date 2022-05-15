FROM golang:alpine as builder

COPY ./ /src
WORKDIR /src
RUN apk add make curl bash
# build as described in readme.md
RUN make build

# final image production ready, without source code and with low privilege user running service  
FROM alpine:3.15.4

RUN mkdir /app
COPY --from=builder /src/bin/* /app
RUN chmod +x /app/*
WORKDIR /app

RUN adduser -D -h /app secure_user
RUN chown -R secure_user:secure_user /app

USER secure_user

EXPOSE 8080

ENTRYPOINT echo -e "\nadmin token: $(/app/token -admin true) \nuser token: $(/app/token)\n" && /app/innsecure