/* eslint-disable */
import axios from 'axios'
// 下面你的代码不动
const request = axios.create({
    baseURL: 'http://localhost:8080', // 后端地址
    timeout: 5000
})
// 请求拦截：自动带 token
request.interceptors.request.use(
    config => {
        const token = localStorage.getItem('token')
        if (token) config.headers.Authorization = 'Bearer ' + token
        return config
    },
    error => Promise.reject(error)
)
// 响应拦截：统一弹错误
request.interceptors.response.use(
    res => res.data,
    err => {
        ElMessage.error(err.response?.data?.msg || '网络错误')
        return Promise.reject(err)
    }
)
export default request
