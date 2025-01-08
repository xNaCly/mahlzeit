<template>
  <UApp>
    <UContainer>
      <UCard class="mt-10">
        <template #header>
          <div class="flex justify-between">
            <h1 class="text-lg">Mahlzeit</h1>
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
        </template>
      </UCard>

      <div class="flex flex-col m-2" style="flex: 1 1 0px">
        <MealPreview
          v-for="(meal, index) in meals"
          :mealAndLock="meal"
          @change="(e: boolean) => (meals[index].lock = e)"
        />
      </div>
    </UContainer>
  </UApp>
</template>

<script setup lang="ts">
useHead({ title: "Mahlzeit" });
import { ref } from "vue";
import { mahlzeitFetch } from "~/fetch";
import type { Meal } from "~/types";

const meals = ref<Array<{ lock: boolean; meal: Meal }>>([]);

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
}

onMounted(async () => {
  meals.value = (await random()).map((meal) => {
    return { lock: false, meal: meal };
  });
});
</script>
