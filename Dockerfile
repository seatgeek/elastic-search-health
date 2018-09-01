FROM alpine:3.8

# Adding ca-certificates for external communication, and openssh
# for attaching to remote nodes
RUN apk add --update curl ca-certificates openssh-client && \
    rm -rf /var/cache/apk/*

ADD ./build/elastic-search-health-linux-amd64 /elastic-search-health
RUN chmod +x /elastic-search-health

ENTRYPOINT ["/elastic-search-health"]
