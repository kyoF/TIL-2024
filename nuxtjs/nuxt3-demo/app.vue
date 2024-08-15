<template>
  <div>
    <NuxtLayout>
      <NuxtRouteAnnouncer />
      <NuxtPage />
      <img src="~/assets/cat.jpg" alt="cat">
    </NuxtLayout>
  </div>
</template>
<script setup>
useHead(
  {
    title: "Nuxt 3 Store",
    link: [
      { rel: "stylesheet", href: "https://fonts.googleapis.com/css2?family=Permanent+Marker&display=swap" }
    ]
  }
);
</script>
