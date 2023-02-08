# syntax=docker/dockerfile:1
FROM golang:1.19
WORKDIR /vesta
COPY . /vesta
RUN ls
RUN go install ./cmd/vestad
CMD ["sh", "./scripts/start-node.sh"]
EXPOSE 26657
EXPOSE 26656
EXPOSE 26658
EXPOSE 1317
EXPOSE 6060
EXPOSE 9090

