export default {
    path:'/order',
    component:() => import('@/views/Order'),
    children: [
        {
            path:'allorder',
            component: () => import('@/views/Order/allorder.vue')
        },
        {
            path:'notdone',
            component: () => import('@/views/Order/notdone.vue')
        },
        {
            path:'doneorder',
            component: () => import('@/views/Order/doneorder.vue')
        },
        {
            path:'/order',
            redirect:'/order/allorder'
        },
    ]
}