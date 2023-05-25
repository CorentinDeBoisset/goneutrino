import { createRouter, createWebHistory, RouteRecordRaw  } from 'vue-router';
import NameSelector from '@/components/NameSelector.vue'
import PeerSelector from '@/components/PeerSelector.vue'
import SecretExchange from '@/components/SecretExchange.vue'

const routes: Array<RouteRecordRaw> = [
  { path: '/', component: NameSelector },
  { path: '/new-chat', component: PeerSelector },
  { path: '/chat/:uuid', component: SecretExchange },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
