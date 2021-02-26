FROM alpine:3.12
MAINTAINER lico
# Install base packages
RUN apk update && apk add curl bash tree tzdata \
    && cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime
COPY lu-short /
RUN chmod u+x /lu-short
EXPOSE 80
CMD ["/lu-short"]
