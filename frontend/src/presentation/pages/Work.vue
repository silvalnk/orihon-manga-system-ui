<script setup lang="ts">
import Foot from '../components/Foot.vue'
import Mast from '../components/Mast.vue'
import Rail from '../components/Rail.vue'
import { useWork } from '../viewmodels/useWork'

const { props, back, open, toggle, openFavorite } = useWork()
</script>

<template>
  <div class="flex h-full flex-col bg-washi text-ink">
    <Mast />
    <div class="flex min-h-0 flex-1 gap-4 px-6 pt-4">
      <Rail :favorites="props.favorites" @open="openFavorite" />
      <main class="shelf-scroll min-w-0 flex-1 overflow-auto pr-2">
        <p v-if="props.error" class="mb-3 text-vermilion">{{ props.error }}</p>
        <article v-if="props.manga.id">
          <div class="flex items-start gap-5">
            <img
              class="h-40 w-28 shrink-0 bg-fold object-cover shadow-[inset_-10px_0_0_rgba(28,25,21,0.06)]"
              :src="props.manga.cover"
              :alt="props.manga.title"
            />
            <div class="min-w-0 flex-1">
              <div class="flex items-start justify-between gap-4">
                <h1 class="text-2xl">{{ props.manga.title }}</h1>
                <button
                  type="button"
                  class="grid size-8 shrink-0 cursor-pointer place-items-center rounded-full border-2 border-vermilion"
                  :class="props.favorited ? 'bg-vermilion text-washi' : 'bg-transparent text-vermilion'"
                  :aria-pressed="props.favorited"
                  :aria-label="props.favorited ? 'Remove from favorites' : 'Add to favorites'"
                  @click="toggle"
                >
                  <svg class="size-[18px]" viewBox="0 0 24 24" aria-hidden="true">
                    <path
                      d="M12 2.4 14.9 8.3 21.2 9.2 16.6 13.7 17.7 20 12 17 6.3 20 7.4 13.7 2.8 9.2 9.1 8.3Z"
                      :fill="props.favorited ? 'currentColor' : 'none'"
                      stroke="currentColor"
                      stroke-width="1.6"
                      stroke-linejoin="round"
                    />
                  </svg>
                </button>
              </div>
              <p class="mt-3 whitespace-pre-wrap leading-relaxed">{{ props.manga.description }}</p>
            </div>
          </div>
          <ol class="mt-6 list-none p-0">
            <li v-for="chapter in props.chapters" :key="chapter.id" class="mb-2">
              <button
                class="w-full cursor-pointer border-0 border-t-[3px] border-vermilion bg-washi px-4 py-3 text-left"
                type="button"
                @click="open(chapter)"
              >
                fold {{ chapter.chapter }} · {{ chapter.lang }} · {{ chapter.title }}
              </button>
            </li>
          </ol>
        </article>
      </main>
    </div>
    <Foot>
      <button
        class="cursor-pointer border border-ink bg-transparent px-2.5 py-1.5 tracking-[0.12em] hover:bg-ink hover:text-washi"
        type="button"
        @click="back"
      >
        Back
      </button>
    </Foot>
  </div>
</template>
