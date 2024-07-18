import Vue from 'vue';
import VueGtag from 'vue-gtag';
import App from './App';
import router from './plugins/router';
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;
// Vue.prototype.$log = console.log;

Vue.use(VueGtag, {
  // The G-tag is not a secret and is available to view by the public
  config: { id: 'G-3D5099EXGN' },
});

new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
