export type Ingredient = {
  name: string;
  unit: string;
  amount: number;
};

export type Meal = {
  id: number;
  name: string;
  ingredients: Array<Ingredient>;
  image: string;
  recipe: string;
};
