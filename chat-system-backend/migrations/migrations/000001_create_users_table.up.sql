-- 创建用户表（PostgreSQL版本）
CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY, -- PostgreSQL的自增主键
  
  -- 基本信息字段
  username varchar(50) NOT NULL, -- 用户名
  nickname varchar(50) DEFAULT '' NOT NULL, -- 昵称
  avatar varchar(255) DEFAULT '' NOT NULL, -- 头像URL
  bio text DEFAULT '' NOT NULL, -- 个人简介
  
  -- 认证相关字段
  password varchar(255) NOT NULL, -- 密码哈希值
  salt varchar(50) NOT NULL, -- 密码盐值
  
  -- 预留的手机号和邮箱字段（在PostgreSQL中，NULL值不参与唯一约束）
  phone varchar(20) UNIQUE, -- 手机号，唯一约束
  email varchar(100) UNIQUE, -- 邮箱，唯一约束
  phone_verified boolean DEFAULT false NOT NULL, -- 手机号是否已验证
  email_verified boolean DEFAULT false NOT NULL, -- 邮箱是否已验证
  
  -- 状态字段
  status integer DEFAULT 1 NOT NULL, -- 状态：1-正常，2-禁用
  
  -- 扩展字段
  last_login_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL, -- 最后登录时间
  last_login_ip varchar(50) DEFAULT '' NOT NULL, -- 最后登录IP
  
  -- 时间戳字段
  created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL, -- 创建时间
  updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL, -- 更新时间
  deleted_at timestamp with time zone -- 删除时间（软删除）
);

-- 创建索引
CREATE UNIQUE INDEX IF NOT EXISTS idx_username ON users USING btree (username);
CREATE UNIQUE INDEX IF NOT EXISTS idx_phone ON users USING btree (phone);
CREATE UNIQUE INDEX IF NOT EXISTS idx_email ON users USING btree (email);
CREATE INDEX IF NOT EXISTS idx_status ON users USING btree (status);
CREATE INDEX IF NOT EXISTS idx_created_at ON users USING btree (created_at);
CREATE INDEX IF NOT EXISTS idx_last_login_at ON users USING btree (last_login_at);