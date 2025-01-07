CREATE TABLE user_credentials (
    id serial PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

create table users (
id serial primary key,
firstname varchar(80),
lastname varchar(80),
phone_number varchar(18),
image varchar(255),
user_credentials_id INT UNIQUE NOT NULL,
created_at timestamp default now(),
updated_at timestamp,
FOREIGN KEY (user_credentials_id) REFERENCES user_credentials(id) ON DELETE CASCADE
);

create table movie (
id serial primary key,
title varchar(100),
image_movie varchar(255),
genre varchar(100),
release_date date,
duration time,
director varchar(255),
cast_actor varchar(255),
synopsis text,
created_at timestamp default now(),
updated_at timestamp
);

-- insert into movie (title, image_movie, genre, release_date, duration, director, cast_actor, synopsis) values
-- ('Spiderman', '54b640ef-0a35-448a-8938-e0122bff3067.jpg', 'action, advanture', '2025-01-01', '02:35:00', 'joko anwar', 'tom holland, mary jane, peter', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.');



create table cinema (
id serial primary key,
cinema_name varchar(255),
price INT,
location varchar(255),
date date,
time TIME,
created_at timestamp default now(),
updated_at timestamp
);

create table movie_cinema (
	id serial primary key,
	movie_id INT,
    cinema_id INT,
    FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE CASCADE,
    FOREIGN KEY (cinema_id) REFERENCES cinema(id) ON DELETE cascade
);

create table orders (
  id serial primary key,
  movie_cinema_id int, 
  quantity int,
  total_price int,
  date date,
  time time,
  created_at timestamp default now(),
  updated_at timestamp,
  FOREIGN KEY (movie_cinema_id) REFERENCES movie_cinema(id) ON DELETE CASCADE
);

-- insert into cinema (cinema_name, price, "location" , "date" , "time") values
-- ('hiflix, XXI', 20000, 'jakarta, bandung', '2025-01-01', '06:35:00'),
-- ('hiflix', 20000, 'jogja, bandung', '2025-01-01', '02:35:00'),
-- ('XXI', 20000, 'jakarta, maiun', '2025-02-01', '03:35:00'),
-- ('hiflix, cinemas', 20000, 'bogor', '2025-08-01', '05:35:00'),
-- ('cinemas XXI', 20000, 'semarang, solo', '2025-12-01', '01:35:00'),
-- ('cinemas', 20000, 'surabaya, malang', '2025-06-01', '01:35:00');

create table seats (
id serial primary key,
seat varchar(80)
);


INSERT INTO movie (title, image_movie, genre, release_date, duration, director, cast_actor, synopsis) VALUES
('Avengers: Endgame', 'bedccff7-f856-4817-a188-770934d6ba71.jpg', 'action, adventure', '2025-02-02', '02:01:01', 'James Cameron', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Iron Man', 'b00db012-1a27-43b3-895a-abd3f540362e.jpg', 'action, adventure, sci-fi', '2025-03-03', '02:02:02', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Thor: Ragnarok', '73f7576a-7a3f-4ee5-99cf-2ef8c35c4d27.jpg', 'action, adventure, sci-fi, comedy', '2025-04-04', '02:03:03', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Black Panther', '78ac36be-9b04-48e4-b4b0-567e3b009231.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-05-05', '02:04:04', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Captain Marvel', 'be09476f-d529-44b5-aaf3-f27e082023a3.jpg', 'action', '2025-06-06', '02:05:05', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Doctor Strange', '4c16491d-67b0-4b62-b9f8-0a269ea5ddcf.jpg', 'action, adventure', '2025-07-07', '02:06:06', 'James Cameron', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Ant-Man', '7b21a349-8201-491a-b49d-b2818cd9aee8.jpg', 'action, adventure, sci-fi', '2025-08-08', '02:07:07', 'Sam Raimi', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Guardians of the Galaxy', '89b39f75-d82a-4a96-9cb7-7c6ec235acd6.jpg', 'action, adventure, sci-fi, comedy', '2025-09-09', '02:08:08', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('The Incredible Hulk', '16225387-42fc-4cb2-a34d-13abd5e68e1b.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-10-10', '02:09:09', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Wonder Woman', '036b9d1b-5a53-4807-bc77-9db36a04c09c.jpg', 'action', '2025-11-11', '02:10:10', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('The Dark Knight', '8d33e70e-ba1d-41f3-8943-3ec4c2e3c44f.jpg', 'action, adventure', '2025-12-12', '02:11:11', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Batman Begins', 'a6b1fcf1-7702-473c-89ea-2973cac3e48e.jpg', 'action, adventure, sci-fi', '2025-01-13', '02:12:12', 'Sam Raimi', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Justice League', '9ca42f1e-5bee-4e1f-a2b9-3a905abacc1d.jpg', 'action, adventure, sci-fi, comedy', '2025-02-14', '02:13:13', 'Christopher Nolan', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Aquaman', '45dd7d9a-d738-4426-bfc8-c2c6a8414dc0.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-03-15', '02:14:14', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('The Flash', '92024a91-7643-4587-857b-30107a529a45.jpg', 'action', '2025-04-16', '02:15:15', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Green Lantern', 'df97629c-020f-4260-851b-b3294e640379.jpg', 'action, adventure', '2025-05-17', '02:16:16', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Man of Steel', 'c78fbfda-77ae-44a3-bb9b-75fe911c97e5.jpg', 'action, adventure, sci-fi', '2025-06-18', '02:17:17', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Suicide Squad', '5bbef21a-4b39-4c5d-9057-1ca0a080ae2b.jpg', 'action, adventure, sci-fi, comedy', '2025-07-19', '02:18:18', 'Christopher Nolan', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Birds of Prey', 'f23dd399-50b7-44ac-a2a8-6665051fd2a7.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-08-20', '02:19:19', 'Steven Spielberg', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Deadpool', '937cb302-bc4f-47f1-96b7-d1f8b264c63c.jpg', 'action', '2025-09-21', '02:20:20', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Logan', '36b0bcbc-0263-4808-8aa5-701c7065989d.jpg', 'action, adventure', '2025-10-22', '02:21:21', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('X-Men: Days of Future Past', '68e4cece-ea7b-4c7a-907f-96a9d05a63f3.jpg', 'action, adventure, sci-fi', '2025-11-23', '02:22:22', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper '),
('X-Men: Apocalypse', 'f8e043bf-7cc9-4bed-8362-3171220172ae.jpg', 'action, adventure, sci-fi, comedy', '2025-12-24', '02:23:23', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Fantastic Four', 'ee5abaf4-31c4-4397-b331-00b78d19a3c1.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-01-25', '02:24:24', 'Steven Spielberg', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Blade', 'cfd2c0c7-84b0-420a-af03-8138d2d1b77c.jpg', 'action', '2025-02-26', '02:25:25', 'Joko Anwar', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi'),
('Kick-Ass', 'b8021ba7-3800-43fd-aadd-55569fcf0a12.jpg', 'action, adventure', '2025-03-27', '02:26:26', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Hellboy', '5995f46f-512f-4a54-bdb4-1efc25bd7423.jpg', 'action, adventure, sci-fi', '2025-04-28', '02:27:27', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('The Mask', '61a4d7a4-9d4c-4aa6-9108-358dd9bd4026.jpg', 'action, adventure, sci-fi, comedy', '2025-05-01', '02:28:28', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('The Crow', 'f62ee781-2083-48d3-9993-e70fe9409238.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-06-02', '02:29:29', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Hancock', '27e9a303-eef4-4aab-88ba-841c1b00344a.jpg', 'action', '2025-07-03', '02:30:30', 'Joko Anwar', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Supergirl', '1f4b17ef-62d2-4c3f-80e7-0d3ac64e4b3a.jpg', 'action, adventure', '2025-08-04', '02:31:31', 'James Cameron', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Captain America: The First Avenger', 'f4cecadd-7a01-4c62-843d-8a5b5215adf6.jpg', 'action, adventure, sci-fi', '2025-09-05', '02:32:32', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Captain America: The Winter Soldier', '98774490-c41e-4875-bbee-d1a0b703158a.jpg', 'action, adventure, sci-fi, comedy', '2025-10-06', '02:33:33', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Captain America: Civil War', '237e9924-f156-4c93-a5c8-16af35e89137.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-11-07', '02:34:34', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Thor: Love and Thunder', 'bdb06abf-f0ca-4831-85c1-c55414eacbcc.jpg', 'action', '2025-12-08', '02:35:35', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Eternals', 'a27e3b7c-fac1-4f2e-9cbe-5788b1f58b33.jpg', 'action, adventure', '2025-01-09', '02:36:36', 'James Cameron', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Shang-Chi and the Legend of the Ten Rings', '8cfeba89-c8b7-4c24-bfe2-76157cbdb6b5.jpg', 'action, adventure, sci-fi', '2025-02-10', '02:37:37', 'Sam Raimi', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi'),
('Spider-Man: Far From Home', '944b3b79-4355-4f89-a396-be36eb92c44c.jpg', 'action, adventure, sci-fi, comedy', '2025-03-11', '02:38:38', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Spider-Man: No Way Home', '31b1038a-4f27-4713-ba70-75f985c5e948.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-04-12', '02:39:39', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Venom', '5ed902cb-dc07-48ed-a933-2455e5858855.jpg', 'action', '2025-05-13', '02:40:40', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Venom: Let There Be Carnage', '9fac11f2-58d5-4ad0-9eae-185133df360b.jpg', 'action, adventure', '2025-06-14', '02:41:41', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Morbius', '1b54de99-99e5-454a-b7f4-27bcc8a319fb.jpg', 'action, adventure, sci-fi', '2025-07-15', '02:42:42', 'Sam Raimi', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi'),
('Daredevil', '79419977-8d87-40a9-af55-ca019f82fad6.jpg', 'action, adventure, sci-fi, comedy', '2025-08-16', '02:43:43', 'Christopher Nolan', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi'),
('Elektra', 'f7653c92-ba4e-4cf1-bbc6-7378456a16c5.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-09-17', '02:44:44', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nis'),
('Ghost Rider', 'f2655ea0-adbe-42a2-8668-bc0592b9435c.jpg', 'action', '2025-10-18', '02:45:45', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Ghost Rider: Spirit of Vengeance', 'b359d575-b79f-4066-8ce8-97ec9aed81a8.jpg', 'action, adventure', '2025-11-19', '02:46:46', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Punisher', 'cdcb9d17-e661-4bef-85b9-d0429a4cc82a.jpg', 'action, adventure, sci-fi', '2025-12-20', '02:47:47', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Punisher: War Zone', '5052331c-5718-476f-8924-6a07661ac4da.jpg', 'action, adventure, sci-fi, comedy', '2025-01-21', '02:48:48', 'Christopher Nolan', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Hulk', 'e898d801-54ec-4b77-aa45-83d0da995d84.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-02-22', '02:49:49', 'Steven Spielberg', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Iron Man 2', 'fd31b4e2-d3b9-47aa-9aca-13c2080e1b84.jpg', 'action', '2025-03-23', '02:50:50', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Iron Man 3', 'cd4d05c0-b48e-4bd9-839a-3b9071d7d180.jpg', 'action, adventure', '2025-04-24', '02:51:51', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('The Avengers', 'c3288379-b36b-4f09-a05d-8764da9511f0.jpg', 'action, adventure, sci-fi', '2025-05-25', '02:52:52', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Avengers: Age of Ultron', '8213e75c-daf9-4a78-88d3-6a8e78b2ab6a.jpg', 'action, adventure, sci-fi, comedy', '2025-06-26', '02:53:53', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Avengers: Infinity War', 'b3ee10b1-c77c-4d06-a691-1294c75703e8.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-07-27', '02:54:54', 'Steven Spielberg', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Thor', '5b1fb98c-bb10-4c7b-9a73-e4708edfb0b8.jpg', 'action', '2025-08-28', '02:55:55', 'Joko Anwar', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Thor: The Dark World', '0c2538db-c95e-465c-90f3-aafe8e270620.jpg', 'action, adventure', '2025-09-01', '02:56:56', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Black Widow', 'bba243f7-a5ba-4f37-b862-9888513999d1.jpg', 'action, adventure, sci-fi', '2025-10-02', '02:57:57', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Hawkeye', '96be7267-56f5-4477-aa21-7588c88b4350.jpg', 'action, adventure, sci-fi, comedy', '2025-11-03', '02:58:58', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Loki', '28b82c9b-bd34-46d2-9dab-6214723effe4.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-12-04', '02:00:00', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Moon Knight', '76fe7f10-be7b-4658-929d-6adce92907ff.jpg', 'action', '2025-01-05', '02:01:01', 'Joko Anwar', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('She-Hulk', '25526e32-8723-46b5-baab-46a8fcf98f89.jpg', 'action, adventure', '2025-02-06', '02:02:02', 'James Cameron', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Ms. Marvel', 'a5c85a35-4105-488f-b7b1-668a67fce079.jpg', 'action, adventure, sci-fi', '2025-03-07', '02:03:03', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('The Marvels', '94cb7e39-ee94-466d-bd86-27418b7d0bac.jpg', 'action, adventure, sci-fi, comedy', '2025-04-08', '02:04:04', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Inhumans', 'ee6aed67-7ee1-475d-b505-e1ad369dd9fe.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-05-09', '02:05:05', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('The Defenders', 'a6d43996-f45f-45cf-863f-bd0e9882fed3.jpg', 'action', '2025-06-10', '02:06:06', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Jessica Jones', '9cf5e296-a59a-42d5-9529-43a6255f8ff4.jpg', 'action, adventure', '2025-07-11', '02:07:07', 'James Cameron', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Luke Cage', 'a17a28ae-7f08-4c4d-bad6-6b323fb93183.jpg', 'action, adventure, sci-fi', '2025-08-12', '02:08:08', 'Sam Raimi', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Iron Fist', '8da994cf-5ba5-4ae5-95d5-89eccc958be2.jpg', 'action, adventure, sci-fi, comedy', '2025-09-13', '02:09:09', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('The Punisher', 'a96987b1-6095-4ed1-9864-dc293904bb28.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-10-14', '02:10:10', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Runaways', '8f6d5268-39cf-418d-a6b1-c6bd67fdb299.jpg', 'action', '2025-11-15', '02:11:11', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Cloak & Dagger', '7ce44053-b13e-46cc-bd07-181c78652d0c.jpg', 'action, adventure', '2025-12-16', '02:12:12', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Agents of S.H.I.E.L.D.', 'edc46d76-e4a0-4516-b73e-bdf3a442fb31.jpg', 'action, adventure, sci-fi', '2025-01-17', '02:13:13', 'Sam Raimi', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. .'),
('Agent Carter', '3378bea8-2671-4bb9-9293-220116f8959f.jpg', 'action, adventure, sci-fi, comedy', '2025-02-18', '02:14:14', 'Christopher Nolan', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('What If...?', 'd7d5805c-b939-4c20-8db0-5cc72086515d.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-03-19', '02:15:15', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('The Boys', '20b817c1-9a4b-4093-9eb3-877522985927.jpg', 'action', '2025-04-20', '02:16:16', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Invincible', '7c80d602-cb52-48bb-8e5e-c4dd0a29804b.jpg', 'action, adventure', '2025-05-21', '02:17:17', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Jupiters Legacy', '638b5c03-c27a-4a30-91b3-f140a1219ae0.jpg', 'action, adventure, sci-fi', '2025-06-22', '02:18:18', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Umbrella Academy', 'a07ec37b-7750-4138-9b1c-6f81c7117f1d.jpg', 'action, adventure, sci-fi, comedy', '2025-07-23', '02:19:19', 'Christopher Nolan', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Doom Patrol', '2751aeaf-cf3c-448a-96ae-8e29c369a280.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-08-24', '02:20:20', 'Steven Spielberg', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Titans', '1b0a731e-e693-4719-83eb-7843846b16b7.jpg', 'action', '2025-09-25', '02:21:21', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Young Justice', '3c085ec2-6cdd-492e-af51-063cb02c5d55.jpg', 'action, adventure', '2025-10-26', '02:22:22', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Teen Titans', '888dbf82-ee8c-45ab-ac6b-9530434977ae.jpg', 'action, adventure, sci-fi', '2025-11-27', '02:23:23', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Harley Quinn', 'c8dd94be-aecd-49e5-8486-32bcedf7a014.jpg', 'action, adventure, sci-fi, comedy', '2025-12-28', '02:24:24', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('The Lego Batman Movie', '7ebd063f-56ff-4004-b15d-9ba8fd52552e.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-01-01', '02:25:25', 'Steven Spielberg', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Spider-Man: Into the Spider-Verse', '24e22a49-9550-4e73-82aa-4714b8ffe536.jpg', 'action', '2025-02-02', '02:26:26', 'Joko Anwar', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Spider-Man: Across the Spider-Verse', '6171afad-9993-405c-b7ab-9b93d6665cf5.jpg', 'action, adventure', '2025-03-03', '02:27:27', 'James Cameron', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Big Hero 6', '90b0cb92-c286-47f8-be59-7bb9ff6f4032.jpg', 'action, adventure, sci-fi', '2025-04-04', '02:28:28', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Incredibles', '2eba5612-4375-480d-a2b5-816514a0398c.jpg', 'action, adventure, sci-fi, comedy', '2025-05-05', '02:29:29', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Incredibles 2', 'b389060e-9ff0-4e46-882c-0928d7c2f5fd.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-06-06', '02:30:30', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Megamind', 'bfc8f233-be21-475b-b14c-98e31cd69eae.jpg', 'action', '2025-07-07', '02:31:31', 'Joko Anwar', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Watchmen', '04af79bc-c2a2-41a4-9524-2c033e763709.jpg', 'action, adventure', '2025-08-08', '02:32:32', 'James Cameron', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Sin City', '6e7703da-a65b-43d6-a40f-e6702ef065d1.jpg', 'action, adventure, sci-fi', '2025-09-09', '02:33:33', 'Sam Raimi', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. .'),
('300', 'e0c95b63-67e0-4620-a7dc-12db3d62c103.jpg', 'action, adventure, sci-fi, comedy', '2025-10-10', '02:34:34', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('V for Vendetta', '9743c1ea-d9ce-4c6b-8693-1057e58ed535.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-11-11', '02:35:35', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Kick-Ass 2', '6556f269-775a-43fc-99fb-daf4a9c5a8d1.jpg', 'action', '2025-12-12', '02:36:36', 'Joko Anwar', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya, Emma Stone, Kirsten Dunst', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Scott Pilgrim vs. The World', '0eba38ec-4429-4a0c-82a8-60e8e2037ec3.jpg', 'action, adventure', '2025-01-13', '02:37:37', 'James Cameron', 'Tom Holland', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. '),
('Kingsman: The Secret Service', '3b5afa62-e5f8-4656-ad07-c1f1f0adf02e.jpg', 'action, adventure, sci-fi', '2025-02-14', '02:38:38', 'Sam Raimi', 'Tom Holland, Andrew Garfield', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Kingsman: The Golden Circle', 'e4459f47-7610-42d8-a529-495c177f9002.jpg', 'action, adventure, sci-fi, comedy', '2025-03-15', '02:39:39', 'Christopher Nolan', 'Tom Holland, Andrew Garfield, Tobey Maguire', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi.'),
('Spiderman', '54b640ef-0a35-448a-8938-e0122bff3067.jpg', 'action, advanture', '2025-01-01', '02:35:00', 'joko anwar', 'tom holland, mary jane, peter', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.'),
('The Kings Man', 'b0500d71-5650-475e-98a2-be1c8070c6dd.jpg', 'action, adventure, sci-fi, comedy, drama', '2025-04-16', '02:40:40', 'Steven Spielberg', 'Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus elementum semper nisi. ');


INSERT INTO cinema (cinema_name, price, "location" , "date" , "time") 
VALUES 
('ebu.id', 20000, 'jakarta', '2025-01-01', '06:35:00'),
('ebu.id', 20000, 'jakarta', '2025-01-01', '08:35:00'),
('ebu.id', 20000, 'jakarta', '2025-01-01', '11:35:00'),
('ebu.id', 20000, 'jakarta', '2025-01-01', '15:35:00'),
('hiflix', 20000, 'jakarta', '2025-01-01', '06:35:00'),
('hiflix', 20000, 'jakarta', '2025-01-01', '08:35:00'),
('hiflix', 20000, 'jakarta', '2025-01-01', '11:35:00'),
('cineone21', 20000, 'jakarta', '2025-01-01', '06:35:00'),
('cineone21', 20000, 'jakarta', '2025-01-01', '08:35:00'),
('cineone21', 20000, 'jakarta', '2025-01-01', '11:35:00'),
('ebu.id', 20000, 'bandung', '2025-01-02', '07:00:00'),
('ebu.id', 20000, 'bandung', '2025-01-02', '09:00:00'),
('ebu.id', 20000, 'bandung', '2025-01-02', '12:00:00'),
('ebu.id', 20000, 'bandung', '2025-01-02', '16:00:00'),
('hiflix', 20000, 'jogja', '2025-01-03', '06:30:00'),
('hiflix', 20000, 'jogja', '2025-01-03', '08:30:00'),
('hiflix', 20000, 'jogja', '2025-01-03', '10:30:00'),
('cineone21', 20000, 'jogja', '2025-01-03', '06:30:00'),
('cineone21', 20000, 'jogja', '2025-01-03', '08:30:00'),
('cineone21', 20000, 'jogja', '2025-01-03', '10:30:00'),
('ebu.id', 20000, 'surabaya', '2025-01-04', '07:30:00'),
('ebu.id', 20000, 'surabaya', '2025-01-04', '09:30:00'),
('ebu.id', 20000, 'surabaya', '2025-01-04', '12:30:00'),
('ebu.id', 20000, 'surabaya', '2025-01-04', '15:30:00'),
('hiflix', 20000, 'surabaya', '2025-01-04', '07:30:00'),
('hiflix', 20000, 'surabaya', '2025-01-04', '09:30:00'),
('hiflix', 20000, 'surabaya', '2025-01-04', '12:30:00');

select id from movie limit 10;
select id from cinema limit 10;

insert into movie_cinema (movie_id , cinema_id) values
(1, 3),
(2, 5),
(3, 5),
(4, 2),
(5, 3),
(6, 2),
(7, 1),
(8, 6),
(9, 6),
(10, 5),
(11, 3),
(12, 2),
(13, 1),
(14, 2),
(15, 5),
(16, 5),
(17, 1),
(18, 2),
(19, 2),
(20, 4),
(21, 1),
(22, 6),
(23, 4),
(24, 3),
(25, 3),
(26, 5),
(27, 3),
(28, 2),
(29, 3),
(30, 4),
(31, 2),
(32, 2),
(33, 2),
(34, 3),
(35, 3),
(36, 3),
(37, 5),
(38, 2),
(39, 3),
(40, 6),
(41, 4),
(42, 4),
(43, 2),
(44, 5),
(45, 2),
(46, 5),
(47, 5),
(48, 4),
(49, 6),
(50, 5),
(51, 5),
(52, 4),
(53, 5),
(54, 4),
(55, 3),
(56, 2),
(57, 1),
(58, 6),
(59, 3),
(60, 2),
(61, 2),
(62, 5),
(63, 2),
(64, 1),
(65, 5),
(66, 4),
(67, 3),
(68, 3),
(69, 3),
(70, 5),
(71, 4),
(72, 4),
(73, 2),
(74, 6),
(75, 5),
(76, 3),
(77, 6),
(78, 5),
(79, 4),
(80, 3),
(81, 2),
(82, 2),
(83, 2),
(84, 4),
(85, 3),
(86, 6),
(87, 4),
(88, 5),
(89, 3),
(90, 4),
(91, 3),
(92, 1),
(93, 4),
(94, 4),
(95, 1),
(96, 4),
(97, 4),
(98, 2),
(99, 2),
(100, 2);



INSERT INTO seats (seat)
VALUES
('A1'), ('A2'), ('A3'), ('A4'), ('A5'), ('A6'), ('A7'), ('A8'), ('A9'), ('A10'), ('A11'), ('A12'), ('A13'), ('A14'),
('B1'), ('B2'), ('B3'), ('B4'), ('B5'), ('B6'), ('B7'), ('B8'), ('B9'), ('B10'), ('B11'), ('B12'), ('B13'), ('B14'),
('C1'), ('C2'), ('C3'), ('C4'), ('C5'), ('C6'), ('C7'), ('C8'), ('C9'), ('C10'), ('C11'), ('C12'), ('C13'), ('C14'),
('D1'), ('D2'), ('D3'), ('D4'), ('D5'), ('D6'), ('D7'), ('D8'), ('D9'), ('D10'), ('D11'), ('D12'), ('D13'), ('D14'),
('E1'), ('E2'), ('E3'), ('E4'), ('E5'), ('E6'), ('E7'), ('E8'), ('E9'), ('E10'), ('E11'), ('E12'), ('E13'), ('E14'),
('F1'), ('F2'), ('F3'), ('F4'), ('F5'), ('F6'), ('F7'), ('F8'), ('F9'), ('F10'), ('F11'), ('F12'), ('F13'), ('F14'),
('G1'), ('G2'), ('G3'), ('G4'), ('G5'), ('G6'), ('G7'), ('G8'), ('G9'), ('G10'), ('G11'), ('G12'), ('G13'), ('G14');






















		
