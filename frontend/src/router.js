import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'root',
      component: Home,
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('./views/About.vue'),
    },
    {
      path: '/contact',
      name: 'contact',
      component: () => import('./views/Contact.vue'),
    },
    {
      path: '/events',
      name: 'events',
      component: () => import('./views/Events.vue'),
    },
    {
      path: '/media',
      name: 'media',
      component: () => import('./views/Media.vue'),
    },
    {
      path: '/members',
      name: 'members',
      component: () => import('./views/Members.vue'),
    },
    {
      path: '/merch',
      name: 'merch',
      component: () => import('./views/Merch.vue'),
    },
    {
      path: '/projects',
      name: 'projects',
      component: () => import('./views/Projects.vue'),
    },
    {
      path: '/resources',
      name: 'resources',
      component: () => import('./views/Resources.vue'),
    },
  ],
});
