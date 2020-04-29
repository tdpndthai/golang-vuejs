<template>
  <div>
    <p>User Edit</p>
    <div>
      <div class="form-group">
        <label>{{$t('nickname')}}</label>
        <input class="form-control" v-model="Edit.nickname">
      </div>
      <div class="form-group">
        <label for="email">{{$t('email')}}:</label>
        <input type="email" class="form-control" v-model="Edit.email">
      </div>
      <div class="form-group">
        <label for="pwd">{{$t('password')}}:</label>
        <input type="password" class="form-control" v-model="Edit.password">
      </div>
      <button type="submit" class="btn btn-default" @click="update">Submit</button>
    </div>

  </div>
</template>

<script>
export default {
  data(){
    return{
      Edit: {
        nickname:"",
        email: "",
        password: ""
      },
      token:"",
    }
  },
  methods:{
    update() {
      var user = this.Edit
      var id = this.GetUser.id
      this.$http.put("users/edit/"+id, user,{
          headers: {
            Authorization: "Bearer " + this.token
          }
      }).then(
        response => {
          console.log(response.data);
          this.$router.push({
            name: "login",
            params: { userprop: this.user }
          });
        },
        error => {
          console.log(error);
          this.error = true;
        }
      );
    },
  }
  ,
  computed: {
    GetUser() {
      return this.$store.getters.getUser;
    },
    GetToken(){
      this.token = this.$store.getters.getToken;
      return this.$store.getters.getToken;
    }
  },
  created(){
    this.token = this.GetToken;
  }
}
</script>

<style>

</style>