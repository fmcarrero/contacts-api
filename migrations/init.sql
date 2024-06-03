CREATE TABLE IF NOT EXISTS contacts (
                                        id numeric NOT NULL CHECK (id > 0),
                                        full_name varchar(255) NOT NULL CHECK (COALESCE(NULLIF(full_name, ''), 'non-empty') <> 'non-empty'),
                                        phone_number varchar(100) NOT NULL CHECK (COALESCE(NULLIF(phone_number, ''), 'non-empty') <> 'non-empty'),
                                        email varchar(255) NOT NULL CHECK (COALESCE(NULLIF(email, ''), 'non-empty') <> 'non-empty'),
                                        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        update_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        CONSTRAINT contacts_pk PRIMARY KEY (id)
);
insert into contacts (id, full_name, phone_number, email, created_at,update_at )
values(
          1795068458442948608,
          'fran carrero',
          '+573143159054',
          'ma@gmail.com',
          now(),
          now()
      );
