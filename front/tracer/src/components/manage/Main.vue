<template>
  <div id="main">
    <el-row :gutter="10">
      <el-button type="primary" :icon="CirclePlusFilled" size="small"
        >Primary</el-button
      >
    </el-row>
    <el-row :gutter="10">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <div class="card-icon">
              <el-icon>
                <svg
                  viewBox="0 0 1024 1024"
                  xmlns="http://www.w3.org/2000/svg"
                  data-v-78e17ca8=""
                  style="display: inline"
                >
                  <path
                    fill="currentColor"
                    d="M283.84 867.84 512 747.776l228.16 119.936a6.4 6.4 0 0 0 9.28-6.72l-43.52-254.08 184.512-179.904a6.4 6.4 0 0 0-3.52-10.88l-255.104-37.12L517.76 147.904a6.4 6.4 0 0 0-11.52 0L392.192 379.072l-255.104 37.12a6.4 6.4 0 0 0-3.52 10.88L318.08 606.976l-43.584 254.08a6.4 6.4 0 0 0 9.28 6.72z"
                  ></path>
                </svg>
              </el-icon>
              <span>星标</span>
            </div>
          </div>
        </template>
        <div class="box-item">
            <div class="box-item-top">
                <span>项目1</span>
            </div>
            <div class="box-item-bottom">
                <div class="box-item-left">用户</div>
                <div class="box-item-right">人数</div>
            </div>
        </div>
      </el-card>
    </el-row>
    <el-row :gutter="10">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <div class="card-icon">
              <el-icon>
                <svg
                  viewBox="0 0 1024 1024"
                  xmlns="http://www.w3.org/2000/svg"
                  data-v-78e17ca8=""
                >
                  <path
                    fill="currentColor"
                    d="M160 448a32 32 0 0 1-32-32V160.064a32 32 0 0 1 32-32h256a32 32 0 0 1 32 32V416a32 32 0 0 1-32 32H160zm448 0a32 32 0 0 1-32-32V160.064a32 32 0 0 1 32-32h255.936a32 32 0 0 1 32 32V416a32 32 0 0 1-32 32H608zM160 896a32 32 0 0 1-32-32V608a32 32 0 0 1 32-32h256a32 32 0 0 1 32 32v256a32 32 0 0 1-32 32H160zm448 0a32 32 0 0 1-32-32V608a32 32 0 0 1 32-32h255.936a32 32 0 0 1 32 32v256a32 32 0 0 1-32 32H608z"
                  ></path>
                </svg>
              </el-icon>
              <span>我创建的</span>
            </div>
          </div>
        </template>
        <div class="box-item">
            <div class="box-item-top">
                <span>项目2</span>
            </div>
            <div class="box-item-bottom">
                <div class="box-item-left">用户</div>
                <div class="box-item-right">人数</div>
            </div>
        </div>
      </el-card>
    </el-row>
    <el-row :gutter="10">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <div class="card-icon">
              <el-icon>
                <svg
                  viewBox="0 0 1024 1024"
                  xmlns="http://www.w3.org/2000/svg"
                  data-v-78e17ca8=""
                >
                  <path
                    fill="currentColor"
                    d="M704 192h160v736H160V192h160v64h384v-64zM288 512h448v-64H288v64zm0 256h448v-64H288v64zm96-576V96h256v96H384z"
                  ></path>
                </svg>
              </el-icon>
              <span>我参与的</span>
            </div>
          </div>
        </template>
        <div class="box-item">
            <div class="box-item-top">
                <span>项目3</span>
            </div>
            <div class="box-item-bottom">
                <div class="box-item-left">用户</div>
                <div class="box-item-right">人数</div>
            </div>
        </div>
      </el-card>
    </el-row>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CirclePlusFilled } from "@element-plus/icons-vue";
export default defineComponent({
  name: "Main",
  setup() {
    return {
      CirclePlusFilled,
    };
  },
});
</script>

<style scoped>
#main {
  width: 94%;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  padding: 10px;
  /* border: 1px solid black; */
}
.el-row {
  margin-top: 10px;
}
.box-card {
  width: 100%;
}
.card-header {
  height: 5px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.text {
  font-size: 14px;
}
.item {
  margin-bottom: 18px;
}
:deep(.el-card__header) {
    background-color: #DCDCDC;
}
.box-item {
    width: 200px;
    height: 160px;
    border: 1px solid #DCDCDC;
    border-radius: 5px;
}
.box-item-top {
    text-align: center;
    height: 120px;
    line-height: 120px;
    background-color: #DCDCDC;
}
.box-item-bottom {
    overflow: hidden;
}
.box-item-left {
    margin-top: 10px;
    margin-left: 10px;
    float: left;
}
.box-item-right {
    margin-top: 10px;
    margin-right: 10px;
    float: right;
}
</style>
