<template>
  <h2>登录</h2>
  <input v-model="form.username" placeholder="admin" />
  <input v-model="form.password" type="password" placeholder="123456" />
  <button @click="onSubmit">登录</button>
</template>

<script setup>
import { reactive } from 'vue'
import axios from 'axios'
const form = reactive({ username: 'admin', password: '123456' })
async function onSubmit() {
  const res = await axios.post('http://localhost:8080/login', form)
  localStorage.setItem('token', res.data.token)
  location.href = '/articles' // 跳转到表格页
}
</script>
