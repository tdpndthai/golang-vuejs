import Vue from "vue";
import Vuex from "vuex";
import createPersistedState from "vuex-persistedstate";

import customer from "./modules/customer.js";
import token from "./modules/token.js";
import photo from "./modules/photo.js";

import * as getters from "./getters.js";
import * as actions from "./actions.js";
import * as mutations from "./mutations.js";
Vue.use(Vuex);

export const store = new Vuex.Store({
  state: {
    lang: "vi"
  },
  plugins: [createPersistedState()],
  getters,
  mutations,
  actions,
  // getters: {
  //   getUser: state => {
  //     return state.customer;
  //   },
  //   getToken: state => {
  //     return state.token;
  //   },
  //   getLangs: state => {
  //     return state.lang;
  //   }
  // },
  // mutations: {
  //   SetCustomer(state, cus) {
  //     state.customer = cus;
  //   },
  //   EditCustomer(state, cus) {
  //     state.customer = cus;
  //   },
  //   SetToken(state, token) {
  //     state.token = token;
  //   },
  //   SetLang(state, lang) {
  //     state.lang = lang;
  //   }
  // },
  // actions: {
  //   SetCustomer: ({ commit }, cus) => {
  //     commit("SetCustomer", cus);
  //   },
  //   EditCustomer: ({ commit }, cus) => {
  //     commit("EditCustomer", cus);
  //   },
  //   SetToken: ({ commit }, token) => {
  //     commit("SetToken", token);
  //   },
  //   SetLang: ({ commit }, lang) => {
  //     commit("SetLang", lang);
  //   }
  // },
  modules: {
    customer: customer,
    token: token,
    photo
  }
});
