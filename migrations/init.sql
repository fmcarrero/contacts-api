CREATE TABLE IF NOT EXISTS contacts (
    id numeric NOT NULL,
    full_name varchar(255) NOT NULL,
    phone_number varchar(100) NOT NULL,
    email varchar(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    update_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT contacts_pk PRIMARY KEY (id)
);
insert into contacts (id, full_name, phone_number, email, created_at,update_at )
values(
          1795068458442948608,
          'fran carrero',
          '+573143159054',
          'ma@gmail.com',
          '2024-06-01 10:00:00-05',
          '2024-06-01 12:00:00-05'
      );