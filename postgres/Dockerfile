FROM postgres:16-alpine

# Copy the initialization scripts into the container
COPY ./init-scripts /docker-entrypoint-initdb.d/

ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=adminpassword
ENV POSTGRES_DB=maildatabase

EXPOSE 5432
