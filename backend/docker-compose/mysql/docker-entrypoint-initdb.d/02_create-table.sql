use chat_rp_db;

create table games (
    id     int unsigned not null auto_increment comment 'ID',
    game_name   varchar(255) not null comment 'ゲーム名',
    created_at  datetime     not null comment '作成日時',
    updated_at  datetime     not null comment '更新日時',
    primary key (id)
);

create table players (
    id   int unsigned not null auto_increment comment 'ID',
    player_name varchar(255) not null comment 'プレイヤー名',
    created_at  datetime     not null comment '作成日時',
    updated_at  datetime     not null comment '更新日時',
    primary key (id)
);

create table participants (
    id int unsigned not null auto_increment comment 'ID',
    game_id        int unsigned not null,
    player_id      int unsigned not null,
    created_at  datetime     not null comment '作成日時',
    updated_at  datetime     not null comment '更新日時',
    primary key (id)
);

alter table participants
    add constraint fk_participants_games foreign key (game_id)
    references games (id)
    on update restrict
    on delete restrict
;    

alter table participants
    add constraint fk_participants_players foreign key (player_id)
    references players (id)
    on update restrict
    on delete restrict
;    

