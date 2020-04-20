import Vue from "vue";
import Vuex from "vuex";
import createPersistedState from "vuex-persistedstate";
import * as Cookies from "js-cookie";

Vue.use(Vuex);

export const store = new Vuex.Store({
  state: {
    customer: {
      id: Number,
      nickname:String,
      email: String,
      password: String
    },
    token:""
  },

  getters: {
    getUser: state => {
      return state.customer;
    },
    getToken: state => {
      return state.token;
    }
  },
  actions: {
    SetCustomer: ({ commit }, cus) => {
      commit("SetCustomer", cus);
    },
    EditCustomer: ({ commit }, cus) => {
      commit("EditCustomer",cus)
    },
    SetToken: ({ commit }, token) => {
      commit("SetToken", token);
    }
  },
  mutations: {
    SetCustomer(state, cus) {
      state.customer = cus;
    },
    EditCustomer(state, cus) {
      state.customer = cus;
    },
    SetToken(state, token) {
      state.token = token;
    }
  },
  plugins: [createPersistedState()]
});
