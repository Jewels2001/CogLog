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
    font-family: Trebuchet, Verdana, sans-serif;
    font-weight: bold;
    background-color: #e3e3e3;
    border: none;
    color: #090909;
    padding: 0px 54px;
    text-align: center;
    display: inline-block;
    font-size: 64px;
    border-radius: 0.2em;
    transition: 0.3s ease-in-out;
    border-color: #e3e3e3;
}

button:hover{
    background-color: #090909;
    color: #e3e3e3;
    transition: 0.3s ease-in-out;
    border: 0.01em solid #e3e3e3;
    border-color: #e3e3e3;
    border-radius: 0.2em;
}

button:active{
    background-color: #e3e3e3;
    color: #090909;
    transition: 0.3s ease-in-out;
}

</style>

