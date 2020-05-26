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
      photo: [],
      percent: 0
    };
  },
  methods: {
    onFileSelected(event) {
      this.selectedFile = event.target.files[0];
    },
    onUpload() {
      let vm = this
      const fd = new FormData();
      fd.append("myfile", this.selectedFile, this.selectedFile.name);
      var token = this.$store.getters.getToken;
      this.$http
        .post("uploadfile", fd, {
          headers: {
            Authorization: "Bearer " + token
          },
          progress(e) {
            //console.log(e);
            if (e.lengthComputable) {
              let currentProgress  = Math.round((e.loaded / e.total) * 100);
              vm.percent = currentProgress;
            }
          }
        })
        .then(
          response => {
            //console.log(response.data);
            var photoUpload = response.data;
            this.photo = photoUpload;
            this.$store.dispatch("SetPhoto", photoUpload);
            setTimeout(() => {
              this.clearStatus()
            },3000)
          },
          error => {
            console.log(error);
          }
        );
    },
    clearStatus(){
      this.selectedFile = null;
      this.percent = 0;
    }
  },
};
</script>

<style>
</style>