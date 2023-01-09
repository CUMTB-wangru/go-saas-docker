import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router";
import basic from "/@/router/basic/index";
import manage from "/@/router/manage/index";

const routes = [
  ...basic,
  ...manage,
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
