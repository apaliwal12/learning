# ─────── MULTI-STAGE DOCKERFILE ───────
# This is the industry standard way to containerize Go applications.
# It results in incredibly small, secure Docker images (often < 15MB!).

# --- Stage 1: Build ---
# We use the official Go image only for COMPILING the code.
# This image is large (~800MB) because it has all the build tools.
FROM golang:1.23 AS builder

WORKDIR /app

# (Optional) Copy go.mod and go.sum and download dependencies first.
# This leverages Docker layer caching to speed up rebuilds.
# COPY go.mod go.sum ./
# RUN go mod download

COPY docker_project.go .

# Build the binary!
# CGO_ENABLED=0 tells Go to build a completely statically linked binary.
# This is CRITICAL because our final image (scratch) won't have any C libraries (like glibc).
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=v1.0.0-docker'" -o server docker_project.go


# --- Stage 2: Final Runtime Image ---
# 'scratch' is a special, empty Docker image. It has NO OS, NO shell, NO utilities.
# It is perfectly secure and tiny.
FROM scratch

# Copy the compiled binary from the 'builder' stage into this empty image.
COPY --from=builder /app/server /server

# Expose the port our app runs on
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/server"]
