<script setup>
import { ref } from 'vue'
const emit = defineEmits(['file-dropped'])
const isDragging = ref(false)

const handleDragOver = (e) => { e.preventDefault(); isDragging.value = true }
const handleDragLeave = () => { isDragging.value = false }
const handleDrop = (e) => {
  e.preventDefault()
  isDragging.value = false
  const files = e.dataTransfer.files
  if (files.length) emit('file-dropped', files[0])
}
</script>

<template>
  <div
      class="upload-zone"
      :class="{ dragging: isDragging }"
      @dragover="handleDragOver"
      @dragleave="handleDragLeave"
      @drop="handleDrop"
    >
      <div class="upload-inner">
        <div class="upload-icon">
          <svg width="48" height="48" viewBox="0 0 48 48" fill="none">
            <rect x="8" y="8" width="32" height="32" rx="4" stroke="currentColor" stroke-width="2"/>
            <circle cx="19" cy="20" r="3" stroke="currentColor" stroke-width="2"/>
            <path d="M8 32l10-10 6 6 6-8 10 10" stroke="currentColor" stroke-width="2" stroke-linejoin="round"/>
          </svg>
        </div>
        <p class="upload-main">Drop your image here</p>
        <p class="upload-sub">or <label class="browse-link">browse files<input type="file" accept="image/*" hidden /></label></p>
        <div class="upload-formats">JPG · RAW · TIFF · WEBP · HEIC</div>
      </div>
    </div>
</template>

<style scoped> 
.upload-zone {
  width: 100%;
  max-width: 580px;
  border: 2px dashed #2a2a2a;
  border-radius: 16px;
  padding: 2rem 2rem;
  transition: border-color 0.2s, background 0.2s;
  cursor: pointer;
  background: #111;
}

.upload-zone:hover,
.upload-zone.dragging {
  border-color: #c8f060;
  background: rgba(200, 240, 96, 0.04);
}

.upload-inner {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.1rem;
  text-align: center;
}

.upload-icon {
  color: #444;
  margin-bottom: 0.3rem;
  transition: color 0.2s;
}

.upload-zone:hover .upload-icon,
.upload-zone.dragging .upload-icon {
  color: #c8f060;
}

.upload-main {
  font-size: 1.1rem;
  font-weight: 700;
  color: #e0ddd8;
  margin-bottom: 0rem;
}

.upload-sub {
  font-size: 0.9rem;
  color: #666;
  margin-top: 0.5rem;
}

.browse-link {
  color: #c8f060;
  cursor: pointer;
  text-decoration: underline;
  text-underline-offset: 3px;
}

.upload-formats {
  margin-top: 0.6rem;
  font-family: 'DM Mono', monospace;
  font-size: 0.7rem;
  letter-spacing: 0.1em;
  color: #444;
}
</style>