FROM alpine
EXPOSE 80
ENV PORT 80
ENTRYPOINT ["/http-pipe"]
ADD http-pipe /

