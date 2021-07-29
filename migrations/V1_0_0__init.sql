CREATE TABLE properties (
                              id serial,
                              application varchar(255) NOT NULL,
                              profile varchar(255) DEFAULT NULL,
                              label varchar(255) DEFAULT NULL,
                              key varchar(255) DEFAULT NULL,
                              value text,
                              PRIMARY KEY (id)
);
