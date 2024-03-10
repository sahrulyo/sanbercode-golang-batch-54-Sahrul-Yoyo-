-- +migrate Up
-- +migrate StatementBegin


CREATE TABLE books (
    id          SERIAL PRIMARY KEY,
    Title       VARCHAR(255) NOT NULL,
    Description TEXT,
    ImageURL    VARCHAR(255),
    ReleaseYear INT,
    Price       DECIMAL(10, 2), 
    TotalPage   INT,
    Thickness   VARCHAR(50),
    CreatedAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CategoryID  INT,
    FOREIGN KEY (CategoryID) REFERENCES categories(id) 
);


CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd