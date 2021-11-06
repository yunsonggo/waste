import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import Message from 'element-ui'
import Notification  from 'element-ui'
import Loading from './components/Loading'
Vue.component('Loading',Loading);
import MessageBox from 'element-ui';

Vue.config.productionTip = false
Vue.prototype.$message = Message
Vue.prototype.$notify = Notification
Vue.prototype.$msgbox = MessageBox
Vue.use(ElementUI)


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
