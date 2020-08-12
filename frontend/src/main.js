import Vue from 'vue';
import App from './App.vue';
import router from './plugins/router';
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;
Vue.prototype.$log = console.log;

new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
