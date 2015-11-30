FROM scratch
COPY app /
ENTRYPOINT ["/app"]
EXPOSE 9000
