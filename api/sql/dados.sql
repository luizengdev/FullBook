insert into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario1", "usuario1@gmail.com", "$2a$10$bpxdStqO5YCi4r7/IItxkOlBUEDHLZS/G6baBkroV.TFM4fVSlUwa"), -- usuario1
("Usuário 2", "usuario2", "usuario2@gmail.com", "$2a$10$bpxdStqO5YCi4r7/IItxkOlBUEDHLZS/G6baBkroV.TFM4fVSlUwa"), -- usuario2
("Usuário 3", "usuario3", "usuario3@gmail.com", "$2a$10$bpxdStqO5YCi4r7/IItxkOlBUEDHLZS/G6baBkroV.TFM4fVSlUwa"); -- usuario3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do usuário 1", "Essa é a publicação do usuário 1! oba!", 1),
("Publicação do usuário 2", "Essa é a publicação do usuário 2! oba!", 2),
("Publicação do usuário 3", "Essa é a publicação do usuário 3! oba!", 3);