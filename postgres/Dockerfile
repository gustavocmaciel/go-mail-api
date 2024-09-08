# Postgres image for the init-container, that usually would be run just once
FROM postgres:16-alpine

# Copy the initialization scripts into the container
COPY ./init-scripts /docker-entrypoint-initdb.d/

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=adminpassword
ENV POSTGRES_DB=maildatabase

# Expose PostgreSQL port
EXPOSE 5432
