<template>
  <div>
    <div class="progress" style="background-color:#b7c1d9 !important">
      <div
        class="progress-bar bg-primary"
        role="progressbar"
        :style="{width:percent+'%'}"
        aria-valuemin="0"
        aria-valuemax="100"
      >{{percent}}%</div>
    </div>
    <input type="file" style="display:none" ref="fileInput" @change="onFileSelected" name="myfile" />
    <button class="btn btn-primary" @click="$refs.fileInput.click()">Chose Image</button>
    <button class="btn btn-alert" @click="onUpload">Upload Image</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedFile: null,
      token: "",
      photo: [],
      percent: 0
    };
  },
  methods: {
    onFileSelected(event) {
      this.selectedFile = event.target.files[0];
    },
    onUpload() {
      const fd = new FormData();
      fd.append("myfile", this.selectedFile, this.selectedFile.name);
      this.$http
        .post("uploadfile", fd, {
          headers: {
            Authorization: "Bearer " + this.token
          },
          progress(e) {
            //console.log(e);
            if (e.lengthComputable) {
              this.percent = Math.round((e.loaded / e.total) * 100);
            }
          }
        })
        .then(
          response => {
            //console.log(response.data);
            var photoUpload = response.data;
            this.photo = photoUpload;
            this.$store.dispatch("SetPhoto", photoUpload);
          },
          error => {
            console.log(error);
          }
        );
    }
  },
  beforeUpdate() {
    GetToken: {
      this.token = this.$store.getters.getToken;
    }
  }
};
</script>

<style>
</style>