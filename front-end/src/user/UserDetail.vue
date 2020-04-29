<template>
  <div>
    <table class="table table-hover table-bordered" style="margin-top:25px">
      <thead>
        <tr class="active">
          <th>#</th>
          <th>##</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>ID</td>
          <td>{{GetUser.id}}</td>
        </tr>
        <tr>
          <td>{{$t('nickname')}}</td>
          <td>{{GetUser.nickname}}</td>
        </tr>
        <tr>
          <td>{{$t('email')}}</td>
          <td>{{GetUser.email}}</td>
        </tr>
        <tr>
          <td>{{$t('password')}}</td>
          <td>{{GetUser.password}}</td>
        </tr>
      </tbody>
    </table>
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