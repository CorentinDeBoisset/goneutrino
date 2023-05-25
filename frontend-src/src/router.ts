import { createRouter, createWebHistory, RouteLocationRaw, RouteRecordRaw  } from 'vue-router';
import { useNeutrinoStore } from './appStore';
import HomePage from '@/components/HomePage.vue'
import NewChatPage from '@/components/NewChatPage.vue'
import SecretExchange from '@/components/SecretExchange.vue'
import ErrorPage from '@/components/ErrorPage.vue'

const routes: Array<RouteRecordRaw> = [
  {
    name: 'home',
    path: '/',
    component: HomePage,
    meta: { initState: "unregistered" }
  },
  {
    name: 'new-chat',
    path: '/new-chat',
    component: NewChatPage,
    meta: { initState: "registered" }
  },
  {
    name: 'chat',
    path: '/chat/:uuid',
    component: SecretExchange,
    meta: { initState: "registered" }
  },
  {
    name: 'error',
    path: '/error',
    component: ErrorPage,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeResolve(async (to): Promise<boolean|RouteLocationRaw> => {
  const store = useNeutrinoStore();

  if (to.meta.initState) {
    // If not inited yet, we start the initialization
    if (!store.initHasRun) {
      await store.initializeData();
      if (store.initError) {
        return { name: 'error', query: { errorType: "init" } };
      }
    }

    // Then we redirect if the user is not registered yet
    if (to.meta.initState == "registered" && store.nickname === null) {
      console.warn("Access to a page requiring a registered session was cancelled.")
      return { name: 'home' };
    } else if (to.meta.initState == "unregistered" && store.nickname !== null) {
      return { name: 'new-chat' };
    }
  }
  return true
})

export default router;
