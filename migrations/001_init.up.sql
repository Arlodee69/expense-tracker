BEGIN;

CREATE TABLE categories(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE CHECK (name != '')
);

CREATE TABLE expenses(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    categoryid BIGINT  NOT NULL REFERENCES categories(id),
    amount BIGINT NOT NULL CHECK (amount > 0),
    spenton DATE NOT NULL,
    description TEXT NOT NULL DEFAULT ''
        CHECK (char_length(description) <= 1000),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()    
);

COMMIT;