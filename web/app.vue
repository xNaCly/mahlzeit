<script setup lang="ts">
useHead({ title: "Mahlzeit" });
import { ref } from "vue";
import { mahlzeitFetch } from "~/fetch";
import type { Ingredient, Meal } from "~/types";

const UCheckbox = resolveComponent("UCheckbox");
const UBadge = resolveComponent("UBadge");

const mealToDay = ["Mon", "Tue", "Wen", "Thu", "Fri", "Sat", "Sun"];

const meals = ref<Array<{ lock: boolean; meal: Meal }>>([]);
const summary = ref<{
  meal_names: Array<string>;
  ingredients_by_kind: Record<string, Array<Ingredient>>;
}>();

async function random(): Promise<Array<Meal>> {
  return await mahlzeitFetch<Meal[]>("meals/random");
}

/**
 *overcomplicated implementation to compute unique new meals, while respecting
 the users choice to keep certain meals
 */
async function regenerate() {
  let lockedMealIdToIndex: Record<number, number> = {};
  meals.value
    .map((m, i) => {
      return { id: m.meal?.id, index: i, lock: m.lock };
    })
    .filter((e) => e.lock)
    .forEach((e) => (lockedMealIdToIndex[e.id] = e.index));
  let randomMeals = await random();
  let filteredMeals = randomMeals.filter((e) => !(e.id in lockedMealIdToIndex));
  for (let i = 0; i < 7; i++) {
    {
      if (Object.values(lockedMealIdToIndex).includes(i)) {
        continue;
      }
      meals.value[i] = { lock: false, meal: filteredMeals.pop()! };
    }
  }
  summary.value = await mahlzeitFetch(
    `meals/summary?meal_ids=${meals.value.map((e) => e.meal.id).join(",")}`
  );
}

onMounted(async () => {
  meals.value = (await random()).map((meal) => {
    return { lock: false, meal: meal };
  });
  summary.value = await mahlzeitFetch(
    `meals/summary?meal_ids=${meals.value.map((e) => e.meal.id).join(",")}`
  );
});

const columns = [
  {
    id: "select",
    header: "Bought",
    cell: ({ row }: any) =>
      h(UCheckbox, {
        modelValue: false,
      }),
  },
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "kind",
    header: "Kind",
    cell: ({ row }: any) => {
      let color = (
        {
          "Oils & Vinegar": "primary",
          Spices: "secondary",
          Vegetables: "success",
          Diary: "info",
          Carbs: "warning",
          "Meat & Fish": "error",
          Sauces: "neutral",
        } as Record<string, string>
      )[row.getValue("kind")];

      return h(
        UBadge,
        { class: "capitalise", variant: "subtle", color },
        row.getValue("kind")
      );
    },
  },
  {
    accessorKey: "amount",
    header: "Amount",
    cell: ({ row }: any) => {
      return `${row.getValue("amount")}${row.getValue("unit")}`;
    },
  },
  {
    accessorKey: "unit",
    header: "",
    cell: ({ row }: any) => {},
  },
];
</script>
<template>
  <UApp>
    <UContainer>
      <UCard class="mt-10">
        <template #header>
          <div class="flex justify-between">
            <h1 class="text-lg">Mahlzeit</h1>
            <div>
              <UButton
                icon="i-lucide-rotate-cw"
                size="md"
                color="primary"
                variant="subtle"
                class="cursor-pointer"
                @click="regenerate()"
                >Regenerate</UButton
              >
            </div>
          </div>
        </template>
      </UCard>

      <div class="flex flex-col m-2" style="flex: 1 1 0px">
        <MealPreview
          :index
          v-for="(meal, index) in meals"
          :mealAndLock="meal"
          @change="(e: boolean) => (meals[index].lock = e)"
        />
      </div>

      <div class="flex flex-col m-3">
        <UTable
          class="border border-gray-300 rounded mt-2"
          :columns="columns"
          :data="Object.values(summary?.ingredients_by_kind || {}).flat()"
        />
      </div>
    </UContainer>
  </UApp>
</template>
