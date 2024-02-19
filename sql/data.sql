insert into users (nameuser, nick, email, pw)
values
("User 1", "user_1", "user1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- User1
("User 2", "user_2", "user2@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- User2
("User 3", "user_3", "user3@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"); -- User3

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into posts(title, content, author_id)
values
("User post 1", "This is the User post 1! Great!", 1),
("User post 2", "This is the User post 2! Great!", 2),
("User post 3", "This is the User post 3! Great!", 3);