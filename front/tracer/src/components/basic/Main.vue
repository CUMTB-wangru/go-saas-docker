<template>
  <div class="demo-image__placeholder">
    <div class="block">
      <el-image :src="src_1"></el-image>
      <el-image :src="src_2"></el-image>
      <el-image :src="src_3"></el-image>
      <el-image :src="src_4"></el-image>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Main",
  setup() {
    return {
      src_1: "images/index-1.png",
      src_2: "images/index-2.png",
      src_3: "images/index-3.png",
      src_4: "images/index-4.png",
    };
  },
});
</script>

<style scoped>
.el-image {
  display: block;
}
</style>