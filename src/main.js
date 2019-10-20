import Vue from 'vue';
import VueRouter from 'vue-router';
import App from './App.vue';
import router from './router';
import vuetify from './plugins/vuetify';

Vue.use(VueRouter);
Vue.config.productionTip = false;
new Vue({
  router,
  vuetify,
  // el: '#app',
  render: h => h(App),
}).$mount('#app'); // does almost the same thing as el: #app
