FROM golang:1.19-alpine3.16 as build

WORKDIR /go/src/app

ARG BRANCH

COPY go.mod go.sum .
RUN go mod download

COPY .. .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -o /go/bin/app -ldflags "-X 'main.Branch=$BRANCH'"

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/app /
CMD ["/app"]
