import Home from '../home/Home.vue';
import Login from '../login_register/Login.vue'
import CreateUser from '../login_register/CreateUser.vue'
import UserDetail from '../user/UserDetail.vue'
import UserEdit from '../user/UserEdit.vue'
import User from '../user/User.vue';

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