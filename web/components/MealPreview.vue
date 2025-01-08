<script setup lang="ts">
import { mahlzeitFetch } from "~/fetch";
import { h, resolveComponent } from "vue";
import type { Ingredient, Meal } from "~/types";
const UBadge = resolveComponent("UBadge");

const slideoverOpen = ref(false);

const { mealAndLock } = defineProps<{
  mealAndLock: { lock: boolean; meal: Meal };
}>();

const ingredients = ref<Array<Ingredient>>();

const emit = defineEmits(["change"]);

onMounted(async () => {
  mealAndLock.meal = await mahlzeitFetch<Meal>("meals/" + mealAndLock.meal.id);
});

const columns = [
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
        :data="[
          { meta: 'Time to cook', value: mealAndLock.meal.minutes + 'min' },
          { meta: 'Cost & Complexity', value: `${mealAndLock.meal.cost}/100` },
        ]"
      />
      <UTable
        class="border border-gray-300 rounded mt-2"
        :columns="columns"
        :data="mealAndLock.meal.ingredients"
      />
    </template>
  </USlideover>
</template>
