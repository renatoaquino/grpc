FROM scratch

WORKDIR /app
EXPOSE 8888
COPY server-linux /app/server

CMD ["/app/server"]