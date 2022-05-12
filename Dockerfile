FROM golang:1.18-alpine as builder

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -v -o /bin/roll cmd/roll/main.go


FROM scratch

COPY --from=builder /bin/roll /bin/roll

LABEL author="Britton Hayes"
LABEL github="https://github.com/brittonhayes/roll"

ENTRYPOINT ["/bin/roll"]

CMD [ "/bin/roll" ]