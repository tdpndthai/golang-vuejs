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
      resource: {}
    };
  },
  methods: {
    submit() {
      this.$http.post("login", this.user, {
          headers: {
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
	          "Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
          }
        })
        .then(
          response => {
            console.log(response);
          },
          error => {
            console.log(error);
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