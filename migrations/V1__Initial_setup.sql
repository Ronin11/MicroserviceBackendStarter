CREATE TABLE items(
   id                   SERIAL,
   created_time         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   data jsonb,
   PRIMARY KEY (id)
)
