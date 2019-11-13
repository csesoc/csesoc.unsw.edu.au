import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';

Vue.use(Router);

export default new Router({

  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/post/:id',
      name: 'Post',
      component: () => import('./views/Post.vue'),
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
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
