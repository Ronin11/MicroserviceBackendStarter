CREATE TABLE items(
   id                   SERIAL,
   created_time         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_weight       INT,
   updated_bp_systolic  INT,
   updated_bp_diastolic INT,
   updated_o2           INT,
   updated_bpm          INT,
   comment              VARCHAR (256),
   PRIMARY KEY (id)
)
