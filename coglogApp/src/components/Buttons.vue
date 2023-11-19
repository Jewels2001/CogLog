<script setup> 
defineProps(['title'])
</script>

<template>
    <button class = "uploadButton">
        {{ title }}
    </button>
</template>

<style scoped>
button {
    font-family: 'Figtree', sans-serif;
    background-color: #202020;
    color: #b9d4b4;
    transition: 0.3s ease-in-out;
    border: 0.01em solid #b9d4b4;
    border-radius: 0.2em;
    padding: 0px 54px;
    text-align: center;
    display: inline-block;
    font-size: 64px;
}

button:hover{
    background-color: #6aae9d;
    color: #202020;
    transition: 0.3s ease-in-out;
    border-radius: 0.2em;
}

button:active{
    background-color: #b9d4b4;
    color: #202020;
    transition: 0.3s ease-in-out;
}

</style>

