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
export default {
  data() {
    return {
    };
  },
  computed: {
    GetUser() {
      return this.$store.getters.getUser;
      console.log(this.$store);
    }
  },
};
</script>

<style>
</style>