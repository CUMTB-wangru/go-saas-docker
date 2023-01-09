import { createApp } from 'vue'
import App from '/@/App.vue'
// 引入路由
import router from "./router/index"
// 引入ElementUI
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import * as ElIconModules from '@element-plus/icons-vue'
// 引入nprogress
import NProgress from 'nprogress' // 进度条
import 'nprogress/nprogress.css' // 这个样式必须引入
// 引入axios
import axios from 'axios'
import VueAxios from 'vue-axios'


const app = createApp(App)

// 全局图标组件注册
for (let iconName in ElIconModules) {
    // app.component(iconName, ElIconModules[iconName])
    app.component(iconName, ElIconModules)
}

router.beforeEach((to, from, next) => {
    NProgress.start()
    next()
})

router.afterEach(() => {
    NProgress.done()
})


app.use(router)
app.use(ElementPlus, { locale: zhCn })
app.use(VueAxios, axios);
app.mount('#app')

// 下面这种写法会导致ElementUI的css不渲染
// createApp(App).use(router, ElementPlus).mount('#app')
