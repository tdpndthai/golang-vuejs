<template>
  <div class="container">
    <div class="row">
      <div class="col-md-3">
        <div>
          <div class="form-group">
            <label>Email address</label>
            <input type="email" class="form-control" v-model="user.email" />
          </div>
          <div class="form-group">
            <label>Password</label>
            <input type="password" class="form-control" v-model="user.password" />
          </div>
          <button class="btn btn-primary" @click="submit">Submit</button>
          <button class="btn btn-primary" @click="fetchData">FetchData</button>
        </div>
      </div>
      <div v-if="error" class="alert alert-secondary">Có lỗi xảy ra khi đăng nhập</div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: {
        email: "",
        password: ""
      },
      users: [],
      resource: {},
      token:"",
      error: false
    };
  },
  methods: {
    submit() {
      this.$http.post("login", this.user)
        .then(
          response => {
            console.log(response);
            var token = response.data;
            //console.log(token);
            this.token = token
            if (this.token != null){
              this.$router.push({name:"register"})
            }
          },
          error => {
            console.log(error);
            this.error = true;
          }
        );
    },
    fetchData() {
      this.$http
        .get("users")
        .then(response => {
          console.log(response);
          //return response.json();
        })
        .then();
      // .then(data => {
      //   const resultArray = [];
      //   for (let key in data) {
      //     resultArray.push(data[key]);
      //   }
      //   this.users = resultArray;
      // });
    }
  }
};
</script>

<style>
</style>