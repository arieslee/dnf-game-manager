import Vue from 'vue';
import VueRouter from 'vue-router';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import tokenHelper from "@/utils/token";


Vue.use(VueRouter);

const need = require('./_import_' + process.env.NODE_ENV); // 获取组件的方法
/**
 * 重写路由的push方法
 */
const routerPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
    return routerPush.call(this, location).catch(error => error)
}
const childRoutes = [];
// if(homeRoutes && homeRoutes.length > 0){
//     childRoutes.push(...homeRoutes);
// }

const constantRoutes = [
    {
        path: '/',
        component: need('/layout/index'),
        children: childRoutes,
        name: '',
        meta: {}
    },
    {
        path: '/login',
        component: need('/login/index'),
        hidden: true,
        meta: {
            title: '登录'
        }
    },
    // 修复无首页组件不显示404
    {
        path: '',
        hidden: true
    },
    {
        path: '*',
        component: need('/404'),
        hidden: true
    },
];
const router = new VueRouter({
    mode: 'history',
    scrollBehavior() {
        return {x: 0, y: 0}
    },
    routes: constantRoutes,
});
const loginRoutePath = '/'+process.env.VUE_APP_LOGIN_PATH;
const whiteList = [loginRoutePath]; // no redirect whitelist
const indexRoutePath = process.env.VUE_APP_INDEX_PATH;
const formatPageTitle = (pageTitle) => {
    if (pageTitle) {
        return `${pageTitle} - ${process.env.VUE_APP_TITLE}`;
    }
    return process.env.VUE_APP_TITLE;
}
const setDocumentTitle = (title) => {
    document.title = title;
}
const auth = () => {
    let accessToken = tokenHelper.get();
    if (!accessToken) {
        let refreshToken = tokenHelper.getRefreshToken();
        if (!refreshToken) {
            return false;
        }
    }
    return true;
}

router.beforeEach((to, from, next) => {
    NProgress.start()
    to.meta && (typeof to.meta.title !== 'undefined' && setDocumentTitle(formatPageTitle(to.meta.title)))
    if (to.path === loginRoutePath) {
        tokenHelper.clean();
    } else if (to.path === "/") {
        next({path: indexRoutePath});
    }
    if (whiteList.includes(to.path)) {
        next();
    } else {
        if (!auth()) {
            next({path: loginRoutePath});
        } else {
            next();
        }
    }
});

router.afterEach(() => {
    NProgress.done() // finish progress bar
})
export default router;
