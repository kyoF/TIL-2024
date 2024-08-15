<template>
  <div>
    <h1>Users {{ $route.params.id }} Page</h1>
  </div>
</template>
<script setup>
  definePageMeta({
    validate: async(route) => {
      return /^\d+$/.test(route.params.id)
    }
  })
</script>
