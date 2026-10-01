<script setup lang="ts">
import Mast from '../components/Mast.vue'
import Rail from '../components/Rail.vue'
import { useShelf } from '../viewmodels/useShelf'

const { props, folds, status, tail, onScroll, open } = useShelf()
</script>

<template>
  <div class="flex h-full flex-col bg-washi text-ink">
    <Mast />
    <div class="flex min-h-0 flex-1 gap-4 px-6 pt-4 pb-3">
      <Rail :favorites="props.favorites" @open="open" />
      <section class="flex min-w-0 flex-1 flex-col">
        <p class="mb-2 h-9 text-sm text-ink/60">{{ status }}</p>
        <main class="shelf-scroll min-h-0 flex-1 overflow-auto pr-2" @scroll="onScroll">
          <div class="grid grid-cols-3 gap-3">
            <button
              v-for="manga in folds"
              :key="manga.id"
              class="flex h-32 w-full cursor-pointer gap-2.5 border-0 bg-paper p-2.5 text-left shadow-[inset_-10px_0_0_rgba(28,25,21,0.06)]"
              type="button"
              @mousedown.prevent="open(manga)"
            >
              <img class="h-[108px] w-[72px] shrink-0 bg-fold object-cover" :src="manga.cover" :alt="manga.title" />
              <span class="min-w-0">
                <span class="line-clamp-2 text-base">{{ manga.title }}</span>
                <span class="mt-1.5 line-clamp-3 text-sm leading-snug text-ink/60">{{ manga.description }}</span>
              </span>
            </button>
          </div>
          <div ref="tail" class="h-px"></div>
        </main>
      </section>
    </div>
  </div>
</template>
