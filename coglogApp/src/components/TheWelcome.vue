<script setup>
import WelcomeItem from './WelcomeItem.vue'
import Dropdown from './Dropdown.vue'
import Buttons from '../components/Buttons.vue'

import { ref } from 'vue'

const examplesDropdown = ref("")

let context;
let request;
let source;

// Change these to prover values
let URL = "192.168.2.14"
let PORT = "8080"

function TestMP3(abc) {
  try {
    // Request from server
    context = new AudioContext();
    request = new XMLHttpRequest();
    request.open(
      "POST",
      "http://" + URL + ":" + PORT + "/bci-decode/audio",
      true,
    );
    request.responseType = "arraybuffer";

    // Play Audio
    request.onload = () => {
      context.decodeAudioData(request.response, (buffer) => {
        source = context.createBufferSource();
        source.buffer = buffer;
        source.connect(context.destination);
        source.start(0);
      });
    };
    console.log(abc);
    request.send(abc);
  } catch (e) {
    console.error(e)
    alert("web audio api not supported");
  }
}

</script>

<template>
  <div class="container">
    <form @submit.prevent>
    <select v-model= "examplesDropdown">
    <label for="exampleSentences">Run an example:</label>
    <option disabled value="">Run one example</option>
            <option value="0">"theocracy reconsidered"</option>
            <option value="1">"rich purchased several signed lithographs"</option>
          <option value=2>"so rules we made in unabashed collusion"</option>
          <option value=3>"lori's costume needed black gloves to be completely elegant"</option>
          <option value=4>"the tooth fairy forgot to come when roger's tooth fell out"</option>
          <option value=5>"that stinging vapor was caused by chloride vaporization"</option>
          <option value=6>"before thursday's exam review every formula"</option>
          <option value=7>"wildfire near sunshine forces park closures"</option>
          <option value=8>"the word means it won't boil away easily nothing else"</option>
          <option value=9>"would a blue feather in a man's hat make him happy all day"</option>
    <!-- <input list="exampleSentences" id="exampleSentences"/> -->

    <!-- <datalist id="exampleSentences">
        <option value="theocracy reconsidered"></option>
        <option value="so rules we made in unabashed collusion"></option>
        <option value="lori's costume needed black gloves to be completely elegant"></option>
        <option value="the tooth fairy forgot to come when roger's tooth fell out"></option>
        <option value="that stinging vapor was caused by chloride vaporization"></option>
        <option value="before thursday's exam review every formula"></option>
        <option value="wildfire near sunshine forces park closures"></option>
        <option value="the word means it won't boil away easily nothing else"></option>
        <option value="would a blue feather in a man's hat make him happy all day"></option>
    </datalist> -->

  </select>
  <Buttons class="temp" @click="TestMP3(examplesDropdown)" title = "UPLOAD"/>
  <!-- <button @click="TestMP3(examplesDropdown)">Submit</button> -->
  </form>
    </div>

  <WelcomeItem>
    <template #heading>Brain to Speech Program</template>
    <a href="https://github.com/Jewels2001/CogLog" target="_blank" rel="noopener">Our GitHub</a>
    provides you with all information you need to get started.
    <br> Accepted files for upload: .............
  </WelcomeItem>
</template>

<style scoped>
.uploadBtn {
  position:absolute;
  margin:auto;
  top: 55%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 10px;
}
.container{
    position: inline-block;
    top: 45%;
    left: 50%;
    transform: translate(-50%, -50%);
    /* padding: 25px; */
}
.temp {
  text-align: center;
}

select {
  /* position: absolute; */
  font-family: 'Figtree', sans-serif;
    background-color: #202020;
    color: #b9d4b4;
    transition: 0.3s ease-in-out;
    border: 0.01em solid #b9d4b4;
    border-radius: 0.3em;
    padding: 1px 0px;
    margin: 5px;
    text-align: center;
}

/* form {
  padding: 2px 2px;
} */

</style>


<!--
<script setup>

import WelcomeItem from './WelcomeItem.vue'
import DocumentationIcon from './icons/IconDocumentation.vue'
import ToolingIcon from './icons/IconTooling.vue'
import EcosystemIcon from './icons/IconEcosystem.vue'
import CommunityIcon from './icons/IconCommunity.vue'
import SupportIcon from './icons/IconSupport.vue'
</script>

<template>
  <WelcomeItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Documentation</template>

    Vue’s
    <a href="https://vuejs.org/" target="_blank" rel="noopener">official documentation</a>
    provides you with all information you need to get started.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <ToolingIcon />
    </template>
    <template #heading>Tooling</template>

    This project is served and bundled with
    <a href="https://vitejs.dev/guide/features.html" target="_blank" rel="noopener">Vite</a>. The
    recommended IDE setup is
    <a href="https://code.visualstudio.com/" target="_blank" rel="noopener">VSCode</a> +
    <a href="https://github.com/johnsoncodehk/volar" target="_blank" rel="noopener">Volar</a>. If
    you need to test your components and web pages, check out
    <a href="https://www.cypress.io/" target="_blank" rel="noopener">Cypress</a> and
    <a href="https://on.cypress.io/component" target="_blank" rel="noopener"
      >Cypress Component Testing</a
    >.

    <br />

    More instructions are available in <code>README.md</code>.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <EcosystemIcon />
    </template>
    <template #heading>Ecosystem</template>

    Get official tools and libraries for your project:
    <a href="https://pinia.vuejs.org/" target="_blank" rel="noopener">Pinia</a>,
    <a href="https://router.vuejs.org/" target="_blank" rel="noopener">Vue Router</a>,
    <a href="https://test-utils.vuejs.org/" target="_blank" rel="noopener">Vue Test Utils</a>, and
    <a href="https://github.com/vuejs/devtools" target="_blank" rel="noopener">Vue Dev Tools</a>. If
    you need more resources, we suggest paying
    <a href="https://github.com/vuejs/awesome-vue" target="_blank" rel="noopener">Awesome Vue</a>
    a visit.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <CommunityIcon />
    </template>
    <template #heading>Community</template>

    Got stuck? Ask your question on
    <a href="https://chat.vuejs.org" target="_blank" rel="noopener">Vue Land</a>, our official
    Discord server, or
    <a href="https://stackoverflow.com/questions/tagged/vue.js" target="_blank" rel="noopener"
      >StackOverflow</a
    >. You should also subscribe to
    <a href="https://news.vuejs.org" target="_blank" rel="noopener">our mailing list</a> and follow
    the official
    <a href="https://twitter.com/vuejs" target="_blank" rel="noopener">@vuejs</a>
    twitter account for latest news in the Vue world.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <SupportIcon />
    </template>
    <template #heading>Support Vue</template>

    As an independent project, Vue relies on community backing for its sustainability. You can help
    us by
    <a href="https://vuejs.org/sponsor/" target="_blank" rel="noopener">becoming a sponsor</a>.
  </WelcomeItem>
</template>
-->