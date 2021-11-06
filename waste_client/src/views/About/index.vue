<template>
    <div id="main">
        <Header title="个人中心"></Header>
        <div id="content">
            <router-view></router-view>
        </div>
        <Tabbar></Tabbar>
    </div>
</template>

<script>
import Header from '@/components/Header'
import Tabbar from '@/components/Tabbar'
import Login from '@/components/Login'
    export default {
        name:"About",
        components:{
            Header,
            Tabbar,
            Login
        }
    }
</script>

<style scoped>

</style>