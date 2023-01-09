import Index from "/@/views/Index.vue"
import Main from "/@/components/manage/Main.vue"
import Dashboard from "/@/components/manage/Dashboard.vue"
import Issues from "/@/components/manage/Issues.vue"
import Statistics from "/@/components/manage/Statistics.vue"
import Wiki from "/@/components/manage/Wiki.vue"
import File from "/@/components/manage/File.vue"
import Settings from "/@/components/manage/Settings.vue"
// import Project form ""
const manage = [
    {
        path: "/home",
        name: "Index",
        component: Index,
        redirect: "/home/projects",
        children: [
            {
                path: "projects",
                name: "项目",
                component: Main,
                // children: [
                    // {
                    //     path: "project",
                    //     name: "项目",
                    //     component: Project,
                    // } 
                // ]
            }, {
                path: "dashboard",
                name: "概览",
                component: Dashboard,
            }, {
                path: "issues",
                name: "问题",
                component: Issues,
            }, {
                path: "statistics",
                name: "统计",
                component: Statistics,
            }, {
                path: "wiki",
                name: "wiki",
                component: Wiki,
            }, {
                path: "file",
                name: "文件",
                component: File,
            }, {
                path: "settings",
                name: "配置",
                component: Settings,
            },
        ]
    },
]
export default manage
