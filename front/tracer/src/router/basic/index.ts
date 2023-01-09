import Home from "/@/views/Home.vue"
import Main from "/@/components/basic/Main.vue"
import Login from "/@/components/basic/Login.vue"
import Register from "/@/components/basic/Register.vue"
import Price from "/@/components/basic/Price.vue"
import Product from "/@/components/basic/Product.vue"
import Help from "/@/components/basic/Help.vue"

const basic = [
    {
        path: "/",
        name: "Home",
        component: Home,
        redirect: "main",
        children: [
            {
                path: "main",
                name: "产品功能",
                component: Main,
            }, {
                path: "product",
                name: "企业方案",
                component: Product,
            }, {
                path: "help",
                name: "帮助文档",
                component: Help,
            }, {
                path: "price",
                name: "价格",
                component: Price,
            }, {
                path: "login",
                name: "登录",
                component: Login,
            }, {
                path: "register",
                name: "注册",
                component: Register,
            },
        ]
    },
]

export default basic
