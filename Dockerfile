FROM alpine
COPY ./sample-build /bin/
ENTRYPOINT "/bin/sample-build"
