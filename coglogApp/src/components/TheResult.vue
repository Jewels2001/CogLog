<script setup>
// Import stuff you need here
import {ref} from 'vue';
import { reactive } from 'vue';
import WelcomeItem from './WelcomeItem.vue';


const result_dat = reactive({
    txt: 'This is your result in text form, This message would be replaced with text data decoded by the model.',
    aup: '@/components/audio/test.mp3'
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
        <button class="restart_btn">+ New</button>
    </div>
    <WelcomeItem>
        <audio class="audio_s" controls>
            <source src=result_dat.aup type="audio/mpeg">
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

.audio_s{
    width: 610px;
    color: #b9d4b4;
    border-radius: 0.2em;
}
.restart_btn{
    font-family: 'Figtree', sans-serif;
    background-color: #202020;
    color: #b9d4b4;
    transition: 0.3s ease-in-out;
    border: 0.01em solid #b9d4b4;
    border-radius: 0.2em;
    padding: 0px 14px;
    text-align: left;
    display: inline-block;
    font-size: 12px;
}

.restart_btn:hover{
    background-color: #6aae9d;
    color: #202020;
    transition: 0.3s ease-in-out;
    border-radius: 0.2em;
}

.restart_btn:active{
    background-color: #b9d4b4;
    color: #202020;
    transition: 0.3s ease-in-out;
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