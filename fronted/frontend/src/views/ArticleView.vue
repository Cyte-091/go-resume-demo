<template>
  <h2>文章列表</h2>
  <table border="1">
    <tr>
      <th>ID</th>
      <th>标题</th>
      <th>正文</th>
    </tr>
    <tr v-for="item in list" :key="item.id">
      <td>{{ item.id }}</td>
      <td>{{ item.title }}</td>
      <td>{{ item.body }}</td>
    </tr>
  </table>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import axios from 'axios'
const list = ref([])

onMounted(async () => {
  const token = localStorage.getItem('token')
  const res = await axios.get('http://localhost:8080/articles', {
    headers: { Authorization: 'Bearer ' + token }
  })

  console.log('原始返回', res)          // 先看控制台结构
  console.log('res.data', res.data)    // 再看 data 层级

  // 万能兜底：不管对象还是数组，都包成数组
  const raw = res.data
  if (Array.isArray(raw)) {
    list.value = raw
  } else if (raw) {
    list.value = [raw]                // 单条也包成数组
  } else {
    list.value = []
  }

  console.log('最终列表', list.value)  // 确认是数组
})
</script>
