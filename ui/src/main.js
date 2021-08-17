import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import vuetify from '@/plugins/vuetify';
Vue.config.productionTip = false;
import 'remixicon/fonts/remixicon.css';
import {LIGHT} from "@/config/theme";
Vue.prototype.theme = LIGHT;
let vm = new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app');
export default vm;
