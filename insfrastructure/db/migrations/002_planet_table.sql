CREATE TABLE planets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    is_retrograde BOOLEAN DEFAULT FALSE,
    element VARCHAR(20) NOT NULL,
    modality VARCHAR(20) NOT NULL,
    degree NOT NULL CHECK (degree >= 0 AND degree <= 30),
    sign VARCHAR(20) NOT NULL,
    house_num INT NOT NULL CHECK (house_num >= 1 AND house_num <= 12),
    timestamp TIMESTAMP DEFAULT NOW(),
    exaltation_sign_id INT,
    debilitation_sign_id INT,
    unique_attributes JSONB
);

CREATE TABLE ruling_signs (
    planet_id INT REFERENCES planets(id) ON DELETE CASCADE,
    sign_id INT NOT NULL,
    PRIMARY KEY (planet_id, sign_id)
);

CREATE TABLE planet_relations (
    planet_id INT REFERENCES planets(id) ON DELETE CASCADE,
    related_planet_id INT REFERENCES planets(id) ON DELETE CASCADE,
    relation_type INT NOT NULL CHECK (relation_type IN (0, 1, 2)),
    PRIMARY KEY (planet_id, related_planet_id)
);


INSERT INTO planets (name, is_retrograde, element, modality, degree, sign, house_num, exaltation_sign_id, debilitation_sign_id)
VALUES ('Mars', FALSE, 'Fire', 'Cardinal', 23.5, 'Aries', 1, 5, 11);


CREATE INDEX idx_planets_name ON planets (name);
CREATE INDEX idx_ruling_signs_planet_id ON ruling_signs (planet_id);
CREATE INDEX idx_planet_relations_planet_id ON planet_relations (planet_id);
