import Vue from "vue";
import App from "./App.vue";
import VueRouter from "vue-router";
import VueResource from "vue-resource";
import { i18n} from './plugins/i18n'

import { routes } from "./router/routes.js";
import { store } from "./store/store.js";
import createPersistedState from 'vuex-persistedstate';

Vue.use(VueRouter);
Vue.use(VueResource);

Vue.http.options.root = "http://localhost:3030/";

const router = new VueRouter({
  routes,
  mode: "history"
});

new Vue({
  el: "#app",
  router,
  store,
  i18n,
  render: h => h(App)
});
