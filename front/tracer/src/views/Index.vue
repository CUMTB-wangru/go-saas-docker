<template>
    <div id="index">
        <el-container>
            <el-header>
                <Header></Header>
            </el-header>
            <el-main>
                <router-view></router-view>
            </el-main>
            <!-- <el-footer>
                <Footer></Footer>
            </el-footer> -->
        </el-container>
        <!-- <Banner></Banner>
        <Footer></Footer> -->
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from "/@/components/manage/Header.vue"
export default defineComponent({
    name: "Index",
    components: {
        Header,
    },
    setup() { }
})
</script>

<style scoped>
.el-header,
.el-footer {
  background-color: #545c64;
  color: var(--el-text-color-primary);  /* 后面可以直接注释掉或者删除 */
  text-align: center;   /* 后面可以直接注释掉或者删除 */
}

.el-aside {
  background-color: #d3dce6;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 200px;
}

.el-main {
  /* background-color: #e9eef3; */
  /* color: var(--el-text-color-primary); */
  /* text-align: center; */
  overflow: hidden;
  padding: 0;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>
