export default {
    path:'/about',
    name:'about',
    component:() => import('@/views/About'),
    children: [
        {
            path:'login',
            component: () => import('@/components/Login') 
        },
        {
            path:'register',
            component: () => import('@/components/SignUp/EmailSign.vue') 
        },
        {
            path:'center',
            component: () => import('@/views/About/mycenter.vue') 
        },
        {
            path:'reset/pass',
            component: () => import('@/views/About/resetPass.vue') 
        },
        {
            path:'edit/info',
            component: () => import('@/views/About/editinfo.vue') 
        },
        {
            path:'address',
            component: () => import('@/views/About/myaddress.vue') 
        },
        {
            path : '/about',
            redirect : 'login'
        }
    ]
}