CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    tenant_id integer references tenant(id)
);


CREATE TABLE tenant {
    id SERIAL PRIMARY_KEY,
    name VARCHAR(100)
};
