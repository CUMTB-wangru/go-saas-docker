<template>
  <div class="demo-price">
    <el-row>
      <div class="slogan">多种方案，适用多种场景</div>
    </el-row>
    <el-row :gutter="24" id="choice">
      <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
        <div id="card">
          <div class="panel-heading" style="text-align: center">免费版</div>
          <div class="panel-body">
            <p style="text-align: center; font-size: 25px">¥ 0 / 年</p>
          </div>
          <div class="input-num">
            <el-button size="mini" style="margin-right: -5px">数量</el-button>
            <el-input-number
              v-model="num"
              size="mini"
              :min="1"
              :max="10"
              controls-position="right"
              @change="handleChange"
            />
            <el-button type="primary" size="mini">立即购买</el-button>
          </div>
          <div class="list-group">最大项目数量：5</div>
          <div class="list-group">最大项目空间：5G</div>
          <div class="list-group">最大文件支持：5M</div>
          <div class="list-group">项目成员：5</div>
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
        <div id="card">
          <div class="panel-heading" style="text-align: center">免费版</div>
          <div class="panel-body">
            <p style="text-align: center; font-size: 25px">¥ 500 / 年</p>
          </div>
          <div class="input-num">
            <el-button size="mini" style="margin-right: -5px">数量</el-button>
            <el-input-number
              v-model="num"
              size="mini"
              :min="1"
              :max="10"
              controls-position="right"
              @change="handleChange"
            />
            <el-button type="primary" size="mini">立即购买</el-button>
          </div>
          <div class="list-group">最大项目数量：50</div>
          <div class="list-group">最大项目空间：50G</div>
          <div class="list-group">最大文件支持：50M</div>
          <div class="list-group">项目成员：50</div>
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8">
        <div id="card">
          <div class="panel-heading" style="text-align: center">免费版</div>
          <div class="panel-body">
            <p style="text-align: center; font-size: 25px">¥ 1000 / 年</p>
          </div>
          <div class="input-num">
            <el-button size="mini" style="margin-right: -5px">数量</el-button>
            <el-input-number
              v-model="num"
              size="mini"
              :min="1"
              :max="10"
              controls-position="right"
              @change="handleChange"
            />
            <el-button type="primary" size="mini">立即购买</el-button>
          </div>
          <div class="list-group">最大项目数量：500</div>
          <div class="list-group">最大项目空间：500G</div>
          <div class="list-group">最大文件支持：500M</div>
          <div class="list-group">项目成员：500</div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import internal from "stream";
import { defineComponent, ref } from "vue";
export default defineComponent({
  name: "Price",
  setup() {
    const num = ref(1);
    const handleChange = (value: any) => {
      console.log(value);
    };
    return {
      num,
      handleChange,
    };
  },
});
</script>

<style scoped>
.el-col {
  height: 400px;
  margin-bottom: 30px;
}
.demo-price {
  width: 100%;
  margin-bottom: 135px;
}
.slogan {
  font-size: 40px;
  text-align: center;
  margin: 100px auto;
}

#card {
  border: 1px solid #c0c0c0;
  border-radius: 4px;
  box-shadow: 0px 0px 10px #c0c0c0;
}
.panel-heading {
  color: white;
  padding: 10px 0;
  border-top-left-radius: 4px;
  border-top-right-radius: 4px;
  background-color: #545c64;
}

#choice {
  width: 90%;
  left: 5%;
  right: 5%;
  margin: 0 !important;
}
.panel-body,
.input-num,
.list-group {
  padding: 10px 40px;
  text-align: left;
  border-top: 1px solid #c0c0c0;
}
</style>
