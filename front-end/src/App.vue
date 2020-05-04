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
import ImageUpload from './components/image_upload/ImageUpload'
export default {
  name: "app",
  data() {
    return {};
  },
  beforeCreate() {
    console.log("khởi động beforeCreate");
    const vuex = JSON.parse(localStorage.getItem("vuex"));
    if (vuex) {
      var locale = vuex.lang;
      this.$i18n.locale = locale;
    }
  },
  components:{
    ImageUpload,
  }
};
</script>

<style>
