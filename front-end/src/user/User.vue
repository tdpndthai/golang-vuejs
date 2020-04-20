<template>
  <div>
    <h1>Component User {{GetUser.id}}</h1>
    <router-link
      :to="{name:'useredit',params:{id:GetUser.id}}"
      tag="a"
      class="btn btn-primary"
      name="useredit"
    >Edit user</router-link>
    <router-link :to="{name:'userdetail',params:{id:GetUser.id}}" tag="a" class="btn btn-primary" name="useredit">Detail User</router-link>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  computed: {
    GetUser() {
      return this.$store.getters.getUser;
    }
  }
};
</script>

<style>
</style>