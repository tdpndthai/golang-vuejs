<template>
  <div>
    <div class="col-md-4"></div>
    <div class="col-md-3">
      <!-- <div>
        <div class="form-group">
          <label>Email address</label>
          <input type="email" class="form-control" v-model="login.email" />
        </div>
        <div class="form-group">
          <label>Password</label>
          <input type="password" class="form-control" v-model="login.password" />
        </div>
        <button class="btn btn-primary" @click="submit">Submit</button>
      </div>-->

      <h2>{{$t('login')}}</h2>
      <div class="info">
        <div v-if="error" class="alert alert-danger">Có lỗi xảy ra khi đăng nhập</div>
      </div>
      <input type="email" class="form-control" v-model="login.email" :placeholder="$t('email')" />
      <input
        type="password"
        class="form-control"
        v-model="login.password"
        :placeholder="$t('password')"
      />
      <button class="btn btn-primary" @click="submit">{{$t('login')}}</button>
      <router-link to="createuser" tag="button" class="btn btn-primary">{{$t('creatuser')}}</router-link>
      <div class="btn-group">
        <a @click="setLocale('vi')" href="#">Tiếng Việt</a>
        <a @click="setLocale('en')" href="#">English</a>
      </div>
    </div>
    <div class="col-md-4"></div>
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
        nickname: "",
        email: "",
        password: ""
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
          //console.log(token);
          this.token = token;
          if (this.token != null) {
            this.$router.push({
              name: "home",
              params: { userprop: this.user }
            });
            this.$store.dispatch("SetToken", this.token);
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
      this.$http
        .get("getbytoken", {
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
          //this.user = response.data
          var user = response.data;
          //mapActions(["SetCustomer"]);
          this.$store.dispatch("SetCustomer", user);
          //console.log(user);
        });
    },
    setLocale(locale) {
      this.$store.dispatch("SetLang", locale);
      this.$i18n.locale = locale;
    }
  },
  created() {
    console.log("khởi động created");
  },
  beforeCreate() {
    console.log("khởi động beforeCreate");
    const vuex = JSON.parse(localStorage.getItem("vuex"));
    if (vuex) {
      var locale = vuex.lang;
      this.$i18n.locale = locale;
    }
  },
  beforeUpdate() {
    console.log("khởi động beforeUpdate");
  },
  updated() {
    console.log("khởi động updated");
  },
  mounted() {
    if (localStorage.getItem("vuex")) {
      this.$router.push({
        name: "home"
      });
    }
  }
};
</script>

<style scoped>
html,
body {
  width: 100%;
  height: 100%;
  margin: 0px;
  font-family: "Work Sans", sans-serif;
}

body {
  background-image: url("https://images-assets.nasa.gov/image/6900952/6900952~orig.jpg");
  background-size: cover;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  color: #fff;
}

section {
  background-color: rgba(0, 0, 0, 0.72);
  width: 25%;
  min-height: 25%;
  display: flex;
  flex-direction: column;
  /*margin-left:auto;
	margin-right:auto;*/
}
form {
  display: flex;
  flex-direction: column;
  padding: 15px;
}
h2 {
  font-family: "segoe ui", sans-serif;
  color: #23a184;
  margin-left: auto;
  margin-right: auto;
}

.info {
  width: 100%;
  /* padding 1em 5px; */
  text-align: center;
  min-height: 45px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* .info.error{
	border:1px solid #a90e0;
	background-color: #ff3c41;
	color:#a90e0;
} */

.info p {
  margin: auto;
  padding: 5px;
}
.info.good {
  border: 1px solid #416d50;
  background-color: #47cf73;
  color: #416d50;
}

input {
  height: 35px;
  padding: 5px 5px;
  margin: 10px 0px;
  background-color: #e0dada;
  border: none;
}
button {
  height: 40px;
  padding: 5px 5px;
  margin: 10px 0px;
  font-weight: bold;
  background-color: #5273be;
  border: none;
  color: #e0dada;
  cursor: pointer;
  font-size: 16px;
}
button:hover {
  background-color: #5273be;
}

@-webkit-keyframes shake {
  from,
  to {
    -webkit-transform: translate3d(0, 0, 0);
    transform: translate3d(0, 0, 0);
  }
  10%,
  30%,
  50%,
  70%,
  90% {
    -webkit-transform: translate3d(-10px, 0, 0);
    transform: translate3d(-10px, 0, 0);
  }
  20%,
  40%,
  60%,
  80% {
    -webkit-transform: translate3d(10px, 0, 0);
    transform: translate(10px, 0, 0);
  }
}

.shake {
  animation-name: shake;
  animation-duration: 1s;
  /*animation-fill-mode: both;*/
}

@media screen and (max-width: 780px) {
  section {
    width: 90%;
  }
}
</style>>