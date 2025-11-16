FROM postgres:17-alpine
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=pullRequests
COPY init.sql /docker-entrypoint-initdb.d/10-init.sql
EXPOSE 80