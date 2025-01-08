export type Ingredient = {
  name: string;
  unit: string;
  kind: string;
  amount: number;
};

export type Meal = {
  id: number;
  name: string;
  ingredients: Array<Ingredient>;
  minutes: number;
  cost: number;
  image: string;
  recipe: string;
};
