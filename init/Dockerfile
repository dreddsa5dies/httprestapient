FROM postgres:13.3

ENV POSTGRES_USER=muser
ENV POSTGRES_PASSWORD=12345678
ENV POSTGRES_DB=matrix_attack

COPY *.sql /docker-entrypoint-initdb.d/