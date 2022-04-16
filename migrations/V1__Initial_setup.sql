CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE items(
   id             uuid DEFAULT uuid_generate_v4 (),
   user_id        uuid,
   created_time   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   data           jsonb,
   PRIMARY KEY (id)
)
