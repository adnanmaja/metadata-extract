<script setup>
import TheHeader from './components/TheHeader.vue'
import UploadZone from './components/UploadZone.vue'
import SkeletonLoader from './components/SkeletonLoader.vue'
import MetadataDisplay from './components/MetadataDisplay.vue'
import TheFooter from './components/TheFooter.vue'
import { useFileProcessor } from './composables/useFileProcessor'

const { processFile, isLoading, error, metadata } = useFileProcessor()
</script>

<template>
  <div class="page">
    <TheHeader />

    <UploadZone 
    @file-dropped="processFile"
    :loading="isLoading" />

    <div v-if="error" class="error-message">
      {{ error }}
    </div>

    <div class="results-placeholder" v-if="isLoading">
      <SkeletonLoader />
    </div>

    <div class="results-placeholder" v-if="metadata">
      <MetadataDisplay 
        title="Basic Info" 
        :metadata="metadata.basic" 
        category="basic" 
      />
      <MetadataDisplay 
        title="Camera Data" 
        :metadata="metadata.camera" 
        category="camera" 
      />
      <MetadataDisplay 
        title="Location" 
        :metadata="metadata.location" 
        category="location" 
      />
    </div>

    <TheFooter />
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