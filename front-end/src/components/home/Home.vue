<template>
  <div class>
    <nav class="nav">
      <h3>home component</h3>
      <!-- {{GetUser.id}} -- {{GetUser.email}} -- {{GetUser.password}} -->
      <router-link to="user" tag="a" class="btn btn-primary">User</router-link>
    </nav>
  </div>
</template>

<script>
import LoginVue from '../login_register/Login.vue';
export default {
  data() {
    return {};
  },
  methods: {},
  computed: {
    GetUser() {
      return this.$store.getters.getUser;
      console.log(this.$store);
    },
  },
  mounted() {
    console.log("mouted");
  },
  beforeUpdate() {
    console.log("beforeUpdate");
    this.checkLogin;
  },
  beforeDestroy() {
    console.log("beforeDestroy");
  },
  destroyed() {
    console.log("destroyed");
  },
  beforeCreate() {
    console.log("beforeCreate");
  },
  created() {
    console.log("created");
  }
};
</script>

<style>
</style>