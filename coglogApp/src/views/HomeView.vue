<script setup>
import TheWelcome from '../components/TheWelcome.vue'
import Buttons from '../components/Buttons.vue'
</script>

<template>
  <main>
    <TheWelcome />
    <form class = "chooseFile">
      <input type = "file" id = "fileUpload" />
    </form>
    <div class = "uploadBtn">
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

.chooseFile{
  position: absolute;
  top: 42%;
  left: 50%;
  transform: translate(-50%, -50%);
}

input[type="file"]::file-selector-button {
  font-family: 'Figtree', sans-serif;
  background-color: #b9d4b4;
  color: #202020;
  transition: 0.3s ease-in-out;
  border: 0.01em solid #b9d4b4;
  border-radius: 0.2em;
  padding: 0px 25px;
  text-align: center;
  display: inline-block;
  font-size: 20px;
}

input[type="file"]::file-selector-button:hover {
  background-color: #6aae9d;
  color: #202020;
  transition: 0.3s ease-in-out;
  border: none;
  border-radius: 0.2em;
}

input[type="file"]::file-selector-button:active {
  background-color: #b9d4b4;
  color: #202020;
  transition: 0.3s ease-in-out;
}

</style>