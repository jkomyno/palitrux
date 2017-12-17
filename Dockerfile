# smallest possible Docker image
FROM scratch
ENV PORT 8080
EXPORT $PORT

# copy only the executable file and the config file
COPY palitrux palitrux
COPY config.json config.json

CMD ["./palitrux"]