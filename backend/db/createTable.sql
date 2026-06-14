-- Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";



-- =========================
-- USERS
-- =========================

CREATE TABLE users (

    id UUID PRIMARY KEY
    DEFAULT gen_random_uuid(),

    email VARCHAR(255)
    UNIQUE NOT NULL,

    username VARCHAR(50)
    UNIQUE NOT NULL,

    password_hash VARCHAR(255),

    google_id VARCHAR(255)
    UNIQUE,

    role VARCHAR(20)
    NOT NULL DEFAULT 'user',

    created_at TIMESTAMP WITH TIME ZONE
    NOT NULL DEFAULT CURRENT_TIMESTAMP
);



-- =========================
-- BOOKS
-- =========================

CREATE TABLE books (

    id UUID PRIMARY KEY
    DEFAULT gen_random_uuid(),

    title VARCHAR(255)
    NOT NULL,

    author VARCHAR(255)
    NOT NULL,

    cover_image_url VARCHAR(2048),

    description TEXT,

    created_at TIMESTAMP WITH TIME ZONE
    NOT NULL DEFAULT CURRENT_TIMESTAMP
);



-- =========================
-- REVIEWS
-- =========================

CREATE TABLE reviews (

    id UUID PRIMARY KEY
    DEFAULT gen_random_uuid(),

    user_id UUID
    NOT NULL,

    book_id UUID
    NOT NULL,

    review_text TEXT
    NOT NULL,

    rating INTEGER
    NOT NULL
    CHECK (rating >= 1 AND rating <= 5),

    updated_at TIMESTAMP WITH TIME ZONE
    NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    FOREIGN KEY (book_id)
    REFERENCES books(id)
    ON DELETE CASCADE
);



-- =========================
-- COMMENTS
-- =========================

CREATE TABLE comments (

    id UUID PRIMARY KEY
    DEFAULT gen_random_uuid(),

    user_id UUID
    NOT NULL,

    review_id UUID
    NOT NULL,

    comment_text TEXT
    NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE
    NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    FOREIGN KEY (review_id)
    REFERENCES reviews(id)
    ON DELETE CASCADE
);



-- =========================
-- QUOTES
-- =========================

CREATE TABLE quotes (

    id UUID PRIMARY KEY
    DEFAULT gen_random_uuid(),

    user_id UUID
    NOT NULL,

    book_id UUID,

    quote_text TEXT
    NOT NULL,

    raw_image_url VARCHAR(2048),

    created_at TIMESTAMP WITH TIME ZONE
    NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    FOREIGN KEY (book_id)
    REFERENCES books(id)
    ON DELETE SET NULL
);



-- =========================
-- FOLLOWS
-- =========================

CREATE TABLE follows (

    follower_id UUID
    NOT NULL,

    following_id UUID
    NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE
    NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (follower_id, following_id),

    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    FOREIGN KEY (following_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    CHECK (follower_id != following_id)
);



-- =========================
-- REVIEW LIKES
-- =========================

CREATE TABLE review_likes (

    user_id UUID
    NOT NULL,

    review_id UUID
    NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE
    NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id, review_id),

    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    FOREIGN KEY (review_id)
    REFERENCES reviews(id)
    ON DELETE CASCADE
);



-- =========================
-- QUOTE LIKES
-- =========================

CREATE TABLE quote_likes (

    user_id UUID
    NOT NULL,

    quote_id UUID
    NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE
    NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id, quote_id),

    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    FOREIGN KEY (quote_id)
    REFERENCES quotes(id)
    ON DELETE CASCADE
);



-- =========================
-- INDEXES
-- =========================

CREATE INDEX idx_reviews_book_id
ON reviews(book_id);

CREATE INDEX idx_reviews_user_id
ON reviews(user_id);

CREATE INDEX idx_comments_review_id
ON comments(review_id);

CREATE INDEX idx_follows_following_id
ON follows(following_id);