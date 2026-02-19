<script setup>
import TheHeader from './components/TheHeader.vue'
import UploadZone from './components/UploadZone.vue'
import MetadataDisplay from './components/MetadataDisplay.vue'

function processFile(file) {
  console.log("Got the file:", file.name)
  // Parsing logic stuff
}
</script>

<template>
  <div class="page">
    <TheHeader />

    <UploadZone @file-dropped="processFile" />

    <div class="results-placeholder">
      <MetadataDisplay title="Basic Info" :rows="3" />
      <MetadataDisplay title="Camera Data" :rows="4" />
      <MetadataDisplay title="Location" :rows="2" />
    </div>

    <footer class="footer">type shi</footer>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Syne:wght@400;700;800&family=DM+Mono:wght@400;500&display=swap');

*, *::before, *::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

.page {
  min-height: 100vh;
  background: #0a0a0a;
  color: #f0ede8;
  font-family: Helvetica, sans-serif;
  padding: 3rem 1.5rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 3rem;
}

.results-placeholder {
  width: 100%;
  max-width: 580px;
  display: flex;
  flex-direction: column;
  gap: 1.3rem;
}

.footer {
  font-family: 'DM Mono', monospace;
  font-size: 0.75rem;
  color: #333;
  letter-spacing: 0.08em;
}
</style>