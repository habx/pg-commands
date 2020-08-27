CREATE TABLE "test_1"
(
    id           serial       not null
        constraint test_1_pkey
            primary key,
    item        varchar(256) not null,
    db_from        varchar(256) not null
);

CREATE TABLE "test_2"
(
    id           serial       not null
        constraint test_2_pkey
            primary key,
    item        varchar(256) not null,
    db_from        varchar(256) not null
);

alter table "test_1"
    owner to "dev_example";

INSERT INTO "test_1" (item, db_from) VALUES
    ('Acnejbam', 'dev'),
    ('Jehiip', 'dev'),
    ('Jopperof', 'dev'),
    ('Gahoiz', 'dev'),
    ('Zilitbi', 'dev'),
    ('Usoibpoj', 'dev'),
    ('Domeovo', 'dev'),
    ('Recweze', 'dev'),
    ('Isiwza', 'dev'),
    ('Kugdoep', 'dev'),
    ('Eveajupuc', 'dev'),
    ('Tuatvi', 'dev');

alter table "test_2"
    owner to "dev_example";

INSERT INTO "test_2" (item, db_from) VALUES
    ('Acnejbam', 'dev'),
    ('Jehiip', 'dev'),
    ('Jopperof', 'dev'),
    ('Gahoiz', 'dev'),
    ('Zilitbi', 'dev'),
    ('Usoibpoj', 'dev'),
    ('Domeovo', 'dev'),
    ('Recweze', 'dev'),
    ('Isiwza', 'dev'),
    ('Kugdoep', 'dev'),
    ('Eveajupuc', 'dev'),
    ('Tuatvi', 'dev');