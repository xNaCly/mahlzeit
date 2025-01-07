<script setup lang="ts">
import { mahlzeitFetch } from "~/fetch";
import type { Ingredient, Meal } from "~/types";

const slideoverOpen = ref(false);

const { mealAndLock } = defineProps<{
  mealAndLock: { lock: boolean; meal: Meal };
}>();

const ingredients = ref<Array<Ingredient>>();

const emit = defineEmits(["change"]);

onMounted(async () => {
  ingredients.value = (
    await mahlzeitFetch<Meal>("meals/" + mealAndLock.meal.id)
  ).ingredients;
});
</script>
<template>
  <UCard class="m-1">
    <div class="flex flex-row justify-between">
      <UButton
        @click="slideoverOpen = true"
        class="text-sm cursor-pointer"
        variant="ghost"
        color="neutral"
        :label="mealAndLock.meal.name"
      />
      <div class="flex items-center justify-center" @click.prevent="">
        <USwitch
          size="lg"
          color="primary"
          @update:modelValue="(v: boolean) => emit('change', v)"
        />
      </div>
    </div>
  </UCard>
  <USlideover
    v-model:open="slideoverOpen"
    :title="mealAndLock.meal.name"
    sticky
  >
    <template #body>
      <UTable
        class="border border-gray-300 rounded"
        :data="
          ingredients?.map((i) => {
            return { name: i.name, amount: `${i.amount}${i.unit}` };
          })
        "
      />
    </template>
  </USlideover>
</template>
