CREATE TABLE chats(
    id SERIAL PRIMARY KEY,
    id_user_1 INTEGER,
    id_user_2 INTEGER
);

CREATE TABLE messages(
    id SERIAL PRIMARY KEY,
    id_chat INTEGER,
    id_sender INTEGER,
    message VARCHAR(500)
)