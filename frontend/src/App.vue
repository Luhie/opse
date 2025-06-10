<template>
  <div id="app">
    <nav class="navbar">
      <router-link to="/">🏠 홈</router-link>
      <router-link to="/assets">📋 자산 목록</router-link>
      <router-link to="/assets/create">➕ 자산 등록</router-link>
    </nav>
    <main>
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
#app {
  min-height: 100vh;
  background-color: #1b1b1b;
  color: white;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.navbar {
  background-color: #2a2a2a;
  display: flex;
  gap: 1rem;
  padding: 1rem 2rem;
}

.navbar a {
  color: #eee;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
}

.navbar a.router-link-active {
  background-color: #4CAF50;
  color: white;
}

main {
  padding: 2rem;
}
</style>
