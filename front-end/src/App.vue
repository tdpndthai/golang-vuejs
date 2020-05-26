<template>
  <div id="app">
    <div class="container">
      <div class="row">
        <ImageUpload />
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import ImageUpload from "./components/image_upload/ImageUpload";
export default {
  name: "app",
  data() {
    return {};
  },
  methods: {
    removeLocalStorage() {
      localStorage.removeItem("vuex");
    },
    checkLogin() {
      const vuex = JSON.parse(localStorage.getItem("vuex"));
      if (vuex == null) {
        this.$router.push({
          name: "login",
          props: {
            error: "Time out login"
          }
        });
      }
    }
  },
  beforeUpdate() {
    //console.log("khởi động beforeCreate");
    const vuex = JSON.parse(localStorage.getItem("vuex"));
    if (vuex) {
      var locale = vuex.lang;
      this.$i18n.locale = locale;
    }
    setTimeout(() => {
      this.removeLocalStorage();
    }, 20000);
  },
  mounted() {
    console.log("mounted");
    setInterval(() => {
      this.checkLogin();
    }, 20000);
  },
  components: {
    ImageUpload
  }
};
</script>

<style>
