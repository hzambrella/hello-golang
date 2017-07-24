--  输入,show create table hz_test.user_info;
-- 右键copy field,将建表的sql复制过来
-- 保存建表语句。

-- 用户基本信息表
CREATE TABLE user_info (
    -- 用户id
      user_id INT(11) NOT NULL AUTO_INCREMENT,
    -- 密码
      password VARCHAR(45) NOT NULL,
    -- 用户名
      user_name VARCHAR(45) NOT NULL,
    -- 状态
        -- 1 正常
        -- 其他 封号
      status INT(11) NOT NULL,
    -- 创建时间
      create_time TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
      PRIMARY KEY (user_id),
      UNIQUE KEY user_name_UNIQUE (user_name)
) ENGINE=InnoDB
