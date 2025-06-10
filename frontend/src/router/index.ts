import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/views/HomePage.vue'
import AssetList from '@/views/asset/AssetList.vue'
import AssetCreate from '@/views/asset/AssetCreate.vue'

const routes = [
    { path: '/', name: 'Home', component: HomePage },
    { path: '/assets', name: 'AssetList', component: AssetList },
    { path: '/assets/create', name: 'AssetCreate', component: AssetCreate }
]

export const router = createRouter({
    history: createWebHistory(),
    routes,
})