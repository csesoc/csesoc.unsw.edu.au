import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);
export default new Router({
  routes: [
    {
      path: '/',
      name: 'root',
      component: () => import('../views/Home/Home')
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/About')
    },
    {
      path: '/resources',
      name: 'resources',
      component: () => import('../views/Resources')
    },
    {
      path: '/sponsors',
      name: 'sponsors',
      component: () => import('../views/Sponsor')
    },
    {
      path: '/engage',
      name: 'engage',
      component: () => import('../views/Engage')
    }
  ],
  scrollBehavior() {
    // Scroll to top for all route navigations
    return { x: 0, y: 0 };
  }
});
