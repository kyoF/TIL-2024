import { createRouter, createWebHistory } from 'vue-router';
import ChatRooms from '../components/ChatRooms.vue'
import ChatRoom from '../components/ChatRoom.vue'

const routes = [
  { path: '/', component: ChatRooms },
  { path: '/rooms/:id', component: ChatRoom, props: route => ({ roomId: route.params.id }) },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
