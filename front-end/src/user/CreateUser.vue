<template>
  <div>
    <p>Create User</p>
    <div>
      <div class="form-group">
        <label>Nick Name</label>
        <input class="form-control" v-model="Create.nickname" />
      </div>
      <div class="form-group">
        <label for="email">Email:</label>
        <input type="email" class="form-control" v-model="Create.email" />
      </div>
      <div class="form-group">
        <label for="pwd">Password:</label>
        <input type="password" class="form-control" v-model="Create.password" />
      </div>
      <button type="submit" class="btn btn-default" @click="createUser">Submit</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      Create: {
        nickname: "",
        email: "",
        password: ""
      }
    };
  },
  methods: {
    createUser() {
      var user = this.Create;
      this.$http.post("createuser", user).then(
        response => {
          console.log(response.data);
          this.$router.push({
            name: "login"
          });
        },
        error => {
          console.log(error);
          this.error = true;
        }
      );
    }
  }
};
</script>

<style>
</style>