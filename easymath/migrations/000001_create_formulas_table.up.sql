CREATE TABLE IF NOT EXISTS formulas (
                                        id SERIAL PRIMARY KEY NOT NULL,
                                        chapter text NOT NULL,
                                        level text NOT NULL,
                                        withvar INT NOT NULL
);