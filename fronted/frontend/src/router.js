import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '@/views/LoginView.vue'
import ArticleView from '@/views/ArticleView.vue'
const routes = [
    { path: '/', redirect: '/login' },
    { path: '/login', component: LoginView },
    { path: '/articles', component: ArticleView }
]
export default createRouter({
    history: createWebHistory(),
    routes
})
