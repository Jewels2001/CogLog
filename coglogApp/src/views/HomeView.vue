<script setup>
import TheWelcome from '../components/TheWelcome.vue'
import Buttons from '../components/Buttons.vue'
</script>

<template>
  <main>
    <TheWelcome />
    <div class="uploadBtn">
      <Buttons title = "UPLOAD"/>
    </div>
  </main>
</template>

<style scoped>
.uploadBtn {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 10px;
}
</style>