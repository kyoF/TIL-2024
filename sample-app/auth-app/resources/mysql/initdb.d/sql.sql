CREATE TABLE `users`
(
  id bigint auto_increment,
  name varchar(255) NOT NULL,
  password varchar(255),
  PRIMARY KEY (`id`)
);

INSERT INTO users (`name`, `password`) VALUES ('山田太郎', 'yamadapassword'), ('鈴木花子', 'suzukipassword');
