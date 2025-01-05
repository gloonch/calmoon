CREATE TABLE signs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    element VARCHAR(20) NOT NULL,
    modality VARCHAR(20) NOT NULL,
    ruling_planet_id INT REFERENCES planets(id),
    unique_attributes JSONB
);

INSERT INTO signs (name, element, modality, ruling_planet_id, unique_attributes)
VALUES
-- Aries
('Aries', 'Fire', 'Cardinal', 1, '{"season": "Spring", "strength": "Courage", "keywords": ["Bold", "Leader", "Impulsive"]}'),
-- Taurus
('Taurus', 'Earth', 'Fixed', 2, '{"season": "Spring", "strength": "Stability", "keywords": ["Persistent", "Reliable", "Patient"]}'),
-- Gemini
('Gemini', 'Air', 'Mutable', 3, '{"season": "Spring", "strength": "Adaptability", "keywords": ["Curious", "Witty", "Communicative"]}'),
-- Cancer
('Cancer', 'Water', 'Cardinal', 4, '{"season": "Summer", "strength": "Empathy", "keywords": ["Nurturing", "Sensitive", "Protective"]}'),
-- Leo
('Leo', 'Fire', 'Fixed', 5, '{"season": "Summer", "strength": "Charisma", "keywords": ["Confident", "Creative", "Dramatic"]}'),
-- Virgo
('Virgo', 'Earth', 'Mutable', 6, '{"season": "Summer", "strength": "Analytical", "keywords": ["Practical", "Detail-oriented", "Efficient"]}'),
-- Libra
('Libra', 'Air', 'Cardinal', 7, '{"season": "Autumn", "strength": "Diplomacy", "keywords": ["Charming", "Fair", "Idealistic"]}'),
-- Scorpio
('Scorpio', 'Water', 'Fixed', 8, '{"season": "Autumn", "strength": "Passion", "keywords": ["Intense", "Determined", "Transformative"]}'),
-- Sagittarius
('Sagittarius', 'Fire', 'Mutable', 9, '{"season": "Autumn", "strength": "Optimism", "keywords": ["Adventurous", "Philosophical", "Independent"]}'),
-- Capricorn
('Capricorn', 'Earth', 'Cardinal', 10, '{"season": "Winter", "strength": "Discipline", "keywords": ["Ambitious", "Responsible", "Hardworking"]}'),
-- Aquarius
('Aquarius', 'Air', 'Fixed', 11, '{"season": "Winter", "strength": "Innovation", "keywords": ["Visionary", "Unconventional", "Humanitarian"]}'),
-- Pisces
('Pisces', 'Water', 'Mutable', 12, '{"season": "Winter", "strength": "Compassion", "keywords": ["Dreamy", "Intuitive", "Artistic"]}');
