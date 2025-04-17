
CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL,
  password text NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  CONSTRAINT unique_username UNIQUE (username),
  CONSTRAINT unique_email UNIQUE (email)
);

CREATE TABLE blogs (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  author UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE follow (
  follower UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  followee UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT unique_follower_followee UNIQUE (follower, followee)
);

-- indexes
CREATE INDEX idx_username ON users(username);
CREATE INDEX idx_email ON users(email);
CREATE INDEX idx_follower ON follow(follower);
CREATE INDEX idx_followee ON follow(followee);
CREATE INDEX idx_blogs_author_created_at ON blogs(author, created_at DESC);
