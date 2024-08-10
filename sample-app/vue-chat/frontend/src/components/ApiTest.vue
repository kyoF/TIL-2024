<template>
  <div>
    <h1>ApiTest</h1>
    <button @click='test'>ApiTest</button>
    <p>{{ message }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: ''
    }
  },
  methods: {
    test() {
      fetch('http://localhost:5000/api/test')
        .then((response) => response.json())
        .then((data) => {
          this.message = data.message
        })
    }
  }
}
</script>
