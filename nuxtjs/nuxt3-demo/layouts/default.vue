<template>
  <div class="container">
    <slot />
    <footer>&copy; proglus.jp</footer>
  </div>
</template>
<style>
.container {
  color: red,
}

footer {
  background-color: orange;
  color: white;
  text-align: center;
}

html {
  font-family: 'Permanent Marker', cursive;
}
</style>
