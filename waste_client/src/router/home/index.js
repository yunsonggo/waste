export default {
  path: "/",
  component: () => import("@/views/Home"),
  children: [
    {
      path: "city",
      component: () => import("@/components/City"),
    },
    {
      path: "hot",
      component: () => import("@/components/Hot"),
    },
    {
      path: "more",
      component: () => import("@/components/More"),
    },
    {
      path: "search",
      component: () => import("@/components/Search"),
    },
    {
      path: "city",
      component: () => import("@/components/City"),
    },
    {
      path: "detail/hot/:detailId",
      components: {
        default: () => import("@/views/Detail"),
      },
      props: {
        default: true,
      },
    },
    {
        path: "detail/more/:detailId",
        components: {
          default: () => import("@/views/Detail"),
        },
        props: {
          default: true,
        },
      },
      {
        path: "order/add/:goodId",
        components: {
          default: () => import("@/views/Order/addorder.vue"),
        },
        props: {
          default: true,
        },
      },
    // {
    //   path: "detail/more/:detailId",
    //   component: () => import("@/views/Detail"),
    //   props: true,
    // },
    // {
    //     path: 'detail/hot/:detailId',
    //     components: {
    //         default: () => import('@/components/Hot'),
    //         // 对应 router-view的命名为 detail
    //         detail: () => import('@/views/Detail')
    //     },
    //     props: {
    //         detail: true
    //     }
    // },
    // {
    //     path: 'detail/more/:detailId',
    //     components: {
    //         default: () => import('@/components/More'),
    //         // 对应 router-view的命名为 detail
    //         detail: () => import('@/views/Detail')
    //     },
    //     props: {
    //         detail: true
    //     }
    // },
    // {
    //     path:'upload',
    //     component: () => import('@/views/UploadFile')
    // }
    {
      path: "/",
      redirect: "/hot",
    },
  ],
};
