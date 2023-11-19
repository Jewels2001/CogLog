<script setup>

</script>

<template>
    <div class="container">
    <form action = "examplesDropdown">
    <label for="exampleSentences">Run an example:</label>
    
    <input list="exampleSentences" id="exampleSentences"/>

    <datalist id="sentences">
        <option value="theocracy reconsidered"></option>
        <option value="so rules we made in unabashed collusion"></option>
        <option value="lori's costume needed black gloves to be completely elegant"></option>
        <option value="the tooth fairy forgot to come when roger's tooth fell out"></option>
        <option value="that stinging vapor was caused by chloride vaporization"></option>
        <option value="before thursday's exam review every formula"></option>
        <option value="wildfire near sunshine forces park closures"></option>
        <option value="the word means it won't boil away easily nothing else"></option>
        <option value="would a blue feather in a man's hat make him happy all day"></option>
    </datalist>
    </form>
</div>
</template>

<style scoped>
.container{
    position: absolute;
    top: 20%;
    left: 50%;
    transform: translate(-50%, -50%);
    padding: 10px;
}
</style>
