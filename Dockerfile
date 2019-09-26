FROM golang:latest

ENV GOPATH /workspace
ENV PATH "$PATH:$GOPATH/bin"

COPY ./ $GOPATH

WORKDIR $GOPATH

RUN make --no-print-directory install

CMD make --no-print-directory run
