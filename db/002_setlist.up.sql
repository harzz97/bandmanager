-- Setlists Table
CREATE TABLE setlists (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    event_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Setlist Songs Table
CREATE TABLE setlist_songs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    setlist_id INT NOT NULL,
    song_id INT NOT NULL,
    musician VARCHAR(255),
    notes TEXT,
    position INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (setlist_id) REFERENCES setlists (id) ON DELETE CASCADE,
    FOREIGN KEY (song_id) REFERENCES songs (id) ON DELETE CASCADE,
    CONSTRAINT unique_setlist_song UNIQUE (setlist_id, song_id)
);
