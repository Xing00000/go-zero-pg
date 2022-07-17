CREATE TABLE `book`
(
  `book` varchar(255) NOT NULL ,
  `price` int NOT NULL ,
  PRIMARY KEY(`book`)
) ;

comment on column book.book is 'book name';
comment on column book.price is 'book price';
