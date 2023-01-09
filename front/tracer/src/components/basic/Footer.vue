<template>
  <div id="gaba">
    <a target="_blank" href="http://www.beian.gov.cn/">
      <img :src="src" style="float: left" />
      <p>京公网安备 11010502048370号</p>
    </a>
    <a target="_blank" href="https://beian.miit.gov.cn/">
      <p>吉ICP备2021006252号</p>
    </a>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Footer",
  setup() {
    return {
      src: "images/gaba.png",
    };
  },
});
</script>

<style scoped>
#gaba {
  width: 300px;
  height: 44px;
  margin: 0 auto;
  padding: 8px 0;
}
a {
  display: inline-block;
  text-decoration: none;
  height: 20px;
  line-height: 20px;
}
p {
  float: left;
  height: 20px;
  line-height: 20px;
  margin: 0px 0px 0px 5px;
  color: #939393;
}
</style>
