<template>
  <div>
    <h1>Top Page</h1>
    <hr>
    <ul>
      <li><NuxtLink to="/price">price</NuxtLink></li>
      <li><NuxtLink to="/users">users</NuxtLink></li>
    </ul>
    <div>
      <ul>
        <li v-for="user in users" :key="user.id">{{ user.id }}: {{ user.name }}</li>
      </ul>
    </div>
    <div>
      <p>{{ title }}</p>
      <button @click="$event => titleState.changeTitle('Hello Nuxt3!')">
        changeTitle
      </button>
    </div>
  </div>
</template>
<script setup>
  const { data: users, error } = await useFetch('https://jsonplaceholder.typicode.com/users')
  if(error.value){
    throw createError({ statusCode: '404', statusMessage: 'Page Not Found!'})
  }

  const titleState = useTitle();
  const { title } = titleState
</script>
