INSERT INTO users(name, nickname, email, password) VALUES 
('user 1', 'user_1', 'user1@gmail.com', '$2a$10$aLda5n3njjhlYDb3qSPPF.CDAhpINU2E3DzgnsMD675b7.1SiQ4e6'),
('user 2', 'user_2', 'user2@gmail.com', '$2a$10$aLda5n3njjhlYDb3qSPPF.CDAhpINU2E3DzgnsMD675b7.1SiQ4e6'),
('user 3', 'user_3', 'user3@gmail.com', '$2a$10$aLda5n3njjhlYDb3qSPPF.CDAhpINU2E3DzgnsMD675b7.1SiQ4e6'),
('user 4', 'user_4', 'user4@gmail.com', '$2a$10$aLda5n3njjhlYDb3qSPPF.CDAhpINU2E3DzgnsMD675b7.1SiQ4e6'),
('user 5', 'user_5', 'user5@gmail.com', '$2a$10$aLda5n3njjhlYDb3qSPPF.CDAhpINU2E3DzgnsMD675b7.1SiQ4e6');

INSERT INTO followers(user_id, follow_id) VALUES 
(1, 2),
(1, 3),
(2, 1),
(2, 4),
(3, 5),
(3, 2),
(4, 1),
(4, 3),
(5, 4),
(5, 1);

INSERT INTO publications(title, content, user_id) VALUES
('Publicação usuário 1', 'Conteúdo de publicação do usuário 1', 1),
('Publicação usuário 2', 'Conteúdo de publicação do usuário 2', 2),
('Publicação usuário 3', 'Conteúdo de publicação do usuário 3', 3),
('Publicação usuário 4', 'Conteúdo de publicação do usuário 4', 4),
('Publicação usuário 5', 'Conteúdo de publicação do usuário 5', 5);