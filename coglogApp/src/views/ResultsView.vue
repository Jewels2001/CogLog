<script setup>
import TheResult from '../components/TheResult.vue';
</script>

<template>
  <main>
    <TheResult />
  </main>
</template>
  
<style scoped>
  .results {
  text-align: center;
  display: flex;
  }
  h1 {
    font-size: 2rem;
    font-weight: 500;
    color: var(--color-heading);
    position: absolute;
    top: 20%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
  </style>
  