import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);
export default new Router({
  routes: [
    {
      path: '/',
      name: 'root',
      component: () => import('../views/Home/Home.vue'),
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/About.vue'),
    },
    {
      path: '/resources',
      name: 'resources',
      component: () => import('../views/Resources.vue'),
      redirect: (to) => {
        console.log(to);
        window.location.href = 'https://media.csesoc.org.au/tag/csesoc/';
      },
    },
    {
      path: '/sponsors',
      name: 'sponsors',
      component: () => import('../views/Sponsor.vue'),
    },
    {
      path: '/engage',
      name: 'engage',
      component: () => import('../views/Engage.vue'),
    }
  ],
  scrollBehavior() {
    // Scroll to top for all route navigations
    return { x: 0, y: 0 };
  }
});
