-- default meals:
INSERT INTO meals (id, name) VALUES 
(0, 'Spaghetti aglio olio')
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount) VALUES 
(0, 'Olive oil', 'ml', 100),
(0, 'Garlic', 'x', 4),
(0, 'Spaghetti', 'g', 500),
(0, 'Parmesan', 'g', 50)
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name) VALUES 
(1, 'Salad à la teo')
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount) VALUES 
(1, 'Salad', 'g', 250),
(1, 'Cucumber', 'x', 1),
(1, 'Mozarella', 'x', 2),
(1, 'Olive oil', 'ml', 100),
(1, 'Balsamico', 'ml', 50),
(1, 'Parmesan', 'g', 50),
(1, 'Red onion', 'x', 1),
(1, 'Spring onion', 'x', 2)
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name) VALUES 
(2, 'Cheese spätzle')
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount) VALUES 
(2, 'Butter', 'g', 50),
(2, 'Eierspätzle', 'g', 500),
(2, 'Cheese', 'g', 250),
(2, 'Roasted Onions', 'g', 50)
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name) VALUES 
(3, 'Alsatian style tarte flambée')
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount) VALUES 
(3, 'Tarte Dough', 'x', 1),
(3, 'Cream cheese', 'g', 250),
(3, 'Diced ham', 'g', 250),
(3, 'Red onions', 'x', 3),
(3, 'Spring onions', 'x', 6)
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name) VALUES 
(4, 'Pizza Salami')
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount) VALUES
(4, 'Pizza dough', 'x', 1),
(4, 'Tomato sauce', 'ml', 250),
(4, 'Mozarella', 'x', 2),
(4, 'Salami', 'g', 250),
(4, 'Pizza cheese', 'g', 250)
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name) VALUES 
(5, 'Carbonara Adaptation')
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount) VALUES
(5, 'Penne', 'g', 250),
(5, 'Diced Ham ', 'g', 250),
(5, 'Red onions', 'x', 3),
(5, 'Spring onions', 'x', 3),
(5, 'Cream', 'ml', 250),
(5, 'Parmesan', 'g', 50)
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name) VALUES 
(6, 'Currywurst & Fries')
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount) VALUES
(6, 'Currywurst', 'x', 2),
(6, 'Fries', 'g', 250),
(6, 'Ketchup', 'ml', 50),
(6, 'Bell pepper powder', 'ml', 50)
ON CONFLICT DO NOTHING;
