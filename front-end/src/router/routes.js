import Home from '../components/home/Home';
import Login from '../components/login_register/Login.vue'
import CreateUser from '../components/login_register/CreateUser.vue'
import UserDetail from '../components/user/UserDetail.vue'
import UserEdit from '../components/user/UserEdit.vue'
import User from '../components/user/User.vue';

export const routes=[
    {
        path:'',component:Login,name:'login'
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