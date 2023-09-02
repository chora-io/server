FROM postgres:15

# Set data directory
ENV PGDATA=/var/pgdata

# Create directory to store PostgreSQL data and logs
RUN mkdir -p ${PGDATA} /tmp /var/log/postgresql

# Change owner of directory to postgres group and user
RUN chown -R postgres:postgres ${PGDATA} /tmp /var/log/postgresql

# Change working directory
WORKDIR /home/postgres

# Copy db migrations
COPY db/migrations /home/postgres/migrations

# Copy init script
COPY docker/scripts/postgres_init.sh /home/postgres/scripts/

# Copy start script
COPY docker/scripts/postgres_start.sh /home/postgres/scripts/

# Make start script executable
RUN ["chmod", "+x", "/home/postgres/scripts/postgres_start.sh"]

# Set user to run init script and db container
USER postgres

# Run init script
RUN /home/postgres/scripts/postgres_init.sh
