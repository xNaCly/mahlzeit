-- default meals:
INSERT INTO meals (id, name, minutes) VALUES 
(0, 'Spaghetti aglio olio', 20)
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount, kind) VALUES 
(0, 'Olive oil', 'ml', 100, "Oils & Vinegar"),
(0, 'Garlic', 'x', 4, "Spices"),
(0, 'Spaghetti', 'g', 500, "Carbs"),
(0, 'Parmesan', 'g', 50, "Diary")
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name, minutes) VALUES 
(1, 'Salad à la teo', 15)
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount, kind) VALUES 
(1, 'Salad', 'g', 250, "Vegetables"),
(1, 'Cucumber', 'x', 1, "Vegetables"),
(1, 'Mozarella', 'x', 2, "Diary"),
(1, 'Olive oil', 'ml', 100, "Oils & Vinegar"),
(1, 'Balsamico', 'ml', 50, "Oils & Vinegar"),
(1, 'Parmesan', 'g', 50, "Diary"),
(1, 'Red onion', 'x', 1, "Vegetables"),
(1, 'Spring onion', 'x', 2, "Vegetables")
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name, minutes) VALUES 
(2, 'Cheese spätzle', 10)
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount, kind) VALUES 
(2, 'Butter', 'g', 50, "Diary"),
(2, 'Eierspätzle', 'g', 500, "Carbs"),
(2, 'Cheese', 'g', 250, "Diary"),
(2, 'Roasted Onions', 'g', 50, "Vegetables")
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name, minutes) VALUES 
(3, 'Alsatian style tarte flambée', 20)
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount, kind) VALUES 
(3, 'Tarte Dough', 'x', 1, "Carbs"),
(3, 'Cream cheese', 'g', 250, "Diary"),
(3, 'Diced ham', 'g', 250, "Meat & Fish"),
(3, 'Red onions', 'x', 3, "Vegetables"),
(3, 'Spring onions', 'x', 6, "Vegetables")
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name, minutes) VALUES 
(4, 'Pizza Salami', 30)
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount, kind) VALUES
(4, 'Pizza dough', 'x', 1, "Carbs"),
(4, 'Tomato sauce', 'ml', 250, "Vegetables"),
(4, 'Mozarella', 'x', 2, "Diary"),
(4, 'Salami', 'g', 250, "Meat & Fish"),
(4, 'Pizza cheese', 'g', 250, "Diary")
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name, minutes) VALUES 
(5, 'Carbonara Adaptation', 30)
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount, kind) VALUES
(5, 'Penne', 'g', 250, "Carbs"),
(5, 'Diced Ham ', 'g', 250, "Meat & Fish"),
(5, 'Red onions', 'x', 3, "Vegetables"),
(5, 'Spring onions', 'x', 3, "Vegetables"),
(5, 'Cream', 'ml', 250, "Diary"),
(5, 'Parmesan', 'g', 50, "Diary")
ON CONFLICT DO NOTHING;

INSERT INTO meals (id, name, minutes) VALUES 
(6, 'Currywurst & Fries', 15)
ON CONFLICT DO NOTHING;
INSERT INTO ingredients (meal_id, name, unit, amount, kind) VALUES
(6, 'Currywurst', 'x', 2, "Meat & Fish"),
(6, 'Fries', 'g', 250, "Carbs"),
(6, 'Ketchup', 'ml', 50, "Sauces"),
(6, 'Bell pepper powder', 'ml', 50, "Spices")
ON CONFLICT DO NOTHING;
