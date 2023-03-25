CREATE TABLE book (
  isbn VARCHAR(13) NOT NULL PRIMARY KEY,
  title VARCHAR(50) NOT NULL,
  author VARCHAR(50),
  published DATETIME,
  genre VARCHAR(50)
);

CREATE TABLE collection (
  name VARCHAR(50) NOT NULL,
  isbn VARCHAR(50) NOT NULL,
  FOREIGN KEY (isbn) REFERENCES books (isbn)
  UNIQUE (name, isbn) ON CONFLICT IGNORE
);

