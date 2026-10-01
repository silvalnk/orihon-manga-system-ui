<script setup lang="ts">
import { useChrome } from '../viewmodels/useChrome'

const { q, session, home, submit } = useChrome()
</script>

<template>
  <header
    class="flex h-[88px] items-center gap-4 border-b border-line bg-fold px-6"
    :class="session.busy ? 'shadow-[inset_0_-2px_0_#c23b22]' : ''"
  >
    <button
      class="flex w-[186px] shrink-0 cursor-pointer items-center gap-3 border-0 bg-transparent text-[22px]"
      type="button"
      @click="home"
    >
      <svg class="size-9 shrink-0" viewBox="0 0 36 36" aria-hidden="true">
        <circle cx="18" cy="18" r="18" class="fill-vermilion" />
        <polyline
          points="10.8,24.3 10.8,11.7 18,17.1 25.2,11.7 25.2,24.3"
          fill="none"
          class="stroke-washi"
          stroke-width="2"
          stroke-linejoin="miter"
          stroke-linecap="butt"
        />
      </svg>
      Orihon
    </button>
    <form class="flex flex-1" @submit.prevent="submit">
      <input
        id="search"
        v-model="q"
        class="h-10 w-full rounded border border-line bg-washi px-2.5 outline-none focus:outline focus:outline-2 focus:outline-vermilion"
        type="search"
        placeholder="search a title…"
        autocomplete="off"
      />
    </form>
    <span
      class="ml-auto grid size-8 place-items-center rounded-full border-2 border-vermilion bg-vermilion text-xs tracking-[0.08em] text-washi"
      aria-label="English"
    >
      EN
    </span>
  </header>
</template>
