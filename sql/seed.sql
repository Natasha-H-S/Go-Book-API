CREATE TABLE Books (
    ID int NOT NULL AUTO_INCREMENT,
    Title varchar(255) NOT NULL,
    Description varchar(255),
    Genre varchar(255),
    PRIMARY KEY (ID)
);

INSERT INTO Books (Title, Description, Genre)
VALUES ('Book 1', 'This is Book 1 from sql', 'Mystery');
