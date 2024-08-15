<template>
  <div>
    <p>{{ num }}</p>
    <button @click="$event => numState.countUp()">+</button>
    <button @click="$event => numState.countDown()">-</button>
    <button @click="$event => numState.countReset()">Reset</button>
  </div>
</template>
<script setup>
  const numState = useCounter();
  const { num } = numState;
</script>
