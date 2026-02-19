<script setup>
defineProps(['title', 'rows']) 
</script>

<template>
    <div class="result-section">
        <div class="section-label">{{ title }}</div>
        <div class="result-rows">
        <div v-for="n in rows" :key="n" class="result-row skeleton">
            <span></span><span></span>
        </div>
        </div>
    </div>
</template>

<style scoped>

.result-section {
  background: #111;
  border: 1px solid #1e1e1e;
  border-radius: 12px;
  padding: 1.25rem 1.5rem;
}

.section-label {
  font-family: 'DM Mono', monospace;
  font-size: 0.68rem;
  letter-spacing: 0.12em;
  color: #555;
  text-transform: uppercase;
  margin-bottom: 1rem;
}

.result-rows {
  display: flex;
  flex-direction: column;
  gap: 0.6rem;
}

.result-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem 0;
  border-bottom: 1px solid #1a1a1a;
  font-family: 'DM Mono', monospace;
  font-size: 0.85rem;
}

.result-row:last-child {
  border-bottom: none;
}

.result-row.skeleton span {
  display: block;
  height: 12px;
  border-radius: 4px;
  background: #1e1e1e;
}

.result-row.skeleton span:first-child { width: 30%; }
.result-row.skeleton span:last-child { width: 45%; }

</style>