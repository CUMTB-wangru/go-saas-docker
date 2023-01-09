<template>
  <el-row :gutter="10">
    <el-col :xs="5" :sm="5" :md="4" :lg="3" :xl="3">
      <router-link :to="{path: '/main'}">
        <b style="color: white; line-height: 60px"> PubbcTracer </b>
      </router-link>
    </el-col>
    <el-col :xs="12" :sm="14" :md="16" :lg="18" :xl="18">
      <el-menu
        :default-active="route.path"
        router
        class="el-menu-demo"
        mode="horizontal"
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b"
      >
        <el-sub-menu index="7">
          <template #title>项目</template>
          <el-sub-menu index="7-1">
            <template #title>我创建的</template>
            <el-menu-item
              v-for="(item, i) in router.options.routes[1].children?.slice(0,1)"
              :key="i"
              :index="item.path"
            >{{item.name}}
            </el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="7-2">
            <template #title>我参与的</template>
            <el-menu-item
              v-for="(item, i) in router.options.routes[1].children?.slice(0,1)"
              :key="i"
              :index="item.path"
            >{{item.name}}
            </el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="7-3">
            <template #title>所有项目</template>
            <el-menu-item
              v-for="(item, i) in router.options.routes[1].children?.slice(0,1)"
              :key="i"
              :index="item.path"
            >{{item.name}}
            </el-menu-item>
          </el-sub-menu>
        </el-sub-menu>
        <el-menu-item
          v-for="(item, i) in router.options.routes[1].children?.slice(1,7)"
          :key="i"
          :index="item.path"
        >{{item.name}}
        </el-menu-item>
      </el-menu>
    </el-col>
    <el-col :xs="7" :sm="5" :md="4" :lg="3" :xl="3">
      <div class="demo-type">
        <el-dropdown>
          <span class="el-dropdown-link" style="color: white">
            <span style="line-height: 60px">未登录</span>
            <el-icon class="el-icon--right">
              <caret-bottom />
            </el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <!-- 这里使用两个0会使效果失效 -->
              <el-dropdown-item>
                <router-link :to="{ path: 'login' }" @click="activeIndex = '4'">
                  <b style="color:black">登录</b>
                </router-link>
              </el-dropdown-item>
              <el-dropdown-item>
                <router-link :to="{ path: 'register' }" @click="activeIndex = '5'">
                  <b style="color:black">注册</b>
                </router-link>
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
        <div class="demo-image">
          <!-- 判断登录状态显示头像 -->
          <el-avatar :icon="UserFilled" style="margin-top: 10px"></el-avatar>
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { UserFilled } from "@element-plus/icons-vue";
import { useRoute, useRouter } from "vue-router";
export default defineComponent({
  name: "Header",
  setup() {
    const router = useRouter();
    const route = useRoute();
    const activeIndex = ref('0');
    return {
        activeIndex,
        UserFilled,
        router,
        route,
    };
  },
});
</script>

<style scoped>
:deep(.el-menu--horizontal){
  border-bottom: 0;
}
.demo-type {
  height: 60px;
  overflow: hidden;
}
.demo-image {
  float: right;
  margin-right: 4px;
}
.el-dropdown {
  float: right;
  margin-left: 4px;
}

a {
  text-decoration: none;
}
</style>
