import Home from './views/Home.vue';
import Login from './views/Login.vue'
import Register from './views/Register.vue';
import CreateUser from './views/CreateUser.vue'

export const routes=[
    {path:'/',component: Home,name:'home'},
    {path:'/login',component: Login,name:'login'},
    {path:'/register',component: Register,name:'register'},
    {path:'/creatuser',component: CreateUser,name:'createuser'},
];