<template>
    <el-skeleton animated :loading="loading">
      <template #template>
        <el-skeleton-item variant="image" style="width: 100%; height: 100%" />
        <el-skeleton-item variant="image" style="width: 100%; height: 100%" />
        <el-skeleton-item variant="image" style="width: 100%; height: 100%" />
        <el-skeleton-item variant="image" style="width: 100%; height: 100%" />
      </template>
      <template #default>
        <div class="demo-image__placeholder">
          <div class="block">
            <el-image :src="src"></el-image>
            <el-image :src="src"></el-image>
            <el-image :src="src"></el-image>
            <el-image :src="src"></el-image>
          </div>
        </div>
      </template>
    </el-skeleton>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: "Prodect",
  setup() {
    return {
      src: "images/index-1.png",
      loading: false,
    }
  },
})
</script>

<style scoped>
</style>
