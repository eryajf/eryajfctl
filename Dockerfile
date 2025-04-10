FROM alpine:3.21

LABEL maintainer=eryajf

ENV TZ=Asia/Shanghai
ENV BINARY_NAME=eryajfctl

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
    && apk upgrade \
    && apk add bash curl wget alpine-conf busybox-extras tzdata \
    && apk del alpine-conf && rm -rf /var/cache/*

ARG TARGETOS
ARG TARGETARCH

COPY config.example.yml /etc/${BINARY_NAME}/config.yml
COPY bin/${BINARY_NAME}_${TARGETOS}_${TARGETARCH} /usr/local/bin/${BINARY_NAME}

RUN chmod +x /usr/local/bin/${BINARY_NAME}