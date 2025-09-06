-- 创建密码重置令牌表
CREATE TABLE IF NOT EXISTS password_reset_tokens (
    id bigserial PRIMARY KEY, -- 主键
    user_id bigint NOT NULL, -- 用户ID
    token text NOT NULL, -- 令牌
    expires_at timestamp NOT NULL, -- 过期时间
    is_used boolean NOT NULL DEFAULT false, -- 是否已使用
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- 添加索引以提高查询性能
CREATE INDEX idx_password_reset_tokens_user_id ON password_reset_tokens(user_id);
CREATE UNIQUE INDEX idx_password_reset_tokens_token ON password_reset_tokens(token);