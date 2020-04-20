<template>
  <div>
    <div class="col-md-4"></div>
    <div class="col-md-3">
      <div>
        <div class="form-group">
          <label>Email address</label>
          <input type="email" class="form-control" v-model="login.email" />
        </div>
        <div class="form-group">
          <label>Password</label>
          <input type="password" class="form-control" v-model="login.password" />
        </div>
        <button class="btn btn-primary" @click="submit">Submit</button>
      </div>
    </div>
    <div class="col-md-4"></div>
    <div v-if="error" class="alert alert-secondary">Có lỗi xảy ra khi đăng nhập</div>
    <router-link to="createuser" tag="a" class="btn btn-primary">Create User</router-link>
  </div>
</template>

<script>
export default {
  data() {
    return {
      login: {
        email: "",
        password: ""
      },
      user: {
        id: 1,
        nickname:"",
        email: "",
        password: "",
      },
      token: "",
      error: false
    };
  },
  methods: {
    submit() {
      this.$http.post("login", this.login).then(
        response => {
          console.log(response);
          var token = response.data;
          console.log(token);
          this.token = token;
          if (this.token != null) {
            this.$router.push({
              name: "home",
              params: { userprop: this.user }
            });
            this.$store.dispatch("SetCustomer", this.user);
            this.$store.dispatch("SetToken",this.token);
          }
          this.fetchData();
        },
        error => {
          console.log(error);
          this.error = true;
        }
      );
    },
    fetchData() {
      this.$http.get("getbytoken", {
          headers: {
            Authorization: "Bearer " + this.token
          }
        })
        .then(response => {
          //console.log(response.data);
          this.user.id = response.data.id;
          this.user.email = response.data.email;
          this.user.password = response.data.password;
          this.user.nickname = response.data.nickname;
          //console.log(this.user);
        });
    }
  }
};
</script>

<style>
</style>