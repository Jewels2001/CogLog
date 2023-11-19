<script setup>
// Import stuff you need here
import {ref} from 'vue'
import { reactive } from 'vue';
import WelcomeItem from './WelcomeItem.vue'


const result_dat = reactive({
    txt: 'This is your result in text form, This message would be replaced with text data decoded by the model.',
    aup: 'tts-service/test.mp3'
})

//const result_txt = ref('This is your result in text form, This message would be replaced with text data decoded by the model.')
//const result_audio = reactive('tts-service\test.mp3') // Path to audio file should be here...?
</script>

<template>
    <h1>Results</h1>
    <div class="left_item">
        {{result_dat.txt}}
    </div>
    <div class="left_item">
        <button>+ New</button>
    </div>
    <WelcomeItem>
        <audio controls>
            <source src={{result_dat.aup}} type="audio/mpeg">
            Your browser does not support the audio element.
        </audio>
    </WelcomeItem>
    
</template>

<style scoped>
h1{
  font-size: 3rem;
  font-weight: 500;
  color: var(--color-heading);
  position: relative;
  text-align: left;
  transform: translate(-50%, -50%);
}

.item {
  text-align: left;
  display: flex;
  position: absolute;
  top: 90%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.right_item {
  text-align: right;
  display: flex;
  position: relative;
  transform: translate(-50%, -50%);
}

.left_item {
  text-align: left;
  display: flex;
  position: relative;
  transform: translate(-50%, -50%);
}
</style>