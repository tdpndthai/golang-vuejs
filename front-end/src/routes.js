import Home from './Home.vue';
import Login from './Login.vue'
import Register from './Register.vue';
import CreateUser from './user/CreateUser.vue'
import UserDetail from './user/UserDetail.vue'
import UserEdit from './user/UserEdit.vue'
import User from './user/User.vue';

export const routes=[
    {
        path:'/',component:Login,name:'login'
    },
    {
        path:'/home',component:Home,name:'home',props:true
    },
    {
        path:'/createuser',component:CreateUser,name:'createuser'
    },
    {
        path:'/register',component:Register,name:'register'
    },
    {
        path:"/user",component:User,name:'user',props:true,
        children:[
            {
                path:':id',component:UserDetail,name:'userdetail'
            },
            {
                path:'edit/:id',component:UserEdit,name:'useredit'
            },
        ]
    }
];