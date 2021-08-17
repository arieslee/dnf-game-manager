const homeRoutes = [
    {
        path: 'home/index',
        acl: 'home/index',
        name: 'home_index',
        parent:'',
        component: () => import('@/views/home/index.vue'),
        meta: {
            title: '首页',
        }
    },
    {
        path: 'tool',
        acl: 'tool',
        name: 'tool',
        parent:'',
        component: () => import('@/views/tool/index.vue'),
        meta: {
            title: '工具',
        }
    },
    {
        path: 'role',
        acl: 'role',
        name: 'role',
        parent:'',
        component: () => import('@/views/role/index.vue'),
        meta: {
            title: '角色',
        }
    },
];

export default homeRoutes;
