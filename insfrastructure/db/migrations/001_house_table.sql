CREATE TABLE houses (
    id SERIAL PRIMARY KEY,
    number INT NOT NULL CHECK (number >= 1 AND number <= 12),
    associated_sign_id INT REFERENCES signs(id),
    ruling_planet_id INT REFERENCES planets(id),
    unique_attributes JSONB,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE house_planet_occupancy (
    house_id INT REFERENCES houses(id) ON DELETE CASCADE,
    planet_id INT REFERENCES planets(id) ON DELETE CASCADE,
    PRIMARY KEY (house_id, planet_id)
);


INSERT INTO houses (number, associated_sign_id, ruling_planet_id, unique_attributes)
VALUES
    (1, 1, 1, '{"focus": "personality", "area_of_life": ["self", "health"], "primary_element": "Fire"}'),
    (2, 2, 2, '{"focus": "wealth", "area_of_life": ["resources", "family"], "primary_element": "Earth"}');


SELECT h.*, p.name AS ruling_planet_name
FROM houses h
         LEFT JOIN planets p ON h.ruling_planet_id = p.id
WHERE h.number = 1;


SELECT p.name AS planet_name
FROM house_planet_occupancy hpo
         JOIN planets p ON hpo.planet_id = p.id
WHERE hpo.house_id = 1;


CREATE INDEX idx_houses_number ON houses (number);
CREATE INDEX idx_house_planet_occupancy_house_id ON house_planet_occupancy (house_id);
