<template>
  <UApp>
    <UContainer>
      <UCard class="mt-10">
        <template #header>
          <div class="flex justify-between">
            <h1 class="text-lg">Mahlzeit</h1>
          </div>
        </template>
      </UCard>
    </UContainer>
  </UApp>
</template>

<script setup lang="ts">
import { mahlzeitFetch } from "./fetch";
const toast = useToast();

onMounted(async () => {
  try {
    await mahlzeitFetch("ping");
  } catch (e: unknown) {
    toast.add({
      title: "Connection failure",
      description: "Failed to connect to server",
      color: "error",
    });
  }
});
</script>
