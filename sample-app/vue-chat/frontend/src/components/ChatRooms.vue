<template>
  <div>
    <h1>VueChat - チャットルーム一覧</h1>
    <ul>
      <li v-for="room in chatRooms" :key="room.id">
        <router-link :to="`/rooms/${room.id}`">{{ room.name }}</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      chatRooms: [],
    };
  },
  mounted() {
    this.fetchChatRooms();
  },
  methods: {
    fetchChatRooms() {
      axios
        .get('http://localhost:5000/rooms')
        .then(response => {
          this.chatRooms = response.data;
        })
        .catch(error => {
          console.error(error);
        });
    },
  },
};
</script>
