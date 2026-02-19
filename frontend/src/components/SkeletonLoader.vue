<script setup>
const rowCounts = [7, 9, 5]
</script>

<template>
  <div class="skeleton-section" v-for="n in 3" :key="n">
    <div class="skeleton-label shimmer" />
    <div class="skeleton-row shimmer" v-for="r in rowCounts[n - 1]" :key="r" />
  </div>
</template>

<style scoped>
.skeleton-section {
  background: #111;
  border: 1px solid #1e1e1e;
  border-radius: 12px;
  padding: 1.25rem 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.skeleton-label {
  width: 30%;
  height: 10px;
  border-radius: 4px;
  margin-bottom: 0.25rem;
}

.skeleton-row {
  width: 100%;
  height: 14px;
  border-radius: 4px;
}

.shimmer {
  background: linear-gradient(90deg, #1a1a1a 25%, #2a2a2a 50%, #1a1a1a 75%);
  background-size: 200% 100%;
  animation: shimmer 1.4s infinite;
}

@keyframes shimmer {
  0%   { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}
</style>