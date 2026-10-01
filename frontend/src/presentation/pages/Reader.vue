<script setup lang="ts">
import Foot from '../components/Foot.vue'
import Mast from '../components/Mast.vue'
import { useReader } from '../viewmodels/useReader'

const { props, step, click, single, back, modeLabel, status } = useReader()
</script>

<template>
  <div class="flex h-full flex-col bg-[linear-gradient(90deg,rgba(194,59,34,0.05),transparent_12px),#f4efe4] text-ink">
    <Mast />
    <main class="shelf-scroll flex-1 overflow-auto p-[22px]">
      <p v-if="props.error" class="mb-3 text-vermilion">{{ props.error }}</p>
      <div
        class="flex min-h-full cursor-pointer flex-col items-stretch justify-center gap-2 sm:flex-row"
        @click="click"
      >
        <img
          v-if="props.frame.hasLeft"
          class="h-auto max-h-[calc(100vh-140px)] w-full bg-fold object-contain sm:w-[min(46vw,520px)]"
          :src="props.urls[props.frame.left]"
          alt=""
        />
        <img
          v-if="props.urls[props.frame.right]"
          class="h-auto max-h-[calc(100vh-140px)] w-full bg-fold object-contain sm:w-[min(46vw,520px)]"
          :src="props.urls[props.frame.right]"
          alt=""
        />
      </div>
    </main>
    <Foot>
      <button
        class="cursor-pointer border border-ink bg-transparent px-2.5 py-1.5 tracking-[0.12em] hover:bg-ink hover:text-washi"
        type="button"
        @click="back"
      >
        Back
      </button>
      <button
        class="cursor-pointer border border-ink bg-transparent px-2.5 py-1.5 tracking-[0.12em] hover:bg-ink hover:text-washi"
        type="button"
        @click="step(false)"
      >
        Prev
      </button>
      <button
        class="cursor-pointer border border-ink bg-transparent px-2.5 py-1.5 tracking-[0.12em] hover:bg-ink hover:text-washi"
        type="button"
        @click="step(true)"
      >
        Next
      </button>
      <button
        class="cursor-pointer border border-ink bg-transparent px-2.5 py-1.5 tracking-[0.12em] hover:bg-ink hover:text-washi"
        type="button"
        @click="single"
      >
        {{ modeLabel }}
      </button>
      <p class="ml-2 self-center text-sm text-ink/60">{{ status }}</p>
    </Foot>
  </div>
</template>
