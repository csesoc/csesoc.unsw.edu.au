import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);
export default new Router({
  mode: 'history',
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
    // {
    //   path: '/resources',
    //   name: 'resources',
    //   component: () => import('../views/Resources')
    // },
    {
      path: '/sponsors',
      name: 'sponsors',
      component: () => import('../views/Sponsor')
    }
    // {
    //   path: '/engage',
    //   name: 'engage',
    //   component: () => import('../views/Engage')
    // }
  ],
  scrollBehavior() {
    // Scroll to top for all route navigations
    document.getElementById('main-app').scrollIntoView({ behavior: 'smooth' });
  }
});
