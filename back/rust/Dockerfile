
# Build stage
FROM rust:1.69-buster as builder

WORKDIR /app

# accept the build argument
ARG DATABASE_URL

ENV DATABASE_URL=$DATABASE_URL

COPY . .

RUN cargo build --release

# Production stage
FROM debian:buster-slim

WORKDIR /usr/local/bin

EXPOSE 8060

COPY --from=builder /app/target/release/rust .

CMD ["./rust"]