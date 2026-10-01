<script setup lang="ts">
import type { Manga } from '../../domain/manga'

defineProps<{ favorites: Manga[] }>()
const emit = defineEmits<{ open: [manga: Manga] }>()
</script>

<template>
  <aside class="shelf-scroll w-[260px] shrink-0 overflow-auto rounded-lg bg-fold px-3 py-3.5">
    <h2 class="px-1 text-base">Favorites</h2>
    <p v-if="favorites.length === 0" class="mt-4 px-1 text-sm leading-relaxed text-ink/60">
      Stamp a title with the vermilion mark to keep it on the shelf.
    </p>
    <ul v-else class="mt-4 flex list-none flex-col gap-2 p-0">
      <li v-for="manga in favorites" :key="manga.id">
        <button
          class="w-full cursor-pointer rounded bg-washi px-2.5 py-4 text-left text-sm"
          type="button"
          @click="emit('open', manga)"
        >
          {{ manga.title }}
        </button>
      </li>
    </ul>
  </aside>
</template>
